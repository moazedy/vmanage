package rest

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"vmanage/pkg/infra/errorx"
	"vmanage/pkg/module/vmanage/application/dto"

	"github.com/gin-gonic/gin"
)

func readRequest[Dto dto.Dto](ctx *gin.Context) (dto Dto) {
	dtoPtr := new(Dto)
	err := ctx.BindJSON(dtoPtr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("missmatched type"))
		return
	}

	return *dtoPtr
}

func handleErrorx(ctx *gin.Context, errx errorx.ErrorX) {
	if errx.IsNil() {
		return
	} else {
		ctx.AbortWithStatusJSON(errx.HttpStatusCode, gin.H{"error": errx.EmbedError.Error()})
		return
	}
}

func generateStateOauthCookie() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get("google.some-repo" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	return contents, nil
}
