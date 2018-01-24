package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Log struct {
	Id      int64 `orm:"pk;auto"`
	Content string
	Time    time.Time
}

func (this *Log) AddLog() {
	o:=orm.NewOrm()
	this.Time=time.Now()
	_,err:=o.Insert(this)
	if err!=nil {
		panic(err)
	}
}
