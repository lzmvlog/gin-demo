package main

import (
	"fmt"
	"gin/main/routers"
	"github.com/gin-gonic/gin"
)

//func main() {
//	// 1.创建路由
//	r := gin.Default()
//	// 2.绑定路由规则，执行的函数
//	// gin.Context，封装了request和response
//	r.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, "hello World!")
//	})
//	// 3.监听端口，默认在8080
//	// Run("里面不指定端口号默认为8080")
//	r.Run(":8000")
//}

//func main() {
//	r := gin.Default()
//	r.GET("/user/:name/*action", func(c *gin.Context) {
//		name := c.Param("name")
//		action := c.Param("action")
//		//截取/
//		action = strings.Trim(action, "/")
//		c.String(http.StatusOK, name+" is "+action)
//	})
//	//默认为监听8080端口
//	r.Run(":8000")
//}

//func main() {
//	r := gin.Default()
//	r.GET("/user", func(c *gin.Context) {
//		//指定默认值
//		//http://localhost:8080/user 才会打印出来默认的值
//		name := c.DefaultQuery("name", "枯藤")
//		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
//	})
//	r.Run()
//}

//func main() {
//	r := gin.Default()
//	r.POST("/form", func(c *gin.Context) {
//		types := c.DefaultPostForm("type", "post")
//		username := c.PostForm("username")
//		password := c.PostForm("userpassword")
//		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
//		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
//	})
//	r.Run()
//}

//func main() {
//	r := gin.Default()
//	r.LoadHTMLGlob("html/*")
//	r.GET("/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
//	})
//	r.Run()
//}

//func main() {
//	r := gin.Default()
//	r.LoadHTMLGlob("html/**/*")
//	r.GET("/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "我是测试", "address": "www.5lmh.com"})
//	})
//	r.Run()
//}

//func main() {
//	r := routers.SetupRouter()
//	if err := r.Run(); err != nil {
//		fmt.Println("startup service failed, err:%v\n", err)
//	}
//}

func main() {
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}