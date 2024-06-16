package presentation

import "github.com/gin-gonic/gin"

type Vehicle interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
	GetByTitle(*gin.Context)
}
