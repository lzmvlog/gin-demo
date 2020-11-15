package api

import (
	"gin/app/modle"
	"gin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 新增学生
/*
http://127.0.0.1:8080/student/save
{
   "id":"1",
   "name":"chenghao_",
   "age":30
}
*/
func Save(c *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu modle.Student
	if err := c.ShouldBindJSON(&stu); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := util.Db.Exec("insert into student(id,name,age) values(?,?,?)", stu.Id, stu.Name, stu.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, gin.H{"success": "添加成功"})
}

// 根据 id 查询
/*
http://127.0.0.1:8080/student/select?id=1
*/
func SelectById(c *gin.Context) {
	var stu []modle.Student
	// 获取查询参数
	id := c.Query("id")
	err := util.Db.Select(&stu, "select id,name,age from student where id =?", id)
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
func Update(c *gin.Context) {
	// 获取传递的参数 转换成你 struct
	var stu modle.Student
	if err := c.ShouldBindJSON(&stu); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := util.Db.Exec("update student set name=?,age=? where id=?", stu.Name, stu.Age, stu.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, gin.H{"success": "修改成功"})
}

// 删除
/*
http://127.0.0.1:8080/student/select?id=1
*/
func Delete(c *gin.Context) {
	// 获取查询参数
	id := c.Query("id")
	result, err := util.Db.Exec("delete from student where id =?", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": result})
}
