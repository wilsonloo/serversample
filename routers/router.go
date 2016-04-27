package routers

import (
	"github.com/go-oauth2/websample/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
