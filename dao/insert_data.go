package dao

import "fmt"

func Insert_data(username string, password string) bool {
	var per user
	sqlStr1 := "select Name from user where username = ?"
	Db.QueryRow(sqlStr1, username).Scan(&per.Id, &per.Name, &per.Password)
	if per.Name != "" {
		return false
	} else {
		sqlStr := "insert into user (username,password) value(?,?)"
		_, err := Db.Exec(sqlStr, username, password)
		if err != nil {
			fmt.Printf("Insert data error:%v", err)
			return false
		}
		return true
	}
}
