package example

import "testing"

func TestManufacturer(t *testing.T) {
	manufacturer := Manufacturer{}

	electricCarBuilder := ElectricCarBuilder{}
	manufacturer.SetBuilder(electricCarBuilder)
	manufacturer.CreateVehicle()
	electric := electricCarBuilder.GetCarType()

	if electric.isElectric != true {
		t.Errorf("Electric Car should be electric but found %s", electric.isElectric)
	}

	mechanicalCarBuilder := MechanicalCarBuilder{}
	manufacturer.SetBuilder(mechanicalCarBuilder)
	manufacturer.CreateVehicle()
	mechanical := mechanicalCarBuilder.GetCarType()

	if mechanical.isElectric != true {
		t.Errorf("Mechanical Car should be mechanic but found %s", mechanical.isElectric)
	}
}