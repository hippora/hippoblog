package main

import (
	_ "hippoblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.EnableAdmin = true
	beego.Run()
}

