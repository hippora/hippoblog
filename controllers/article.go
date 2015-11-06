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

	c.Data["IsArticle"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplNames = "article_add.html"
}

func (c *ArticleController) Post() {
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	if len(title) !=0 && len(content) !=0 {
		models.AddArticle(title,content)
	}
	c.Redirect("/article",301)
}