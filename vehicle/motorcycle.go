package vehicle

// Motorcycle embeds LandVehicle, obtaining its methods and fields.
type Motorcycle struct {
	LandVehicle
	HasCargo bool
}

// Motorcycle constructor
func newMotorcycle(licensePlate string, hasCargo bool) Motorcycle {
	return Motorcycle{
		LandVehicle: LandVehicle{
			numberOfWheels: 2,
			licensePlate:   licensePlate,
			category:       A,
		},
		HasCargo: hasCargo,
	}
}

var _ Vehicle = (*Motorcycle)(nil) // Will cause a compile error if Motorcycle lacks any Vehicle methods
