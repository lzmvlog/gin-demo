package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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

//func main() {
//	//r := gin.Default()
//	// 加载路由
//	r := routers.Router()
//	r.Run()
//}

//func main() {
//	c, err := redis.Dial("tcp", "localhost:6379")
//	if err != nil {
//		fmt.Println("conn redis failed,", err)
//		return
//	}
//
//	defer c.Close()
//	_, err = c.Do("Set", "abc", 100)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	r, err := redis.Int(c.Do("Get", "abc"))
//	if err != nil {
//		fmt.Println("get abc failed,", err)
//		return
//	}
//
//	fmt.Println(r)
//
//	c, err := redis.Dial("tcp", "localhost:6379")
//	if err != nil {
//		log.Println("conn redis failed,", err)
//		return
//	}
//
//	_, err = c.Do("Set", "name", "Shao Jie")
//	if err != nil {
//		log.Println(err.Error())
//		return
//	}
//
//	name, err := redis.String(c.Do("Get", "name"))
//	if err != nil {
//		log.Println(err.Error())
//		return
//	}
//	log.Println(name)
//}

//func main() {
//	// 获取操作redis的链接
//	c, err := redis.Dial("tcp", "localhost:6379")
//	if err != nil {
//		log.Println("conn redis failed,", err)
//		return
//	}
//
//	//_, err = c.Do("MSet", "name", "shaojie_", "age", 22)
//	//if err != nil {
//	//	log.Println(err.Error())
//	//	return
//	//}
//	//
//	//result, err := redis.Strings(c.Do("MGet", "name", "age"))
//	//if err != nil {
//	//	log.Println("get name failed,", err)
//	//	return
//	//}
//	//
//	//for _, v := range result {
//	//	log.Println(v)
//	//}
//
//	// 设置一个key为name 值为shaojie_
//	//_, err = c.Do("Set", "name", "shaojie_")
//	//if err != nil {
//	//	log.Println(err.Error())
//	//	return
//	//}
//	//
//	//// 设置过期时间 10 为秒
//	//_, err = c.Do("expire", "name", 10)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//
//	//// 取值 取一个key为name的值
//	//name, err := redis.String(c.Do("Get", "name"))
//	//if err != nil {
//	//	log.Println(err.Error())
//	//	return
//	//}
//	//log.Println(name)
//	//
//	//// 休眠十秒
//	//time.Sleep(10 * time.Second)
//	//
//	//// 取值  十秒取一个key为name的值
//	//name1, err := redis.String(c.Do("Get", "name"))
//	//if err != nil {
//	//	log.Println(err.Error())
//	//	return
//	//}
//	//log.Println(name1)
//
//	// 设置一个key为namelist 值为shaojie_ 、chenghao_
//	//_, err = c.Do("lpush", "namelist", "shaojie_","chenghao_")
//	//if err != nil {
//	//	log.Println(err.Error())
//	//	return
//	//}
//	//
//	//// 取值 获取一个值为 namelist的值 获取这个值之后会删除这个值
//	//namelist, err := redis.String(c.Do("lpop", "namelist"))
//	//if err != nil {
//	//	log.Println("get namelist failed,", err)
//	//	return
//	//}
//	//
//	//log.Println(namelist)
//
//	// 设置 shaojie_ 的年龄age为21
//	_, err = c.Do("HSet", "shaojie_", "age", 21)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	// 获取 shaojie_ 的年龄age
//	age, err := redis.Int(c.Do("HGet", "shaojie_", "age"))
//	if err != nil {
//		fmt.Println("get shaojie_'s age failed,", err)
//		return
//	}
//
//	log.Println(age)
//}

//var pool *redis.Pool //创建redis连接池
//
//func init() {
//	pool = &redis.Pool{ //实例化一个连接池
//		MaxIdle: 16, //最初的连接数量
//		// MaxActive:1000000,    //最大连接数量
//		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按   需分配
//		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
//		Dial: func() (redis.Conn, error) { //要连接的redis数据库
//			return redis.Dial("tcp", "localhost:6379")
//		},
//	}
//}

