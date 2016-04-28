package controllers

import (
	"net/http"
	"net/url"

	"github.com/go-oauth2/serversample/models"

	"github.com/astaxie/beego"
)

// AuthorizeRequest 验证请求
type AuthorizeRequest struct {
	ResponseType string
	ClientID     string
	RedirectURI  string
	Scope        string
	State        string
}

type AuthorizeController struct {
	beego.Controller
}

func (ac *AuthorizeController) Get() {
	redirectURI, err := url.QueryUnescape(ac.GetString("redirect_uri"))
	if err != nil {
		panic(err)
	}
	authorizeReq := AuthorizeRequest{
		ResponseType: ac.GetString("response_type"),
		ClientID:     ac.GetString("client_id"),
		RedirectURI:  redirectURI,
		Scope:        ac.GetString("scope"),
		State:        ac.GetString("state"),
	}
	if authorizeReq.ResponseType == "" ||
		authorizeReq.ClientID == "" ||
		authorizeReq.RedirectURI == "" {
		panic("未知的认证信息")
	}
	ac.SetSession("Authorize", authorizeReq)
	if ac.GetSession("UID") == nil {
		ac.Redirect("/login", http.StatusFound)
		return
	}
	cli := new(models.Client)
	clientInfo, err := cli.GetByID(authorizeReq.ClientID)
	if err != nil {
		panic(err)
	}
	ac.Data["ClientName"] = clientInfo.Name
	ac.TplName = "authorization.tpl"
}

func (ac *AuthorizeController) Post() {
	authorize := ac.GetSession("Authorize")
	if authorize == nil {
		panic("无效的授权")
	}
	uid := ac.GetSession("UID")
	if uid == nil {
		panic("未知的用户")
	}
	userID := uid.(string)
	authorizeReq := authorize.(AuthorizeRequest)
	switch authorizeReq.ResponseType {
	case "code":
		code, err := oAuthManager.GetACManager().GenerateCode(authorizeReq.ClientID, userID, authorizeReq.RedirectURI, authorizeReq.Scope)
		if err != nil {
			panic(err)
		}
		redirectURL, err := url.Parse(authorizeReq.RedirectURI)
		if err != nil {
			panic(err)
		}
		redirectValues := url.Values{}
		redirectValues.Set("code", code)
		redirectValues.Set("state", authorizeReq.State)
		redirectURL.RawQuery = redirectValues.Encode()
		ac.Redirect(redirectURL.String(), http.StatusFound)
	default:
		panic("未识别的授权类型")
	}
}
