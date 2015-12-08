package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hippora/hippoblog/models"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	awcs, err := models.GetArticleWithCataNames()
	if err != nil {
		beego.Error(err)
	}

	c.Data["awcs"] = awcs

	c.Data["IsArticle"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "article.html"
}

func (c *ArticleController) Add() {
	if ! IsLogin(c.Ctx) {
		c.Redirect("/login",301)
		return
	}
	catagories,err := models.GetCatagories()
	if err!= nil {
		beego.Error(err)
	}
	c.Data["catagories"] = catagories

	c.Data["IsArticle"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "article_add.html"
}

func (c *ArticleController) Post() {
	if ! IsLogin(c.Ctx) {
		c.Redirect("/login",301)
		return
	}
	title := c.Input().Get("title")
	cid := c.Input().Get("cid")
	content := c.Input().Get("content")
	if len(title) !=0 && len(content) !=0 {
		models.AddArticle(title,cid,content)
	}
	c.Redirect("/article",301)
}

func (c *ArticleController) View() {
	id := c.Ctx.Input.Params["0"]
	article, err := models.GetArticle(id)
	if err != nil {
		beego.Error(err)
	}

	c.Data["article"] = article

	c.Data["IsArticle"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "article_view.html"
}