package routers

import (
	"gin/app/api"
	"gin/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//func helloHandler(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Hello www.topgoer.com!",
//	})
//}
//
//func SetupRouter() *gin.Engine {
//	r := gin.Default()
//	r.GET("/topgoer", helloHandler)
//	return r
//}

func Router() *gin.Engine {
	r := gin.Default()
	// 初始化数据连接
	util.Init()
	s := r.Group("/student")
	{
		s.PUT("/save", api.Save)
		s.GET("/select", api.SelectById)
		s.POST("/update", api.Update)
		s.DELETE("/delete", api.Delete)
	}
	return r
}
