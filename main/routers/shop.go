package routers

import "github.com/gin-gonic/gin"

func LoadBlog(e *gin.Engine)  {
	e.GET("/helloblog", helloHandler)
}
