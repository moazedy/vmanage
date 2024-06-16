package dto

import (
	"context"
	"vmanage/pkg/infra/errorx"
	"vmanage/pkg/infra/validator"
)

type Vehicle struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (v Vehicle) IsDto() {}

// ----------------------------------------------------------
type CreateVehicleRequest struct {
	Title string `json:"title" validate:"required,lte=100,gte=3"`
}

func (cvr CreateVehicleRequest) Validate(ctx context.Context) errorx.ErrorX {
	return errorx.NewBadRequestErrorX(validator.Validate(ctx, cvr))
}

func (cvr CreateVehicleRequest) IsDto() {}

type CreateVehicleResponse struct {
	Vehicle Vehicle `json:"vehicle"`
}

// ----------------------------------------------------------

type UpdateVehicleRequest struct {
	VehicleID string `json:"vehicleId" validate:"required,uuid"`
	Title     string `json:"title" validate:"required,lte=100,gte=3"`
}

func (uvr UpdateVehicleRequest) Validate(ctx context.Context) errorx.ErrorX {
	return errorx.NewBadRequestErrorX(validator.Validate(ctx, uvr))
}

func (uvr UpdateVehicleRequest) IsDto() {}

// ----------------------------------------------------------

type DeleteVehicleRequest struct {
	VehicleID string `json:"vehicleId" validate:"required,uuid"`
}

func (dvr DeleteVehicleRequest) Validate(ctx context.Context) errorx.ErrorX {
	return errorx.NewBadRequestErrorX(validator.Validate(ctx, dvr))
}

// ----------------------------------------------------------

type GetVehicleByIDRequest struct {
	VehicleID string `json:"vehicleId" validate:"required,uuid"`
}

func (gvr GetVehicleByIDRequest) Validate(ctx context.Context) errorx.ErrorX {
	return errorx.NewBadRequestErrorX(validator.Validate(ctx, gvr))
}

type GetVehicleByIDResponse struct {
	Vehicle Vehicle `json:"vehicle"`
}

// ----------------------------------------------------------

type GetVehicleByTitleRequest struct {
	VehicleTitle string `json:"vehicleTitle" validate:"required"`
}

func (gvr GetVehicleByTitleRequest) Validate(ctx context.Context) errorx.ErrorX {
	return errorx.NewBadRequestErrorX(validator.Validate(ctx, gvr))
}

type GetVehicleByTitleResponse struct {
	Vehicle Vehicle `json:"vehicle"`
}
