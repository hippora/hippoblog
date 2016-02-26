package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hippora/hippoblog/models"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	// 是否是删除操作
	switch c.Input().Get("op") {
	case "del":
		if ! IsLogin(c.Ctx) {
			c.Redirect("/login",301)
			return
		}
		models.DelArticle(c.Input().Get("id"))
		break
	}
	//获取显示第几页
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

	c.Data["IsArticle"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplName = "article.html"
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
	c.TplName = "article_add.html"
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
	id := c.Ctx.Input.Params()["0"]
	//获取所有文章
	article, err := models.GetArticle(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["article"] = article

	c.Data["IsArticle"] = true
	c.Data["IsLogin"] = IsLogin(c.Ctx)
	c.TplName = "article_view.html"
}