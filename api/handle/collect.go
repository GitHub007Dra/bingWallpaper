package handle

import (
	"bingWallpaper/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	data := map[string]interface{}{"status": "ok"}
	fmt.Print("==HealthCheck===")
	c.JSON(http.StatusOK, data)
}

func haha(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	c.JSON(http.StatusCreated, nil)
	fmt.Println(json)
}

type Image struct {
	Url string `json:"url"`
}

type BingWallPaper struct {
	Images []Image `json:"images"`
}

func bingWallpaper(c *gin.Context) {
	logs.Info("====someone get ====")

	logs.Info(c.Request.RemoteAddr + " " + c.Request.Header.Get("User-Agent"))

	get := util.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")

	var config BingWallPaper
	err := json.Unmarshal([]byte(get), &config)
	if err != nil {
		fmt.Println(err)
	}
	today := "http://cn.bing.com" + config.Images[0].Url

	c.Redirect(http.StatusFound, today)
}
