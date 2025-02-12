package vehicle

import "fmt"

// LandVehicle implements Vehicle
type LandVehicle struct {
	numberOfWheels int
	licensePlate   string
	category       Category
}

// implements methods defined in interface
func (v LandVehicle) Accelerate() {
	fmt.Println("Land vehicle is accelerating...")
}

func (v LandVehicle) SlowDown() {
	fmt.Println("Land vehicle is slowing down...")
}

func (v LandVehicle) GetInfo() string {
	info := fmt.Sprintf("LAND VEHICLE INFO\nLicense Plate: %v\nCategory: %v", v.licensePlate, v.category.String())
	return info
}
