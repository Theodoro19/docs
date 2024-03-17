package example

import (
	"errors"
	"fmt"
)

const (
	HATCH_CAR = 1
	SEDAN_CAR = 2
)

type CarFactory struct{}

func (f *CarFactory) NewVehicle(carType int) (Vehicle, error){
	switch carType {
	case HATCH_CAR:
		return new(Hatch), nil
	case SEDAN_CAR:
		return new(Sedan), nil
	default:
		return nil, errors.New(fmt.Sprintf("Car type not found:%d", carType))
	}
}