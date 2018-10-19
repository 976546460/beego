package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beego/models"
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
	txt := ""
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = "你好世界"

	if _,err := o.Insert(user); err != nil {
		txt = err.Error()

	}else{
		//c.Ctx.WriteString(" + " + strconv.FormatInt(id,10))

		//where := models.User{Name: "你好世界"}
		var users []*models.User
		//err = o.Read(&where,"name"); err != nil
		num , err :=o.QueryTable("user").Filter("name","你好世界").All(&users,"id","name")
		if err != nil {
			if err == orm.ErrNoRows{
				txt = "找不到数据!"
			}else if err == orm.ErrMissPK {
				txt = "找不到主键!"
			}else{
				txt = "未知错误!"
			}
			c.Ctx.WriteString(txt)
		}else{
			c.Data["a"] = users
			c.Data["num"] = num
			/*for _,m := range users {
				q_id := m.Id
				fmt.Println(q_id)
			}*/
			//fmt.Println("id = %d ----   name = %s ",num)
		}
		c.TplName = "index.html"


	}
	c.Ctx.WriteString(txt)
}
