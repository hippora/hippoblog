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
	awcs, err := models.GetArticleCatagory(page)
	if err != nil {
		beego.Error(err)
	}

	totalNum,err := models.GetArticleCount()
	if err != nil {
		beego.Error(err)
	}
	c.Data["pages"] = CalcPaginate(page,10,totalNum)
	c.Data["awcs"] = awcs

	catagories, err := models.GetCatagories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["catagories"] = catagories

	c.Data["IsHome"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplName = "home.html"
}
