package example

type VehicleFactory interface {
	NewVehicle(int) (Vehicle, error)
}