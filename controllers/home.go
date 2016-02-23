package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hippora/hippoblog/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	page,err := strconv.ParseInt(c.Input().Get("p"),10,64)
	if err != nil ||  page < 1 {
		page = 1
	}
	awcs, err := models.GetArticleWithCataNames(page)
	if err != nil {
		beego.Error(err)
	}

	c.Data["pages"] = CalcPaginate(page,100)
	c.Data["awcs"] = awcs

	catagories, err := models.GetCatagories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["catagories"] = catagories

	c.Data["IsHome"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "home.html"
}
