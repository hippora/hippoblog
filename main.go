package main

import (
	_ "github.com/hippora/hippoblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.Listen.AdminEnable = true
	beego.Run()
}

