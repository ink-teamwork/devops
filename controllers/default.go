package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Name string
	Age int
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test() {
	//user := &User{"zhaokai", 31}
	c.Data["json"] = &User{"zhaokai", 31}
	c.ServeJSON()
}
