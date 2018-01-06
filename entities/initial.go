package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var eng *xorm.Engine

func init() {
	//创建一个数据库引擎
	db, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	eng = db
	db.Sync2(new(UserInfo))
}
