package routers

import (
	"github.com/astaxie/beego"
	"beego/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Index")
}
