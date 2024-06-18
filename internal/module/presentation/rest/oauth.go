package rest

import (
	"net/http"
	"vmanage/pkg/infra/cookie"
	presentation "vmanage/pkg/module/vmanage/persentation"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	loginUrl    = "http://localhost:4853/auth/google/login"
	callbackUrl = "http://localhost:4853/auth/google/callback"
)

var oauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:4853/auth/google/callback",
	ClientID:     "",
	ClientSecret: "",
	// NOTE : just for example on scopes
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
	// NOTE : in a real project, redirect  path will be to a oauth provider like google.
	// and after user permission grant, callback will be called by provider.
	// in this mock example grant is being bypassed.
	/*
		oauthState := generateStateOauthCookie()
		redirectUrl := oauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	*/
	redirectUrl := callbackUrl
	ctx.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}

func (oa oauth) Callback(ctx *gin.Context) {
	// NOTE : in a real project, code will be exchanged with access token and using access token,
	// user data with defined scopes will be fetched from data resource.
	// in this example data will be considered a mock data.
	/*
		data, err := getUserDataFromGoogle(ctx.Request.FormValue("code"))
		if err != nil {
			fmt.Println("error on reading user data from google: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	*/
	data := []byte("user data")

	cookie.SetCookie(ctx, data)
	ctx.JSON(http.StatusOK, string(data))
}

func (oa oauth) CheckLogin(ctx *gin.Context) {
	cookie.LoginCheck(ctx, loginUrl)
	if ctx.IsAborted() {
		return
	}
}