//func main() {
//	util.Init()
//	c := util.Pool.Get() //从连接池，取一个链接
//	defer c.Close()      //函数运行结束 ，把连接放回连接池
//
//	//_, err := c.Do("Set", "abc", 200)
//	//if err != nil {
//	//	log.Println(err)
//	//	return
//	//}
//
//	r, err := redis.Int(c.Do("HGet", "shaojie_", "age"))
//	if err != nil {
//		fmt.Println("get abc faild :", err)
//		return
//	}
//	log.Println(r)
//	util.Pool.Close() //关闭连接池
//}

//// 定义中间
//func MiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//		fmt.Println("中间件开始执行了")
//		// 设置变量到Context的key中，可以通过Get()取
//		c.Set("request", "中间件")
//		status := c.Writer.Status()
//		fmt.Println("中间件执行完毕", status)
//		t2 := time.Since(t)
//		fmt.Println("time:", t2)
//	}
//}
/*
http://127.0.0.1:8080/intermediate
*/
//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// 注册中间件
//	r.Use(MiddleWare())
//	// {}为了代码规范
//	{
//		r.GET("/intermediate", func(c *gin.Context) {
//			// 取值
//			req, _ := c.Get("request")
//			fmt.Println("request:", req)
//			// 页面接收
//			c.JSON(200, gin.H{"request": req})
//		})
//
//	}
//	r.Run()
//}

// 定义中间
//func MiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//		fmt.Println("中间件开始执行了")
//		// 设置变量到Context的key中，可以通过Get()取
//		c.Set("request", "中间件")
//		// 执行函数
//		c.Next()
//		// 中间件执行完后续的一些事情
//		status := c.Writer.Status()
//		fmt.Println("中间件执行完毕", status)
//		// 从 t 开始经过得时间
//		t2 := time.Since(t)
//		fmt.Println("time:", t2)
//	}
//}
//
//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// 注册中间件
//	r.Use(MiddleWare())
//	// {}为了代码规范
//	{
//		r.GET("/intermediate", func(c *gin.Context) {
//			// 取值
//			req, _ := c.Get("request")
//			fmt.Println("request:", req)
//			// 页面接收
//			c.JSON(200, gin.H{"request": req})
//		})
//	}
//	r.Run()
//}

// 定义中间
//func MiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//		fmt.Println("中间件开始执行了")
//		// 设置变量到Context的key中，可以通过Get()取
//		c.Set("request", "中间件")
//		// 执行函数
//		c.Next()
//		// 中间件执行完后续的一些事情
//		status := c.Writer.Status()
//		fmt.Println("中间件执行完毕", status)
//		t2 := time.Since(t)
//		fmt.Println("time:", t2)
//	}
//}
//
//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	//局部中间键使用
//	r.GET("/intermediate", MiddleWare(), func(c *gin.Context) {
//		// 取值
//		req, _ := c.Get("request")
//		fmt.Println("request:", req)
//		// 页面接收
//		c.JSON(200, gin.H{"request": req})
//	})
//	r.Run()
//}

// 定义中间
//func myTime(c *gin.Context) {
//	start := time.Now()
//	c.Next()
//	// 统计时间
//	since := time.Since(start)
//	fmt.Println("程序用时：", since)
//}
//
//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// 注册中间件
//	r.Use(myTime)
//	// {}为了代码规范
//	shoppingGroup := r.Group("/shopping")
//	{
//		shoppingGroup.GET("/index", shopIndexHandler)
//		shoppingGroup.GET("/home", shopHomeHandler)
//	}
//	r.Run()
//}
//
//func shopIndexHandler(c *gin.Context) {
//	time.Sleep(5 * time.Second)
//}
//
//func shopHomeHandler(c *gin.Context) {
//	time.Sleep(3 * time.Second)
//}
//func main() {
//	// 1.创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// 服务端要给客户端cookie
//	r.GET("cookie", func(c *gin.Context) {
//		// 获取客户端是否携带cookie
//		cookie, err := c.Cookie("key_cookie")
//		if err != nil {
//			cookie = "NotSet"
//			// 给客户端设置cookie
//			//  maxAge int, 单位为秒
//			// path,cookie所在目录
//			// domain string,域名
//			//   secure 是否智能通过https访问
//			// httpOnly bool  是否允许别人通过js获取自己的cookie
//			c.SetCookie("key_cookie", "value_cookie", 60, "/",
//				"localhost", false, true)
//		}
//		fmt.Printf("cookie的值是： %s\n", cookie)
//	})
//	r.Run()
//}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	// 1.创建路由
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run()
}
