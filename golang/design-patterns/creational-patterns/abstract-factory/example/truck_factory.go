package example

import (
	"errors"
	"fmt"
)

const (
	SIMPLE_MECHANICAL_HORSE = 1
	ROAD_TRAIN = 2
)

type TruckFactory struct{}

func (f *TruckFactory) NewVehicle(truckType int) (Vehicle, error){
	switch truckType {
	case SIMPLE_MECHANICAL_HORSE:
		return new(Mechanical), nil
	case ROAD_TRAIN:
		return new(Road),nil
	default:
		return nil, errors.New(fmt.Sprintf("Truck type not found:%d", truckType))
	}
}
