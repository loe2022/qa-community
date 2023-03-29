package api

import (
	"github.com/gin-gonic/gin"
	"qa-community/util"
)

func root(c *gin.Context) {
	cuname, err1 := c.Cookie("username")
	//只有当cookie username 和 password同时存在时
	resp, err := util.Cookie_check(cuname, err1)
	if err != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"ok":   true,
			"data": "你好，" + cuname,
		})
	}
}
