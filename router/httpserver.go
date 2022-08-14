package router

import (
	"context"
	"github.com/noovertime7/kubemanage/config"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/public"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	public.PrintColor()
	dao.InitDB()
	go func() {
		logger.Info("启动容器websocket服务")
		http.HandleFunc("/ws/terminal", service.Terminal.WsHandler)
		logger.Info("websocket开始监听,地址:", config.WebSocketListenAddr)
		http.ListenAndServe(config.WebSocketListenAddr, nil)
	}()
	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           config.ListenAddr,
		Handler:        r,
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(10) * time.Second,
		MaxHeaderBytes: 1 << uint(20),
	}
	go func() {
		logger.Info("HttpServerRun:%s\n", config.ListenAddr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			logger.Fatal("HttpServerRun:%s err:%v\n", config.ListenAddr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		logger.Fatal("HttpServerStop err:%v\n", err)
	}
	logger.Info("HttpServerStop stopped\n")
}
