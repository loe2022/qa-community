package api

import (
	"github.com/gin-gonic/gin"
	"qa-community/dao"
)

// 检查用户名，密码对不对
func checkUserPass(username string, password string) bool {
	datalist, err := dao.SelectFromUserName(username)
	if err != nil {
		return false
	}
	if datalist == nil {
		return false
	}
	data := datalist[0]
	if password == data.Password {
		return true
	}
	return false
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"ok":   true,
			"data": "no username or password",
		})
	}

	ok := checkUserPass(username, password)

	if ok {
		c.SetCookie("username", username, 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{
			"ok":   true,
			"data": "用户登录成功",
		})
	} else {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "登录失败或密码错误",
		})
	}
}
