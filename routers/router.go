package routers

import (
	"github.com/astaxie/beego"
	"beego/mybeego/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Index")
}
