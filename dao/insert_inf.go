package dao

func InsertMes(username string, content string, parentid int) error {
	sqlStr := "insert into message (username,message,parentid) value(?,?,?)"
	_, err := Db.Exec(sqlStr, username, content, parentid)
	if err != nil {
		return err
	}
	return nil
}
