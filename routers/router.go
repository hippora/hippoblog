package routers

import (
	"hippoblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/logout", &controllers.LogoutController{})
    beego.Router("/catagory", &controllers.CatagoryController{})
    beego.Router("/article", &controllers.ArticleController{})
    beego.AutoRouter(&controllers.ArticleController{})
    beego.Router("/about", &controllers.AboutController{})
}
