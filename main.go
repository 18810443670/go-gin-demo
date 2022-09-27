package main

import (
	"context"
	"fmt"
	"gin-1/Router"
	"gin-1/Services"
	"github.com/go-ini/ini"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var ConfigIni *ini.File

func init()  {
	fmt.Println("init ------------------ start")

	ConfigIni = Services.InitIni()

	//注册数据库
	Services.InitMysql(ConfigIni)

	fmt.Println("init ------------------ end")
}

func main() {
	fmt.Println("main ------------------ start")

	//加载路由
	router := Router.SetupRouter()

	//启动服务
	var port = ConfigIni.Section("service").Key("PORT").String()
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		// 监听请求
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅Shutdown（或重启）服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) // syscall.SIGKILL
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
	}
	log.Println("Server exiting")
}
