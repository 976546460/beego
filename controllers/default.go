package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beego/mybeego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
func (c *MainController) Index(){

	o := orm.NewOrm()
	user := new(models.User)
	user.Name = "你好世界"

	if id,err := o.Insert(user); err != nil {
		c.Ctx.WriteString(err.Error())
	}else{
		c.Ctx.WriteString(" + " +string(id))
	}
}
