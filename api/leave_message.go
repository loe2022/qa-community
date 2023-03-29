package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"qa-community/dao"
	"qa-community/util"
)

type MessageForm struct { //一条留言应该具有的内容
	Id       int    `form:"id"`       //本身的id
	ParentId int    `form:"parentid"` //父评论的id
	Name     string `form:"name"`     //发留言人的名字
	Message  string `form:"message"`  //留言内容
	IsDelete int    `form:"isdelete"`
}

func insertMessage(username string, content string, parentid int) bool {
	err := dao.InsertMes(username, content, parentid)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func LeaveMessage(c *gin.Context) {
	//获取cookie
	cuname, err1 := c.Cookie("username")
	resp, err := util.Cookie_check(cuname, err1)
	if err != nil {
		c.JSON(200, gin.H{"ok": false, "data": resp})
	}
	var messageform MessageForm
	//绑定参数，方便接收
	if err3 := c.ShouldBind(&messageform); err3 != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "bind data wrong",
		})
		return
	}

	if err != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "no such user",
		})
	} else {
		ok := insertMessage(messageform.Name, messageform.Message, messageform.ParentId)
		if ok {
			c.JSON(200, gin.H{
				"ok":   true,
				"data": "插入评论成功",
			})
		} else {
			c.JSON(200, gin.H{
				"ok":   false,
				"data": "插入评论失败",
			})
		}
	}
}
