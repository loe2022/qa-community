package dao

import "log"

func Del_info(message_id int) error {
	sqlStr := "update message set isdelete = 1 where id=?"
	_, err := Db.Exec(sqlStr, message_id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
