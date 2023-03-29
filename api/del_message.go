package api

import (
	"github.com/gin-gonic/gin"
	"qa-community/dao"
	"qa-community/util"
)

func update_isdelete(message_id int) bool {
	err := dao.Del_info(message_id)
	if err != nil {
		return false
	}
	return true
}

func DeleteMessage(c *gin.Context) {
	cuname, err1 := c.Cookie("username")
	resp, err := util.Cookie_check(cuname, err1)
	if err != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": resp,
		})
	}
	var message MessageForm
	if err3 := c.ShouldBind(&message); err3 != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "bind data wrong",
		})
		return
	}
	ok := update_isdelete(message.Id)
	if ok {
		c.JSON(200, gin.H{
			"ok":   true,
			"data": "delete successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "sql wrong",
		})
	}
}
