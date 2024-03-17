package example

import "fmt"
import "errors"

const (
	WHITE = 1
	GREEN = 2
	BLACK = 3
)

type ProductInfoGetter interface {
	GetInfo() string
}

type Vehicle struct {
	Model string 
	Color string
}

func (v Vehicle) GetInfo() string {
	return fmt.Sprintf("Vehicle model: '%s', Color: '%s'.", v.Model, t.Color)
}

var whiteVehicle *Vehicle = &Vehicle{ "Ford Ka", "white" }
var greenVehicle *Vehicle = &Vehicle{ "Ford Ecosport", "green" }
var greenVehicle *Vehicle = &Vehicle{ "Ford Escort", "black" }

func GetVehicleClone(t int) (ProductInfoGetter, error) {
	switch v {
		case WHITE:
			return whiteVehicle, nil
		case GREEN:
			return greenVehicle, nil
		case BLACK:
			return greenVehicle, nil
		default:
		return nil, errors.New(fmt.Sprintf("Prototype id %d not recognized.", v))
	}
}