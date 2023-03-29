package main

import (
	"log"
	"qa-community/api"
	"qa-community/dao"
)

func main() {
	//初始化数据库
	err := dao.InitSql()
	if err != nil {
		log.Fatal(err)
		return
	}
	//初始化路由
	api.InitRouter()
}
