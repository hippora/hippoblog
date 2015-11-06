package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}


func (c *LoginController) Get() {
	if IsLogin(c.Ctx) {
		c.Redirect("/", 301)
		return
	}
	c.TplNames = "login.html"
}

func (c *LoginController) Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	if beego.AppConfig.String("username") == username &&
	beego.AppConfig.String("password") == password {
		maxAge := 0
		if c.Input().Get("autologin") == "on" {
			maxAge = 1 << 31 - 1
		}
		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}
	c.Redirect("/", 301)
}