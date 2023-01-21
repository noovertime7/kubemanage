package webshell

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/runtime"
	"github.com/noovertime7/kubemanage/runtime/wait"
)

type HostConnectionInfo struct {
	Address            string
	Port               uint
	UserName, Password string
	UsePrivateKey      bool
	PrivateKey         string
}

var (
	UpGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024 * 1024 * 10,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type RecordData struct {
	Event string  `json:"event"` // 输入输出事件
	Time  float64 `json:"time"`  // 时间差
	Data  []byte  `json:"data"`  // 数据
}

type Meta struct {
	TERM      string
	Width     int
	Height    int
	UserName  string
	ConnectId string
	HostId    uint
	HostName  string
}

var SSHWsHandler = NewSSHWsHandler()

func NewSSHWsHandler() *sshWsHandler {
	return &sshWsHandler{}
}

type sshWsHandler struct {
	sync.RWMutex
	Terminal    *Terminal       // ssh客户端
	Conn        *websocket.Conn // socket 连接
	messageType int             // 发送的数据类型
	recorder    []*RecordData   // 操作记录
	CreatedAt   time.Time       // 创建时间
	UpdatedAt   time.Time       // 最新的更新时间
	Meta        Meta            // 元信息
	written     bool            // 是否已写入记录, 一个流只允许写入一次
}

func (s *sshWsHandler) SetUp(w http.ResponseWriter, r *http.Request, info *HostConnectionInfo, cols, rows int) error {
	// 获取主机信息
	// 设置默认xterm窗口大小
	terminalConfig := Config{
		IpAddress:     info.Address,
		Port:          strconv.Itoa(int(info.Port)),
		UserName:      info.UserName,
		Password:      info.Password,
		PrivateKey:    info.PrivateKey,
		KeyPassphrase: "",
		Width:         cols,
		Height:        rows,
	}

	ws, err := UpGrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	terminal, err := NewTerminal(terminalConfig)
	if err != nil {
		if err = sendMsg(ws, err.Error()); err != nil {
			return err
		}
		return ws.Close()
	}
	resizeCh := make(chan interface{}, 10)
	stream, err := NewTerminalSession(resizeCh, ws)
	go func() {
		for {
			if terminal.IsClosed() {
				return
			}
			select {
			case data := <-resizeCh:
				msg := data.(TerminalMessage)
				terminal.SetWinSize(msg.Cols, msg.Rows)
			}
		}
	}()
	if err != nil {
		return err
	}

	logger.LG.Info(fmt.Sprintf("start ssh connection to %s", info.Address))
	err = terminal.Connect(stream, stream, stream)

	if err != nil {
		_ = ws.WriteMessage(websocket.BinaryMessage, []byte(err.Error()))
		_ = ws.Close()
		return err
	}

	// 断开ws和ssh的操作
	stream.WsConn.SetCloseHandler(func(code int, text string) error {
		if err := stream.Close(); err != nil {
			return err
		}
		if err = terminal.Close(); err != nil {
			return err
		}
		return nil
	})

	go func() {
		// 每5秒检测一次UpdateAt 超时退出
		err = wait.PollImmediateUntil(5*time.Second, func() (done bool, err error) {
			if stream.IsClosed() || terminal.IsClosed() {
				return true, err
			}
			// 5分钟超时时间
			if time.Now().Unix()-stream.UpdateAt.Unix() > 60*5 {
				//超时退出发送消息
				if err = sendMsg(stream.WsConn, "\r\n检测到终端闲置，已断开连接..."); err != nil {
					return false, err
				}

				err = stream.Close()
				if err != nil {
					return false, err
				}

				return true, err
			}
			return false, nil
		}, runtime.SystemContext.Done())

		if err != nil {
			logger.LG.Warn("websocket timeout timer error", zap.Error(err))
			return
		}
	}()

	return nil
}

func sendMsg(socket *websocket.Conn, data string) error {
	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      data,
	})
	if err != nil {
		return err
	}
	return socket.WriteMessage(websocket.TextMessage, msg)
}
