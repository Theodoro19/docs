package example

type Road struct{
	wheels int
	doors int
	speed int
	axles bool
}

func (truck Road) WheelCount() int {
	return truck.wheels
}

func (truck Road) NumberOfDoors() int {
	return truck.doors
}

func (truck Road) MaximumSpeed() int {
	return truck.speed
}

func (truck Road) NumberOfAxles() int {
	return truck.axles
}