package rest

import (
	"net/http"
	"vmanage/pkg/module/vmanage/application/appservice"
	"vmanage/pkg/module/vmanage/application/dto"
	presentation "vmanage/pkg/module/vmanage/persentation"

	"github.com/gin-gonic/gin"
)

type vehicle struct {
	vehicleAppService appservice.Vehicle
}

func NewVehicle(vehicleAppService appservice.Vehicle) presentation.Vehicle {
	return vehicle{
		vehicleAppService: vehicleAppService,
	}
}

func (v vehicle) Create(ctx *gin.Context) {
	req := readRequest[dto.CreateVehicleRequest](ctx)
	response, errx := v.vehicleAppService.Create(ctx, req)
	handleErrorx(ctx, errx)

	ctx.JSON(http.StatusCreated, response)
}

func (v vehicle) Update(ctx *gin.Context) {
	req := readRequest[dto.UpdateVehicleRequest](ctx)
	errx := v.vehicleAppService.Update(ctx, req)
	handleErrorx(ctx, errx)

	ctx.JSON(http.StatusOK, nil)
}

func (v vehicle) Delete(ctx *gin.Context) {
	vehicleID := ctx.Param("id")
	var req dto.DeleteVehicleRequest
	req.VehicleID = vehicleID
	errx := v.vehicleAppService.Delete(ctx, req)
	handleErrorx(ctx, errx)

	ctx.JSON(http.StatusOK, nil)

}

func (v vehicle) GetByID(ctx *gin.Context) {
	vehicleID := ctx.Param("id")
	var req dto.GetVehicleByIDRequest
	req.VehicleID = vehicleID
	response, errx := v.vehicleAppService.GetByID(ctx, req)
	handleErrorx(ctx, errx)

	ctx.JSON(http.StatusOK, response)
}

func (v vehicle) GetByTitle(ctx *gin.Context) {
	vehicleTitle := ctx.Param("title")
	var req dto.GetVehicleByTitleRequest
	req.VehicleTitle = vehicleTitle
	response, errx := v.vehicleAppService.GetByTitle(ctx, req)
	handleErrorx(ctx, errx)

	ctx.JSON(http.StatusOK, response)
}
