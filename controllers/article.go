package controllers

import (
	"github.com/astaxie/beego"
	"hippoblog/models"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	articles, err := models.GetArticles()
	if err != nil {
		beego.Error(err)
	}

	c.Data["articles"] = articles

	c.Data["IsArticle"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "article.html"
}

func (c *ArticleController) Add() {
	if ! IsLogin(c.Ctx) {
		c.Redirect("/login",301)
		return
	}
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
	content := c.Input().Get("content")
	if len(title) !=0 && len(content) !=0 {
		models.AddArticle(title,content)
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