package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"qa-community/dao"
)

func selectTableFromParentId(id int) ([]MessageForm, error) {
	var data MessageForm                                                    //定义了一个Message类型的结构体
	var list []MessageForm                                                  //结构体切片
	rows, err := dao.Db.Query("select * from message where parentid=?", id) //找parentid=114514的
	if err != nil {
		return []MessageForm{}, err
	}
	defer rows.Close()
	for rows.Next() { //循环取值，把值放进切片
		err := rows.Scan(&data.Id, &data.ParentId, &data.Message, &data.IsDelete)
		if err != nil {
			log.Fatal(err)
			return []MessageForm{}, err
		}
		list = append(list, data)
	}
	return list, nil //返回一个list结构体切片，里面所有信息的parentid都是114514
}

//根据id在数据库里面找到信息

func SelectTableFromId(id int) ([]MessageForm, error) {
	var data MessageForm
	var list []MessageForm
	rows, err := dao.Db.Query("select * from message where id=?", id)
	if err != nil {
		return []MessageForm{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&data.Id, &data.Name, &data.ParentId, &data.Message, &data.IsDelete)
		if err != nil {
			log.Fatal(err)
			return []MessageForm{}, err
		}
		list = append(list, data)
	}
	return list, nil
}

// 一开始的id是问题的id,一开始datalist里面只有问题，返回的时候是包括这个问题所有的回答
func get_from_parentid(id int, dataform []MessageForm) []MessageForm {
	datalist, err := selectTableFromParentId(id)
	if err != nil {
		return nil
	}
	if len(datalist) == 0 {
		return dataform
	}
	//datalist是所有父id是指定id的信息
	for _, i := range datalist {
		//遍历当前层的所有元素的id，向下查找元素
		//先将这一层的所有数据append进入切片
		for _, j := range datalist {
			dataform = append(dataform, j)
		}

		return get_from_parentid(i.Id, dataform)
	}
	return nil
}

// 返回了所有父id是114514的message的id列表，其实这些就是问题，返回了所有问题的id号
func get_all_root() (resp []int, err error) {
	datalist, err := selectTableFromParentId(114514)
	if err != nil {
		return nil, err
	}
	for _, i := range datalist {
		resp = append(resp, i.Id)
	}
	return resp, nil
}

func ShowMessage(c *gin.Context) {
	var table [][]MessageForm
	rootlist, err := get_all_root()
	if err != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "get table fail",
		})
	}
	//rootlist里面是所有问题的id
	for _, i := range rootlist {
		var datalist []MessageForm        //又声明了一个Message结构体切片变量
		self, err := SelectTableFromId(i) //返回id为i的信息，得到了问题
		if err != nil {
			c.JSON(200, gin.H{
				"ok":   false,
				"data": "gettable fail",
			})
		}
		datalist = append(datalist, self[0]) //把问题加进去
		datalist = get_from_parentid(i, datalist)
		if datalist == nil {
			c.JSON(200, gin.H{
				"ok":   false,
				"data": "datalist nil",
			})
		}
		table = append(table, datalist)
	}
	if err != nil {
		c.JSON(200, gin.H{
			"ok":   false,
			"data": "get table fail",
		})
	} else {
		c.JSON(200, gin.H{
			"ok":   true,
			"data": table,
		})
	}
}
