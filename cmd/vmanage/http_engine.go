package main

import "github.com/gin-gonic/gin"

func run(address string, items serverItems) {
	engine := gin.Default()

	// APIs
	engine.POST("/vehicles/vehicle", items.vehiclePresentation.Create)
	engine.PUT("/vehicles/vehicle", items.vehiclePresentation.Update)
	engine.DELETE("/vehicles/vehicle/:id", items.vehiclePresentation.Delete)
	engine.GET("/vehicles/vehicle/:id", items.vehiclePresentation.GetByID)
	engine.GET("/vehicles/vehicle/title/:title", items.vehiclePresentation.GetByTitle)

	// running the http engine
	err := engine.Run(address)
	if err != nil {
		panic(err)
	}
}
