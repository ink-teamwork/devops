package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Id string
	Username string
	Password string

}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test() {
	o := orm.NewOrm()
	var users []User
	n, err := o.Raw("select * from user").QueryRows(&users)
	if err != nil {
		println("查询失败")
	}
	println(n)
	c.Ctx.Output.JSON(&users, true, false)

}
