package main

import (
	"bingWallpaper/api/handle"
	"bingWallpaper/util"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

//# 不使用这个库了 github.com/reujab/wallpaper v0.0.0-20210630195606-5f9f655b3740

func main() {

	/*	代码注释是做的壁纸服务
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
	*/

	//获取今天的壁纸url
	getResponse := util.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
	var config handle.BingWallPaper
	err := json.Unmarshal([]byte(getResponse), &config)
	if err != nil {
		fmt.Println(err)
	}
	todayUrl := "http://cn.bing.com" + config.Images[0].Url

	//判断目录是否存在
	usr, err := user.Current()
	wallPaperdir := filepath.Join(usr.HomeDir, "Pictures", "wallpaper")
	_, exist := os.Stat(wallPaperdir)
	if os.IsNotExist(exist) {
		os.Mkdir(wallPaperdir, os.ModePerm)
	}
	timeStr := time.Now().Format("2006-01-02")
	join := filepath.Join(usr.HomeDir, "Pictures", "wallpaper", timeStr+".jpg")

	print(join)
	//将壁纸下载到指定目录下
	s, err := util.DownloadImageToFile(todayUrl, join)
	if err != nil {
		print(s)
	}

	//设置壁纸
	errw := util.SetFromURL(todayUrl)
	if errw != nil {
		print("====")
	}

	//删除本地缓存
	cacheDir, _ := util.GetCacheDir()

	file, err := os.Create(filepath.Join(cacheDir, "wallpaper"))
	if err != nil {
		return
	}
	os.Remove(file.Name())

}
