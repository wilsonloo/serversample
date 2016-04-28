package controllers

import (
	"net/url"
	"time"

	"github.com/astaxie/beego"
)

// TokenRequest 令牌请求参数
type TokenRequest struct {
	GrantType    string `form:"grant_type"`
	Code         string `form:"code"`
	RedirectURI  string `form:"redirect_uri"`
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
}

// TokenResponse 令牌响应参数
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// TokenController 令牌
type TokenController struct {
	beego.Controller
}

// Post Post
func (tc *TokenController) Post() {
	var tokenReq TokenRequest
	if err := tc.ParseForm(&tokenReq); err != nil {
		panic(err)
	}
	tc.Ctx.Output.Header("Cache-Control", "no-store")
	tc.Ctx.Output.Header("Pragma", "no-cache")
	switch tokenReq.GrantType {
	case "authorization_code": // 授权码模式
		redirectURI, err := url.QueryUnescape(tokenReq.RedirectURI)
		if err != nil {
			panic(err)
		}
		token, err := oAuthManager.GetACManager().
			GenerateToken(tokenReq.Code, redirectURI, tokenReq.ClientID, tokenReq.ClientSecret, true)
		if err != nil {
			panic(err)
		}
		tokenRes := TokenResponse{
			AccessToken:  token.AccessToken,
			TokenType:    "Bearer",
			ExpiresIn:    int64(token.ATExpiresIn / time.Second),
			RefreshToken: token.RefreshToken,
			Scope:        token.Scope,
		}
		tc.Data["json"] = tokenRes
		tc.ServeJSON()
	default:
		panic("未知的授权类型")
	}
}
