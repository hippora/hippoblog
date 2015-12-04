package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hippora/hippoblog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	articles, err := models.GetArticles()
	if err != nil {
		beego.Error(err)
	}
	c.Data["articles"] = articles

	catagories,err := models.GetCatagories()
	if err!= nil {
		beego.Error(err)
	}
	c.Data["catagories"] = catagories

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