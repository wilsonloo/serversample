package routers

import (
	"github.com/astaxie/beego"
	"github.com/go-oauth2/serversample/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/authorize?:response_type&:client_id&:redirect_uri", &controllers.AuthorizeController{})
	beego.Router("/token?:grant_type", &controllers.TokenController{})
}
