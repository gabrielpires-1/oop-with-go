package main

import (
	"fmt"

	v "github.com/gabrielpires-1/oop-with-go.git/vehicle"
)

func main() {
	// Create a Car and Motorcycle using the constructor functions
	car := v.NewCar("XYZ-1234", 5, 4)
	motorcycle := v.NewMotorcycle("ABC-5678", true)

	car.Accelerate()           // Calls overridden method
	fmt.Println(car.GetInfo()) // Uses inherited method from LandVehicle
	car.SlowDown()             // Calls overridden method

	fmt.Println()

	motorcycle.Accelerate()           // Uses inherited method from LandVehicle
	fmt.Println(motorcycle.GetInfo()) // Uses inherited method from LandVehicle
	motorcycle.SlowDown()             // Uses inherited method from LandVehicle

	fmt.Println()

	fmt.Println(v.Drive(car))
}
