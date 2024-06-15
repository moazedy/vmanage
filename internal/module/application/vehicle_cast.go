package application

import (
	"vmanage/pkg/module/vmanage/application/dto"
	"vmanage/pkg/module/vmanage/model/entity"
)

func createDtoToVehicleEntity(in dto.CreateVehicleRequest) (out entity.Vehicle) {
	out.Title = in.Title
	return
}

func toVehicleDto(in entity.Vehicle) (out dto.Vehicle) {
	out.ID = in.ID
	out.Title = in.Title
	return
}

func updateDtoToVehicleEntity(in dto.UpdateVehicleRequest) (out entity.Vehicle) {
	out.ID = in.VehicleID
	out.Title = in.Title
	return
}
