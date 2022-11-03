package router

import (
	"context"
	"github.com/wonderivan/logger"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun(addr string) {
	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(10) * time.Second,
		MaxHeaderBytes: 1 << uint(20),
	}
	go func() {
		logger.Info("HttpServerRun:%s\n", addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			logger.Fatal("HttpServerRun:%s err:%v\n", addr, err)
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
