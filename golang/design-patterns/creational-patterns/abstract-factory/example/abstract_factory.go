package example

import (
	"errors"
	"fmt"
)

const (
	TRUCK = 1
	CAR = 2
)

func CreateVehicleFactory(vehicleType int) (VehicleFactory, error) {
	switch vehicleType {
	case TRUCK:
		return new(TruckFactory), nil
	case CAR:
		return new(CarFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle type not found:%d", truckType))
	}
}