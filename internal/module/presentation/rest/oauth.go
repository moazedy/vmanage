package rest

import (
	"fmt"
	"net/http"
	"vmanage/pkg/infra/cookie"
	presentation "vmanage/pkg/module/vmanage/persentation"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var loginUrl = "http://localhost:4853/auth/google/login"

var oauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:4853/auth/google/callback",
	ClientID:     "",
	ClientSecret: "",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/photoslibrary.readonly",
	},
	Endpoint: google.Endpoint,
}

type oauth struct{}

func NewOAuth() presentation.OAuth {
	return oauth{}
}

func (oa oauth) Login(ctx *gin.Context) {
	oauthState := generateStateOauthCookie()
	redirectUrl := oauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	ctx.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}

func (oa oauth) Callback(ctx *gin.Context) {
	data, err := getUserDataFromGoogle(ctx.Request.FormValue("code"))
	if err != nil {
		fmt.Println("error on reading user data from google: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	cookie.SetCookie(ctx, data)
	ctx.JSON(http.StatusOK, string(data))
}

func (oa oauth) CheckLogin(ctx *gin.Context) {
	cookie.LoginCheck(ctx, loginUrl)
}
