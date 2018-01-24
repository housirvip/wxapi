package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type User struct {
	Id       int64 `orm:"pk;auto"` //主键，自动增长
	Uid      string
	Name     string
	Sex      int
	Language string
	City     string
	Province string
	Country  string
	Time     time.Time
	Ban      bool
	Ex       int
}

func (this *User) GetUserByUid() bool {
	if this.Uid == NULL {
		panic(fmt.Errorf("没有给uid赋值"))
	}
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("uid", this.Uid).One(this)
	if err != nil {
		return false//错误不为nil表示没有找到此用户
	}
	return true//获取用户对象成功
}

func (this *User) AddUser() bool {
	if this.GetUserByUid() {
		return false//表示已经有此用户了，返回添加失败
	}
	o := orm.NewOrm()
	this.Ban = false
	this.Time = time.Now()
	this.Ex = 0
	_, err := o.Insert(this)
	if err != nil {
		panic(err)//报出此意外错误
		return false
	}
	return true
}

func (this *User) DelUser() {
	if this.GetUserByUid() {
		o := orm.NewOrm()
		_, err := o.Delete(this)
		if err != nil {
			panic(err)//删除失败，报出此意外错误
		}
	}
}