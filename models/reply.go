package models

import "github.com/astaxie/beego/orm"

type Reply struct {
	Id      int64  `orm:"pk;auto"` //主键，自动增长
	Key     int
	Content string
}

func (this *Reply) GetReplyByKey(key string) bool {
	o := orm.NewOrm()
	err := o.QueryTable("reply").Filter("key", key).One(this)
	if nil != err {
		return false
	}
	return true
}

func (this *Reply) AddReply(key, content string) {
	o := orm.NewOrm()
	_, err := o.Insert(this)
	if err != nil {
		panic(err)
	}
}