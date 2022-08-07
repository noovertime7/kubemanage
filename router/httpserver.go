package router

import (
	"context"
	"github.com/noovertime7/kubemanage/config"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/public"
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
