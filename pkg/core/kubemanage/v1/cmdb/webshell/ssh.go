package webshell

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/noovertime7/kubemanage/pkg/utils"

	"github.com/gorilla/websocket"
)

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

func (s *sshWsHandler) SetUp(w http.ResponseWriter, r *http.Request, cols, rows int) error {
	// 设置默认xterm窗口大小
	terminalConfig := Config{
		IpAddress:     "yunxue521.top",
		Port:          "22",
		UserName:      "root",
		Password:      "1qaz@WSXchenteng@",
		PrivateKey:    "",
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
		_ = ws.WriteMessage(websocket.BinaryMessage, []byte(err.Error()))
		_ = ws.Close()
		return err
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
				fmt.Printf("监听到resize信号%v\n", msg)
				terminal.SetWinSize(msg.Cols, msg.Rows)
			}
		}
	}()
	if err != nil {
		log.Printf("NewTerminalSession error %v\n", err)
		return err
	}

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
		if err := terminal.Close(); err != nil {
			return err
		}
		log.Printf("ws断开成功")
		return nil
	})

	go func() {
		for {
			// 每5秒
			timer := time.NewTimer(5 * time.Second)
			<-timer.C

			if stream.IsClosed() || terminal.IsClosed() {
				_ = timer.Stop()
				break
			}
			// 如果有 10 分钟没有数据流动，则断开连接
			if time.Now().Unix()-stream.UpdateAt.Unix() > 60*10 {
				stream.WsConn.WriteMessage(websocket.TextMessage, utils.Str2Bytes("检测到终端闲置，已断开连接...\r\n"))
				stream.WsConn.WriteMessage(websocket.BinaryMessage, utils.Str2Bytes("检测到终端闲置，已断开连接..."))
				stream.Close()
				_ = timer.Stop()
				break
			}
		}
	}()

	return nil
}
