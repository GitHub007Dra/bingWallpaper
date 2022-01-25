package main

import (
	"bingWallpaper/api/handle"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	conf, err := config.NewConfig("ini", "./conf/server.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile,`{"filename":"logs/bingWallPaper.log","level":7,"maxlines":10000,"daily":true,"maxdays":10}`)
	listenAddress := fmt.Sprintf(":%d", port)

	engine := gin.New()
	handle.Init(engine)
	server := &http.Server{
		Addr:    listenAddress,
		Handler: engine,
	}

	server.ListenAndServe()
}
