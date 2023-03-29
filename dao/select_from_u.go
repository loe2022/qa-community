package dao

import "log"

type user struct {
	Id       int
	Name     string
	Password string
}

func SelectFromUserName(username string) ([]user, error) {
	var data user   //data是一个user类型的结构体
	var list []user //list是一个结构体切片
	sqlStr := "select * from user where username=?"
	rows, err := Db.Query(sqlStr, username)
	if err != nil {
		return []user{}, err
	}
	defer rows.Close() //要释放连接池
	for rows.Next() {
		// row.scan 必须按照先后顺序 &获取数据
		err := rows.Scan(&data.Id, &data.Name, &data.Password)
		if err != nil {
			log.Println(err)
			return []user{}, err
		}
		list = append(list, data)
	}
	return list, nil
}
