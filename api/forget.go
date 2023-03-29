package api

import (
	"github.com/gin-gonic/gin"
	"qa-community/dao"
)

func Forget(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "lack of content",
		})
	}
	datalist, err := dao.SelectFromUserName(username)
	if err != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "data get wrong,user may not exist",
		})
	} else {
		data := datalist[0]
		c.JSON(200, gin.H{
			"ok":   true,
			"info": "你的密码是:",
			"data": data.Password,
		})
	}
}
