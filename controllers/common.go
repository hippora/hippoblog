package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"math"
)

type PageItem struct {
	IsCurrent bool
	Current int
}

type Paginate struct {
	HasPrevious,HasNext bool
	Previous,Next,TotalPages int
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

func CalcPaginate(pageNum,pageCount,maxPages int64) *Paginate {
	p := int(pageNum)
	x := int(math.Ceil(float64(maxPages)/float64(pageCount)))
	if p < 1 {
		p = 1
	}
	if p > x {
		p = x
	}
	paginate := new(Paginate)
	paginate.HasPrevious = p > 1
	paginate.HasNext = p < x
	paginate.Previous = p -1
	paginate.Next = p + 1
	paginate.TotalPages = x
	for i,v := range [...]int{-4,-3,-2,-1,0,1,2,3,4} {
		item := new(PageItem)
		item.Current = v + p
		if item.Current > x {
			item.Current = - 1 //hidden page
		}
		item.IsCurrent = v == 0
		paginate.Items[i] = item
	}
	return paginate
}