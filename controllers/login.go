package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) Get() {
	lc.TplName = "login.tpl"
	if lc.GetSession("LoginError") != nil {
		lc.Data["ErrorMessage"] = "登录错误"
		lc.DelSession("LoginError")
		return
	}
	if lc.GetSession("UID") != nil {
		lc.Redirect("/", http.StatusOK)
	}
}

func (lc *LoginController) Post() {
	userName := lc.Ctx.Request.FormValue("UserName")
	userPwd := lc.Ctx.Request.FormValue("Password")
	if userName == "admin" && userPwd == "admin" {
		lc.SetSession("UID", "000000")
		if v := lc.GetSession("Authorize"); v != nil {
			authorizeReq := v.(AuthorizeRequest)
			redirectURI := fmt.Sprintf("/authorize?response_type=%s&client_id=%s&redirect_uri=%s&scope=%s&state=%s",
				authorizeReq.ResponseType, authorizeReq.ClientID, url.QueryEscape(authorizeReq.RedirectURI), authorizeReq.Scope, authorizeReq.State)
			lc.DelSession("Authorize")
			lc.Redirect(redirectURI, http.StatusFound)
			return
		}
		lc.Redirect("/", http.StatusFound)
	} else {
		lc.SetSession("LoginError", "1")
		lc.Redirect("/login", http.StatusFound)
	}
}
