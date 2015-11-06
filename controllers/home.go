package controllers

import (
	"github.com/astaxie/beego"
	"hippoblog/models"
	"github.com/astaxie/beego/context"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*v := c.GetSession("asta")
	if v == nil {
		c.SetSession("asta", int64(1))
	}
	c.Data["username"] = v.(int64)
	*/

	//	models.AddUser()
	users, err := models.GetUsers()
	if err != nil {
		beego.Error(err)
	}

	c.Data["users"] = users

	c.Data["IsHome"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "home.html"
}

func IsLogin(ctx *context.Context) bool {
	username := ctx.GetCookie("username")
	password := ctx.GetCookie("password")
	if beego.AppConfig.String("username") == username &&
	beego.AppConfig.String("password") == password {
		return true
	}
	return false
}