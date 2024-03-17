package example

type Sedan struct{
	wheels int
	doors int
	speed int
	isElectric bool
}

func (car Sedan) WheelCount() int {
	return car.wheels
}

func (car Sedan) NumberOfDoors() int {
	return car.doors
}

func (car Sedan) MaximumSpeed() int {
	return car.speed
}

func (car Sedan) HasElectricEngine() bool {
	return car.isElectric
}