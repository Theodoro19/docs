package example

import (
	"testing"
)

func TestCarFactory(test *testing.T) {
	carFactory, err := CreateVehicleFactory(CAR)
	if err != nil {
		test.Fatal(err)
	}

	hatchCarVehicle, err := carFactory.NewVehicle(HATCH_CAR)
	if err != nil {
		test.Fatal(err)
	}

	sedanCarVehicle, err := carFactory.NewVehicle(SEDAN_CAR)
	if err != nil {
		test.Fatal(err)
	}

	hatchCar, ok := hatchCarVehicle.(Car)
	if !ok {
		test.Fatal("Assertion failed")
	}

	sedanCar, ok := sedanCarVehicle.(Car)
	if !ok {
		test.Fatal("Assertion failed")
	}
}

func TestTruckFactory(test *testing.T) {
	truckFactory, err := CreateVehicleFactory(TRUCK)
	if err != nil {
		test.Fatal(err)
	}

	mechanicalTruckVehicle, err := truckFactory.NewVehicle(SIMPLE_MECHANICAL_HORSE)
	if err != nil {
		test.Fatal(err)
	}

	roadTruckVehicle, err := truckFactory.NewVehicle(ROAD_TRAIN)
	if err != nil {
		test.Fatal(err)
	}

	mechanicalTruck, ok := mechanicalTruckVehicle.(Truck)
	if !ok {
		test.Fatal("Assertion failed")
	}

	roadTruck, ok := roadTruckVehicle.(Truck)
	if !ok {
		test.Fatal("Assertion failed")
	}
}