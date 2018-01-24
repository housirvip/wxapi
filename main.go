package main

import (
	_ "wxapi/routers"
	_ "wxapi/service"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
