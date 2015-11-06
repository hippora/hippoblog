package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}


func (c *AboutController) Get() {
	c.Data["IsAbout"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "about.html"
}
