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
	logs.Info("=============")
	data := map[string]interface{}{"status": "ok"}
	c.JSON(http.StatusOK, data)
}

type Images struct {
	Url string `json:"url"`
}

type bingWallPaper struct {

	Images []Images `json:"images"`

}
func bingWallpaper(c *gin.Context) {
	logs.Info("====someone get ====")

	logs.Info(c.Request.RemoteAddr +" "+c.Request.Header.Get("User-Agent") )

	get := util.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")

	var config bingWallPaper
	err := json.Unmarshal([]byte(get), &config)
	if err !=nil{
		fmt.Println(err)
	}
	today:="http://cn.bing.com"+config.Images[0].Url

	c.Redirect(http.StatusFound, today)
}
