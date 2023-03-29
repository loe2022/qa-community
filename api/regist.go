package api

import (
	"github.com/gin-gonic/gin"
	"qa-community/dao"
)

func regis(username string, password string) (bool, string) {
	ok := dao.Insert_data(username, password)
	if !ok {
		return ok, "插入失败，可能已经存在该用户"
	}
	return true, "插入成功"
}

func Register(c *gin.Context) {
	//获取提交的表单参数，用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "lack of content",
		})
	}
	//当用户名和密码都不为空的时候，执行插入
	ok, _ := regis(username, password)

	if ok {
		//设置cookie，名字叫username,值是用户提交的用户名，存活时间一小时，在所有页面都可以获取，本地
		c.SetCookie("username", username, 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{
			"ok":   true,
			"data": "用户创建成功",
		})
	} else {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "用户创建失败",
		})
	}
}
