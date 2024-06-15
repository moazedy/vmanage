package appservice

import (
	"context"
	"vmanage/pkg/infra/errorx"
	"vmanage/pkg/module/vmanage/application/dto"
)

type Vehicle interface {
	Create(context.Context, dto.CreateVehicleRequest) (dto.CreateVehicleResponse, errorx.ErrorX)
	Update(context.Context, dto.UpdateVehicleRequest) errorx.ErrorX
	Delete(context.Context, dto.DeleteVehicleRequest) errorx.ErrorX
	GetByID(context.Context, dto.GetVehicleByIDRequest) (dto.GetVehicleByIDResponse, errorx.ErrorX)
	GetByTitle(context.Context, dto.GetVehicleByTitleRequest) (dto.GetVehicleByTitleResponse, errorx.ErrorX)
}
