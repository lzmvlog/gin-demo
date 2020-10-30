package util

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init() {
	// 获取 MySQL 链接需要自己导入 _ "github.com/go-sql-driver/mysql" 😂
	database, err := sqlx.Open("mysql", "root:Root5683@@tcp(127.0.0.1:3306)/myschool")
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	// 关闭数据库连接
	//defer database.Close() // 注意这行代码要写在上面err判断的下面
}
