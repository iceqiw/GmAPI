package model

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var orm *xorm.Engine 

func InitDB() (*xorm.Engine , error) {
	engeine, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err == nil {
		orm = engeine
		orm.Sync(new(TrainSearch))
		return engeine, err
	}
	return nil, err
}