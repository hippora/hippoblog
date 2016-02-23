package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type PageItem struct {
	IsCurrent bool
	Current int
}

type Paginate struct {
	HasPrevious,HasNext bool
	Previous,Next int
	Items [9]*PageItem
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

func CalcPaginate(pageNum,maxPages int64) *Paginate {
	paginate := new(Paginate)
	paginate.HasPrevious = pageNum > 1
	paginate.HasNext = pageNum < maxPages
	paginate.Previous = int(pageNum) -1
	paginate.Next = int(pageNum) + 1
	for i,v := range [...]int{-4,-3,-2,-1,0,1,2,3,4} {
		item := new(PageItem)
		item.Current = v + int(pageNum)
		item.IsCurrent = v == 0
		paginate.Items[i] = item
	}
	return paginate
}