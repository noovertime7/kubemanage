package webshell

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	"github.com/noovertime7/kubemanage/pkg/logger"
)

// TerminalMessage 定义了终端和容器 shell 交互内容的格式 Operation 是操作类型
// Data 是具体数据内容 Rows和Cols 可以理解为终端的行数和列数，也就是宽、高
type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      int    `json:"rows"`
	Cols      int    `json:"cols"`
}

// TerminalSession 定义 TerminalSession 结构体，实现 PtyHandler 接口 // wsConn 是 websocket 连接 // sizeChan 用来定义终端输入和输出的宽和高 // doneChan 用于标记退出终端
type TerminalSession struct {
	WsConn      *websocket.Conn
	messageChan chan interface{}
	CreatedAt   time.Time
	UpdateAt    time.Time
	closed      bool
}

// NewTerminalSession 该方法用于升级 http 协议至 websocket，并new一个 TerminalSession 类型的对象返回
func NewTerminalSession(resizeCh chan interface{}, conn *websocket.Conn) (*TerminalSession, error) {
	session := &TerminalSession{
		messageChan: resizeCh,
		WsConn:      conn,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	return session, nil
}

// 用于读取web端的输入，接收web端输入的指令内容
func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.WsConn.ReadMessage()
	if err != nil {
		return copy(p, "\u0004"), err
	}
	// 反序列化
	var msg TerminalMessage
	if err = json.Unmarshal(message, &msg); err != nil {
		return copy(p, "\u0004"), err
	}
	t.UpdateAt = time.Now()
	// 逻辑判断
	switch msg.Operation {
	case "closePty":
		if err := t.WsConn.Close(); err != nil {
			return 0, err
		}
		logger.LG.Info("websocket close successful")
		return 0, nil
	// 如果是标准输入
	case "stdin":
		return copy(p, msg.Data), nil
	// 窗口调整大小
	case "resize":
		if msg.Cols > 0 && msg.Rows > 0 {
			logger.LG.Info("terminal start resize")
			t.messageChan <- msg
		}
		return copy(p, "resize success"), nil
	// ping	无内容交互
	case "ping":
		return copy(p, "pong\u0004"), nil
	default:
		return copy(p, "unknown message type\u0004"), fmt.Errorf("unknown message type")
	}
	return 0, err
}

// 写数据的方法，拿到 api-server 的返回内容，向web端输出
func (t *TerminalSession) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      string(p),
	})
	if err != nil {
		return 0, err
	}
	if err = t.WsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return 0, err
	}
	return len(p), nil
}

const closedMsg = "websocket timeout shutdown"

// Close 用于关闭websocket连接
func (t *TerminalSession) Close() error {
	if t.IsClosed() {
		return nil
	}
	if err := t.WsConn.Close(); err != nil {
		return err
	}
	t.closed = true
	return nil
}

// IsClosed 终端是否已关闭
func (t *TerminalSession) IsClosed() bool {
	return t.closed
}
