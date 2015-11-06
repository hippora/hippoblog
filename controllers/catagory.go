package controllers

import (
	"github.com/astaxie/beego"
	"hippoblog/models"
)

type CatagoryController struct {
	beego.Controller
}

func (c *CatagoryController) Get() {
	switch c.Input().Get("op") {
	case "del":
		models.DelCatagory(c.Input().Get("id"))
		break
	}
	catagories,err := models.GetCatagories()
	if err!= nil {
		beego.Error(err)
	}
	c.Data["catagories"] = catagories

	c.Data["IsCatagory"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "catagory.html"
}

func (c *CatagoryController) Post() {
	catagory := c.Input().Get("catagory")
	if len(catagory) != 0 {
		models.AddCatagory(catagory)
	}
	c.Redirect("/catagory",301)
}
