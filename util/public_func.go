package util

import (
	"github.com/gin-gonic/gin"
	"qa-community/dao"
)

func Cookie_check(username string, err1 error) (gin.H, error) {
	//第一种情况，没有cookie，这里的err1是获取cookie时返回的那一个err1
	if err1 != nil {
		return gin.H{
			"ok":   false,
			"data": "no cookie",
		}, err1
	} else {
		//第二种情况：提供了cookie
		datalist, err := dao.SelectFromUserName(username)
		if err != nil {
			return gin.H{"ok": false, "data": "no such username"}, err
		}
		data := datalist[0]
		if data.Name == username {
			return gin.H{"ok": true, "data": "check_success"}, nil
		}
	}
	return nil, err1
}
