package dao

import "log"

func Update_info(content string, message_id int) error {
	sqlStr := "update message set message = ? where id = ? "
	_, err := Db.Exec(sqlStr, content, message_id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
