package cookie

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var key = []byte("some secret")

// Claims are information that being stored in jwt
type Claims struct {
	UserData []byte `json:"userData"`
	jwt.StandardClaims
}

func SetCookie(ctx *gin.Context, userData []byte) {
	expTime := time.Now().Add(10 * time.Minute)
	claims := Claims{
		UserData: userData,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetCookie(
		"vmanage",
		tokenString,
		int(expTime.Unix()),
		"/",
		"localhost",
		false,
		true,
	)
}

func LoginCheck(ctx *gin.Context, redirectPath string) (claims Claims) {
	tokenString, err := ctx.Cookie("vmanage")
	if err != nil {
		fmt.Println(err)
		if err == http.ErrNoCookie {
			ctx.Redirect(http.StatusTemporaryRedirect, redirectPath)
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		fmt.Println(err)
		if err == jwt.ErrSignatureInvalid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !token.Valid {
		fmt.Println("token is invalid")
		ctx.Redirect(http.StatusTemporaryRedirect, redirectPath)
		return
	}

	return
}
