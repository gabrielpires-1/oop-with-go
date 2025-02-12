package vehicle

import "fmt"

// Car embeds LandVehicle, obtaining its methods and fields.
type Car struct {
	LandVehicle
	numberOfSeats int
	numberOfDoors int
}

// Car constructor
func NewCar(licensePlate string, numberOfSeats, numberOfDoors int) Car {
	return Car{
		LandVehicle: LandVehicle{
			numberOfWheels: 4,
			licensePlate:   licensePlate,
			category:       B,
		},
		numberOfSeats: numberOfSeats,
		numberOfDoors: numberOfDoors,
	}
}

// Overrides inherited methods from LandVehicle
func (car Car) Accelerate() {
	fmt.Println("Car is accelerating...")
}

func (car Car) SlowDown() {
	fmt.Println("Car is slowing down...")
}

var _ Vehicle = (*Car)(nil) // Ensures Car implements Vehicle
