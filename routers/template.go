package routers

import (
	"gin/app/api"
	"gin/util"
	"github.com/gin-gonic/gin"
)

func Template() *gin.Engine {
	r := gin.Default()
	// 初始化数据连接
	util.Init()
	// LoadHTMLFiles加载一片HTML文件//并将结果与HTML渲染器关联。
	//r.LoadHTMLFiles()
	// LoadHTMLGlob加载由glob模式//标识的HTML文件，并将结果与HTML渲染器关联。
	r.LoadHTMLGlob("template/**/*")
	r.GET("/", api.Index)
	return r
}
