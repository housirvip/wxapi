package routers

import (
	"wxapi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.MenuController{})
}
