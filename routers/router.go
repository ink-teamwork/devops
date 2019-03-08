package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

var FilterUser = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/login") {
		return
	}
	_, ok := ctx.Input.Session("uid").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.MainController{}, "get:Test")
	beego.Router("/login", &controllers.UserController{}, "get:Login")

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

}
