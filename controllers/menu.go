package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"wxapi/models"
)

type MenuController struct {
	beego.Controller
}

func (this *MenuController) Create() {
	this.Ctx.WriteString("Create 成功\n"+models.CreateMenu())
}

func (this *MenuController) Delete() {
	this.Ctx.WriteString("Delete 成功\n"+models.DeleteMenu())
}

func (this *MenuController) Get() {
	this.Ctx.WriteString("Get 成功\n"+models.GetMenu())
}

func (this *MenuController) Post() {
	xml_bytes, _ := ioutil.ReadAll(this.Ctx.Request.Body)
	this.Ctx.Request.Body.Close()
	msg := new(models.Wxmsg)
	msg.Xml2Msg(xml_bytes)
	msg.Match()
	this.Ctx.WriteString(msg.Str2Msg())
}