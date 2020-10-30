package routers

import (
	"fmt"
	"gin/app/student"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
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

var Db *sqlx.DB

func init() {
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

func Router() *gin.Engine {
	r := gin.Default()
	s := r.Group("/student")
	{
		s.PUT("/save", save)
		s.GET("/select", selectById)
		s.POST("/update", update)
		s.DELETE("/delete", delete)
	}
	return r
}

// 新增学生
/*
http://127.0.0.1:8080/student/save
{
    "id":"1",
    "name":"chenghao_",
    "age":30
}
*/
func save(c *gin.Context) {
	// 获取传递的参数 转换成你 struct
	var stu student.Student
	if err := c.ShouldBindJSON(&stu); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := Db.Exec("insert into student(id,name,age) values(?,?,?)", stu.Id, stu.Name, stu.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, gin.H{"success": "添加成功"})
}

// 根据 id 查询
func selectById(c *gin.Context) {
	var stu []student.Student
	// 获取查询参数
	id := c.Query("id")
	err := Db.Select(&stu, "select id,name,age from student where id =?", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": stu})
}

// 修改
/*
http://127.0.0.1:8080/student/update
{
    "id":"1",
    "name":"chenghao_",
    "age":19
}
*/
func update(c *gin.Context) {
	// 获取传递的参数 转换成你 struct
	var stu student.Student
	if err := c.ShouldBindJSON(&stu); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := Db.Exec("update student set name=?,age=? where id=?", stu.Name, stu.Age, stu.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, gin.H{"success": "修改成功"})
}

// 删除
func delete(c *gin.Context) {
	// 获取查询参数
	id := c.Query("id")
	result, err := Db.Exec("delete from student where id =?", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": result})
}
