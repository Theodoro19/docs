package example

type Mechanical struct{
	wheels int
	doors int
	speed int
	axles bool
}

func (truck Mechanical) WheelCount() int {
	return truck.wheels
}

func (truck Mechanical) NumberOfDoors() int {
	return truck.doors
}

func (truck Mechanical) MaximumSpeed() int {
	return truck.speed
}

func (truck Mechanical) NumberOfAxles() int {
	return truck.axles
}