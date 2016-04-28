package routers

import (
	"github.com/astaxie/beego"
	"github.com/go-oauth2/serversample/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/authorize", &controllers.AuthorizeController{})
	beego.Router("/token", &controllers.TokenController{})
}
