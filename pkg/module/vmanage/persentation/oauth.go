package presentation

import "github.com/gin-gonic/gin"

type OAuth interface {
	Login(*gin.Context)
	Callback(*gin.Context)
	CheckLogin(*gin.Context)
}
