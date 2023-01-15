package cmdb

import (
	"context"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/runtime/queue"
)

// StartHostCheck 从数据库中不断查询放到queue中
func (h *hostService) StartHostCheck() {
	// TODO 考虑是否处理error
	hosts, _ := h.GetHostList(context.TODO(), model.CMDBHost{})
	for _, host := range hosts {
		if h.queue.IsClosed() {
			return
		}
		h.queue.Push(&queue.Event{Type: "AddHOST", Data: host})
	}
}
