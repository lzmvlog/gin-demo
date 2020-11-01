package util

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB // 创建mysql连接

var Pool *redis.Pool //创建redis连接池

func Init() {
	// 获取 MySQL 链接需要自己导入 _ "github.com/go-sql-driver/mysql" 😂
	database, err := sqlx.Open("mysql", "root:Root5683@@tcp(127.0.0.1:3306)/myschool")
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		log.Println("open mysql failed,", err)
		return
	}

	Db = database
	// 关闭数据库连接
	//defer database.Close() // 注意这行代码要写在上面err判断的下面

	Pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按   需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
