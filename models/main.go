package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql",
		 //beego.AppConfig.String("mysqluser")+":"+
			//		beego.AppConfig.String("mysqlpass")+"@/"+
			//		beego.AppConfig.String("mysqldb")+"?charset=utf8")

		"root:root@/test?charset=utf8")
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile 	:= new(Profile)
	profile.Age  = 30

	user := new(User)

	user.Name 	 = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}