package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"wxapi/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Ctx.WriteString(this.Input().Get("echostr"))
}

func (this *IndexController) Post() {
	xml_bytes, _ := ioutil.ReadAll(this.Ctx.Request.Body)
	this.Ctx.Request.Body.Close()
	msg := new(models.Wxmsg)
	msg.Xml2Msg(xml_bytes)
	msg.Match()
	this.Ctx.WriteString(msg.Str2Msg())
}