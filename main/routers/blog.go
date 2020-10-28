package routers

import "github.com/gin-gonic/gin"

func LoadShop(e *gin.Engine)  {
	e.GET("/hello", helloHandler)
}
