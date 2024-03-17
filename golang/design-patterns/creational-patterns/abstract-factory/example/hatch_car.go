package example

type Hatch struct{
	wheels int
	doors int
	speed int
	isElectric bool
}

func (car Hatch) WheelCount() int {
	return car.wheels
}

func (car Hatch) NumberOfDoors() int {
	return car.doors
}

func (car Hatch) MaximumSpeed() int {
	return car.speed
}

func (car Hatch) HasElectricEngine() bool {
	return car.isElectric
}