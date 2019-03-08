package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	username :=c.Input().Get("username")
	password :=c.Input().Get("password")

	println(username)
	println(password)

	c.Ctx.WriteString("123")

}
