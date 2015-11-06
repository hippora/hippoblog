package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	c.Ctx.SetCookie("username", "", -1, "/")
	c.Ctx.SetCookie("password", "", -1, "/")
	c.Redirect("/",301)
}
