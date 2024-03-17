package example

type SedanCar struct {
	wheels int
	doors int
	speed int
	isElectric bool
}

type CarBuilder interface {
	SetWheelCount() CarBuilder
	SetNumberOfDoors() CarBuilder
	SetMaximumSpeed() CarBuilder
	SetHasElectricEngine() CarBuilder
	GetCarType() SedanCar
}

type ElectricCarBuilder struct {
	sedan SedanCar
}

func (electric *ElectricCarBuilder) SetWheelCount() CarBuilder {
	electric.sedan.wheels = 4
	return electric
}

func (electric *ElectricCarBuilder) SetNumberOfDoors() CarBuilder {
	electric.sedan.doors = 4
	return electric
}

func (electric *ElectricCarBuilder) SetMaximumSpeed() CarBuilder {
	electric.sedan.speed = 300
	return electric
}

func (electric *ElectricCarBuilder) SetHasElectricEngine() CarBuilder {
	electric.sedan.isElectric = true
}

func (electric *ElectricCarBuilder) GetCarType() SedanCar {
	return electric.sedan
}

type MechanicalCarBuilder struct {
	sedan SedanCar
}

func (mechanical *MechanicalCarBuilder) SetWheelCount() CarBuilder {
	mechanical.sedan.wheels = 4
	return mechanical
}

func (mechanical *MechanicalCarBuilder) SetNumberOfDoors() CarBuilder {
	mechanical.sedan.doors = 4
	return mechanical
}

func (mechanical *MechanicalCarBuilder) SetMaximumSpeed() CarBuilder {
	mechanical.sedan.speed = 320
	return mechanical
}

func (mechanical *MechanicalCarBuilder) SetHasElectricEngine() CarBuilder {
	mechanical.sedan.isElectric = false
}

func (mechanical *MechanicalCarBuilder) GetCarType() SedanCar {
	return mechanical.sedan
}

type Manufacturer struct {
	car CarBuilder
}

func (m *Manufacturer) SetBuilder(builder CarBuilder) {
	m.car = builder
}

func (m *Manufacturer) CreateVehicle() {
	m.car.SetWheelCount().SetNumberOfDoors().SetMaximumSpeed().SetHasElectricEngine()
}