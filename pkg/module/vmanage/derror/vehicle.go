package derror

import "errors"

var (
	ErrVehicleWithTitleAlreadyExists = errors.New("vehicle with requested title already exists")
	ErrVehicleWithIDDoesNotExist     = errors.New("vehicle with requested id does not exist")
	ErrVehicleWithTitleDoesNotExist  = errors.New("vehicle with requested title does not exist")
)
