package application

import (
	"context"
	"vmanage/pkg/infra/errorx"
	"vmanage/pkg/infra/tx"
	"vmanage/pkg/module/vmanage/application/appservice"
	"vmanage/pkg/module/vmanage/application/dto"
	"vmanage/pkg/module/vmanage/derror"
	"vmanage/pkg/module/vmanage/model/entity"
	"vmanage/pkg/module/vmanage/repository"
)

const (
	fieldNameTitle = "title"
	fieldNameID    = "id"
)

type vehicle struct {
	txFactory          tx.TXFactory
	vehicleRepoFactory repository.GenericRepoFactory[entity.Vehicle]
}

func NewVehicle(txFactory tx.TXFactory, vehicleRepoFactory repository.GenericRepoFactory[entity.Vehicle]) appservice.Vehicle {
	return vehicle{
		txFactory:          txFactory,
		vehicleRepoFactory: vehicleRepoFactory,
	}
}

func (v vehicle) Create(ctx context.Context, req dto.CreateVehicleRequest) (out dto.CreateVehicleResponse, errx errorx.ErrorX) {
	if errx = req.Validate(ctx); !errx.IsNil() {
		return
	}

	tx := v.txFactory.NewTX()
	out, errx = v.create(ctx, tx, req)
	errx = tx.AutoCR(errx)
	return
}

func (v vehicle) create(ctx context.Context, tx tx.TX, req dto.CreateVehicleRequest) (out dto.CreateVehicleResponse, errx errorx.ErrorX) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	vehicleRepo := v.vehicleRepoFactory.NewGenericRepo(tx)
	sameTitledVehicle, err := vehicleRepo.GetByStringField(ctx, fieldNameTitle, req.Title)
	if err != nil {
		errx = errorx.NewInternalErrorX(err)
		return
	}

	if sameTitledVehicle != nil {
		errx = errorx.NewBadRequestErrorX(derror.ErrVehicleWithTitleAlreadyExists)
		return
	}

	vehicleEntity, err := vehicleRepo.Create(ctx, createDtoToVehicleEntity(req))
	if err != nil {
		errx = errorx.NewInternalErrorX(err)
		return
	}

	out.Vehicle = toVehicleDto(*vehicleEntity)
	return
}

func (v vehicle) Update(ctx context.Context, req dto.UpdateVehicleRequest) (errx errorx.ErrorX) {
	if errx = req.Validate(ctx); !errx.IsNil() {
		return
	}

	tx := v.txFactory.NewTX()
	errx = v.update(ctx, tx, req)
	errx = tx.AutoCR(errx)
	return
}

func (v vehicle) update(ctx context.Context, tx tx.TX, req dto.UpdateVehicleRequest) (errx errorx.ErrorX) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	vehicleRepo := v.vehicleRepoFactory.NewGenericRepo(tx)
	sameTitledVehicle, err := vehicleRepo.GetByStringField(ctx, fieldNameTitle, req.Title)
	if err != nil {
		errx = errorx.NewInternalErrorX(err)
		return
	}

	if sameTitledVehicle != nil {
		errx = errorx.NewBadRequestErrorX(derror.ErrVehicleWithTitleAlreadyExists)
		return
	}

	_, err = vehicleRepo.Update(ctx, updateDtoToVehicleEntity(req))
	if err != nil {
		errx = errorx.NewInternalErrorX(err)
		return
	}

	return
}

func (v vehicle) Delete(ctx context.Context, req dto.DeleteVehicleRequest) (errx errorx.ErrorX) {
	if errx = req.Validate(ctx); !errx.IsNil() {
		return
	}

	tx := v.txFactory.NewTX()
	errx = v.delete(ctx, tx, req)
	errx = tx.AutoCR(errx)
	return
}

func (v vehicle) delete(ctx context.Context, tx tx.TX, req dto.DeleteVehicleRequest) (errx errorx.ErrorX) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	vehicleRepo := v.vehicleRepoFactory.NewGenericRepo(tx)
	err := vehicleRepo.Delete(ctx, req.VehicleID)
	if err != nil {
		errx = errorx.NewInternalErrorX(err)
		return
	}
	return
}

func (v vehicle) GetByID(ctx context.Context, req dto.GetVehicleByIDRequest) (out dto.GetVehicleByIDResponse, errx errorx.ErrorX) {
	if errx = req.Validate(ctx); !errx.IsNil() {
		return
	}

	tx := v.txFactory.NewTX()
	out, errx = v.getByID(ctx, tx, req)
	errx = tx.AutoCR(errx)
	return
}

func (v vehicle) getByID(ctx context.Context, tx tx.TX, req dto.GetVehicleByIDRequest) (out dto.GetVehicleByIDResponse, errx errorx.ErrorX) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	vehicleRepo := v.vehicleRepoFactory.NewGenericRepo(tx)
	theVehicle, err := vehicleRepo.GetByStringField(ctx, fieldNameID, req.VehicleID)
	if err != nil {
		errx = errorx.NewInternalErrorX(err)
		return
	}

	if theVehicle == nil {
		errx = errorx.NewBadRequestErrorX(derror.ErrVehicleWithIDDoesNotExist)
		return
	}

	out.Vehicle = toVehicleDto(*theVehicle)
	return
}

func (v vehicle) GetByTitle(ctx context.Context, req dto.GetVehicleByTitleRequest) (out dto.GetVehicleByTitleResponse, errx errorx.ErrorX) {
	if errx = req.Validate(ctx); !errx.IsNil() {
		return
	}

	tx := v.txFactory.NewTX()
	out, errx = v.getByTitle(ctx, tx, req)
	errx = tx.AutoCR(errx)
	return
}

func (v vehicle) getByTitle(ctx context.Context, tx tx.TX, req dto.GetVehicleByTitleRequest) (out dto.GetVehicleByTitleResponse, errx errorx.ErrorX) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	vehicleRepo := v.vehicleRepoFactory.NewGenericRepo(tx)
	theVehicle, err := vehicleRepo.GetByStringField(ctx, fieldNameTitle, req.VehicleTitle)
	if err != nil {
		errx = errorx.NewInternalErrorX(err)
		return
	}

	if theVehicle == nil {
		errx = errorx.NewBadRequestErrorX(derror.ErrVehicleWithTitleDoesNotExist)
		return
	}

	out.Vehicle = toVehicleDto(*theVehicle)
	return
}
