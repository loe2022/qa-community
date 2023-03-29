package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB //声明全局变量连接池，方便使用

func InitSql() (err error) {
	//数据库信息 用户名:密码@tcp(地址:端口)/数据库的名字
	dsn := "root:2019219033@tcp(127.0.0.1:3306)/community"
	//mysql驱动获取连接池
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect format error : %v", err)
		return
	}
	err = Db.Ping()
	if err != nil {
		fmt.Printf("mysql connect error:%v", err)
		return
	}
	return nil
}
