package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
	"time"
)

// TerminalMessage定义了终端和容器shell交互内容的格式 //Operation是操作类型
// Data是具体数据内容 //Rows和Cols可以理解为终端的行数和列数，也就是宽、高

type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      uint16 `json:"rows"`
	Cols      uint16 `json:"cols"`
}

// 初始化一个websocket.Upgrader类型的对象，用于http协议升级为websocket协议
var upgrader = func() websocket.Upgrader {
	upgrader := websocket.Upgrader{}
	upgrader.HandshakeTimeout = time.Second * 2
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return upgrader
}()

// TerminalSession 定义TerminalSession结构体，实现PtyHandler接口 //wsConn是websocket连接 //sizeChan用来定义终端输入和输出的宽和高 //doneChan用于标记退出终端
type TerminalSession struct {
	wsConn   *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

// Terminal 定义Terminal全局变量
var Terminal terminal

// 定义terminal结构体
type terminal struct{}

// WsHandler 定义websocket的handler方法
func (t *terminal) WsHandler(w http.ResponseWriter, r *http.Request) {
	// 加载k8s配置
	conf, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\18495\\.kube\\config")
	if err != nil {
		logger.Error("初始化kubernetes配置失败,错误信息," + err.Error())
	}
	// 解析form入参，获取namespace、podName、containerName参数
	// 如果解析失败
	if err := r.ParseForm(); err != nil {
		logger.Error("解析参数失败,错误信息," + err.Error())
		return
	}
	// 如果解析成功
	namespace := r.Form.Get("namespace")
	podName := r.Form.Get("pod_name")
	containerName := r.Form.Get("container_name")
	logger.Info("exec pod: %s, container: %s, namespace: %s\n", podName, containerName, namespace)

	// new一个TerminalSession类型的pty实例
	pty, err := NewTerminalSession(w, r, nil)
	if err != nil {
		logger.Error("get pty failed: %v\n", err)
		return
	}
	// 处理关闭
	defer func() {
		logger.Info("close session successfully!")
		pty.Close()
	}()
	/* 初始化pod所在的corev1资源组
	PodExecOptions struct 包括Container stdout stdout  Command 等结构
	scheme.ParameterCodec 应该是pod 的GVK （GroupVersion & Kind）之类的
	URL长相:
	https://192.168.1.11:6443/api/v1/namespaces/default/pods/nginx-wf2-778d88d7c-7rmsk/exec?command=%2Fbin%2Fbash&container=nginx-wf2&stderr=true&stdin=true&stdout=true&tty=true
	*/
	// 组装POST请求
	req := kube.K8s.ClientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   []string{"/bin/bash"},
			Stderr:    true,
			Stdin:     true,
			Stdout:    true,
			TTY:       true,
		}, scheme.ParameterCodec)
	logger.Info("exec post request url: ", req)

	// remotecommand 主要实现了http 转 SPDY 添加X-Stream-Protocol-Version相关header 并发送请求
	executor, err := remotecommand.NewSPDYExecutor(conf, "POST", req.URL())
	if err != nil {
		logger.Error("建立SPDY连接失败," + err.Error())
		return
	}
	// 与kubelet建立stream连接
	err = executor.Stream(remotecommand.StreamOptions{
		Stdout:            pty,
		Stdin:             pty,
		Stderr:            pty,
		TerminalSizeQueue: pty,
		Tty:               true,
	})
	if err != nil {
		logger.Error("exec pod command failed," + err.Error())
		// 将报错返回给web端
		pty.Write([]byte("exec pod command failed," + err.Error()))
		// 标记关闭terminal
		pty.Done()
	}
}

// NewTerminalSession 该方法用于升级http协议至websocket，并new一个TerminalSession类型的对象返回
func NewTerminalSession(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*TerminalSession, error) {
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	session := &TerminalSession{
		wsConn:   conn,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
	return session, nil
}

// 用于读取web端的输入，接收web端输入的指令内容
func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.wsConn.ReadMessage()
	if err != nil {
		logger.Error(errors.New("读取parse信息失败,错误信息," + err.Error()))
		return copy(p, "\u0004"), err
	}
	// 反序列化
	var msg TerminalMessage
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		logger.Error(errors.New("读取parse信息失败,错误信息," + err.Error()))
		// return 0, nil
		return copy(p, "\u0004"), err
	}
	// 逻辑判断
	switch msg.Operation {
	// 如果是标准输入
	case "stdin":
		return copy(p, msg.Data), nil
	// 窗口调整大小
	case "resize":
		t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	// ping	无内容交互
	case "ping":
		return 0, nil
	default:
		logger.Info(errors.New("无法确认的message类型,当前类型为 " + msg.Operation))
		// return 0, nil
		return copy(p, "\u0004"), fmt.Errorf("unknown message type '%s'",
			msg.Operation)
	}
}

// 写数据的方法，拿到apiserver的返回内容，向web端输出
func (t *TerminalSession) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      string(p),
	})
	if err != nil {
		logger.Error("write parse message err: %v", err)
		return 0, err
	}
	if err := t.wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
		logger.Info("write message err: %v", err)
		return 0, err
	}
	return len(p), nil
}

// Done 标记关闭doneChan,关闭后触发退出终端
func (t *TerminalSession) Done() {
	close(t.doneChan)
}

// Close 用于关闭websocket连接
func (t *TerminalSession) Close() error {
	return t.wsConn.Close()
}

// Next 获取web端是否resize,以及是否退出终端
func (t *TerminalSession) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}
