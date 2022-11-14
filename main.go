package main

import (
	"context"
	"github.com/yujian0213/self-web/framework/gin"
	"github.com/yujian0213/self-web/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	core.Use(gin.Recovery())
	core.Use(middleware.Coast())
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr: ":8888",
	}
	go func() {
		_ = server.ListenAndServe()
	}()
	//当前goroutine等待信号量
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	 <- quit
	 timeoutCtx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	 defer cancel()
	 if err := server.Shutdown(timeoutCtx);err != nil {
	 	log.Fatal("server shutdown",err)
	 }
}