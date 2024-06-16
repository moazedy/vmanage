package rest

import (
	"errors"
	"net/http"
	"vmanage/pkg/infra/errorx"
	"vmanage/pkg/module/vmanage/application/dto"

	"github.com/gin-gonic/gin"
)

func readRequest[Dto dto.Dto](ctx *gin.Context) (dto Dto) {
	dtoPtr := new(Dto)
	err := ctx.BindJSON(dto)
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
		ctx.AbortWithError(errx.HttpStatusCode, errx.EmbedError)
	}
}
