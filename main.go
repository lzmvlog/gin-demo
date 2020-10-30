package main

import (
	"gin/routers"
	_ "github.com/go-sql-driver/mysql"
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

//func main() {
//	r := gin.Default()
//	routers.LoadBlog(r)
//	routers.LoadShop(r)
//	if err := r.Run(); err != nil {
//		fmt.Println("startup service failed, err:%v\n", err)
//	}
//}

//// 定义接收数据的结构体
//type Login struct {
//	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
//	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
//	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
//}
//
//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// JSON绑定
//	r.POST("loginJSON", func(c *gin.Context) {
//		// 声明接收的变量
//		var json Login
//		// 将request的body中的数据，自动按照json格式解析到结构体
//		if err := c.ShouldBindJSON(&json); err != nil {
//			// 返回错误信息
//			// gin.H封装了生成json数据的工具
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		// 判断用户名密码是否正确
//		if json.User != "root" || json.Pssword != "admin" {
//			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"status": "200"})
//	})
//	r.Run(":8000")
//}

//// 多种响应方式
//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// 1.json
//	r.GET("/someJSON", func(c *gin.Context) {
//		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
//	})
//	// 2. 结构体响应
//	r.GET("/someStruct", func(c *gin.Context) {
//		var msg struct {
//			Name    string
//			Message string
//			Number  int
//		}
//		msg.Name = "root"
//		msg.Message = "message"
//		msg.Number = 123
//		c.JSON(200, msg)
//	})
//	// 3.XML
//	r.GET("/someXML", func(c *gin.Context) {
//		c.XML(200, gin.H{"message": "abc"})
//	})
//	// 4.YAML响应
//	r.GET("/someYAML", func(c *gin.Context) {
//		c.YAML(200, gin.H{"name": "zhangsan"})
//	})
//	// 5.protobuf格式,谷歌开发的高效存储读取的工具
//	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
//	r.GET("/someProtoBuf", func(c *gin.Context) {
//		reps := []int64{int64(1), int64(2)}
//		// 定义数据
//		label := "label"
//		// 传protobuf格式数据
//		data := &protoexample.Test{
//			Label: &label,
//			Reps:  reps,
//		}
//		c.ProtoBuf(200, data)
//	})
//
//	r.Run(":8000")
//}

//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// 1.异步
//	r.GET("/long_async", func(c *gin.Context) {
//		// 需要搞一个副本
//		copyContext := c.Copy()
//		// 异步处理
//		go func() {
//			time.Sleep(3 * time.Second)
//			log.Println("异步执行：" + copyContext.Request.URL.Path)
//		}()
//	})
//	// 2.同步
//	r.GET("/long_sync", func(c *gin.Context) {
//		time.Sleep(3 * time.Second)
//		log.Println("同步执行：" + c.Request.URL.Path)
//	})
//
//	r.Run(":8000")
//}

//type Student struct {
//	Id      string `orm:"id,primary" json:"id"`      //
//	Name    string `orm:"name"       json:"name"`    //
//	Age     int    `orm:"age"        json:"age"`     //
//}
//
//var Db *sqlx.DB
//
//func init() {
//	// 获取 MySQL 链接需要自己导入 _ "github.com/go-sql-driver/mysql" 😂
//
//	database, err := sqlx.Open("mysql", "root:Root5683@@tcp(127.0.0.1:3306)/myschool")
//	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//
//	Db = database
//	defer database.Close() // 注意这行代码要写在上面err判断的下面
//}
//
//func main() {
//	var stu []Student
//	err := Db.Select(&stu, "select id, name, age from student where id=?", "27afbc17ebb944a39efb094c7de15355")
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		return
//	}
//
//	fmt.Println("select succ:", stu)
//}

//func main() {
//	c, err := redis.Dial("tcp", "127.0.0.1:6379")
//	if err != nil {
//		fmt.Println("conn redis failed,", err)
//		return
//	}
//
//	fmt.Println("redis conn success")
//
//	defer c.Close()
//}

func main() {
	//r := gin.Default()
	// 加载路由
	r := routers.Router()
	r.Run()
}
