package main

import "github.com/gin-gonic/gin"

func run(address string, items serverItems) {
	engine := gin.Default()
	engine.GET("/auth/google/login", items.oauthPresentation.Login)
	engine.GET("/auth/google/callback", items.oauthPresentation.Callback)

	privateApisGroup := engine.Group("private")
	privateApisGroup.Use(items.oauthPresentation.CheckLogin)
	privateApisGroup.POST("/vehicles/vehicle", items.vehiclePresentation.Create)
	privateApisGroup.PUT("/vehicles/vehicle", items.vehiclePresentation.Update)
	privateApisGroup.DELETE("/vehicles/vehicle/:id", items.vehiclePresentation.Delete)
	privateApisGroup.GET("/vehicles/vehicle/:id", items.vehiclePresentation.GetByID)
	privateApisGroup.GET("/vehicles/vehicle/title/:title", items.vehiclePresentation.GetByTitle)

	// running the http engine
	err := engine.Run(address)
	if err != nil {
		panic(err)
	}
}
