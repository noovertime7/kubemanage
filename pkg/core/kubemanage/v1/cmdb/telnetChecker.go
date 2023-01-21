package cmdb

import (
	"fmt"
	"net"
	"time"

	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/runtime"
	"github.com/noovertime7/kubemanage/runtime/queue"
)

func NewTelnetHandler(factory dao.ShareDaoFactory, queue queue.Queue, lg logger.Logger) *telnetHandler {
	handler := &telnetHandler{factory: factory, queue: queue, logger: lg}
	runtime.CloserHandler.AddCloser(handler)
	return handler
}

type telnetHandler struct {
	factory dao.ShareDaoFactory
	queue   queue.Queue
	logger  logger.Logger
}

func (t *telnetHandler) Close() error {
	t.queue.Close()
	t.logger.Info("telnet checker successful quit")
	return nil
}

func (t *telnetHandler) Run() {
	t.logger.Info("telnet checker start watch queue...")
	for {
		event, err := t.queue.Get()
		// 这里返回的错误对于checker来说都是致命的 所有需要直接return
		if err != nil {
			t.HandlerErr(err)
			return
		}
		if err := t.Check(event); err != nil {
			t.HandlerErr(err)
			return
		}
	}
}

func (t *telnetHandler) Check(event *queue.Event) error {
	host := event.Data.(*model.CMDBHost)
	timeout := time.Duration(config.SysConfig.CMDB.HostCheck.HostCheckTimeout) * time.Second
	addr := fmt.Sprintf("%s:%d", host.Address, host.Port)
	t.logger.Info(fmt.Sprintf("Start port connectivity detection, destination address:%s,timeout time:%v", addr, timeout.String()))
	_, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		t.logger.ErrorWithErr("telnet failed will change host status to offline", err)
		if err = t.checkFailed(*host); err != nil {
			return err
		}
		return nil
	}
	t.logger.Infof("%v Server connection successful\n", addr)
	if err = t.checkSuccess(*host); err != nil {
		return err
	}
	return nil
}

func (t *telnetHandler) checkFailed(host model.CMDBHost) error {
	return checkFailedHandler(t.factory, host)
}

func (t *telnetHandler) checkSuccess(host model.CMDBHost) error {
	return checkSuccessHandler(t.factory, host)
}

func (t *telnetHandler) HandlerErr(err error) {
	if t.queue.IsClosed() {
		t.logger.Info("telenet handler queue closed")
		return
	}
	if err != nil {
		t.logger.Error(fmt.Sprintf("telenet handler error %v", err))
	}
}
