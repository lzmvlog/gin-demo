package api

import (
	"gin/app/modle"
	"gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	var stu []modle.Student
	err := util.Db.Select(&stu, "select id,name,age from student where id =?", "545289da9a8740b3a7cba5554e5f3eb7")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.html", stu)
}
