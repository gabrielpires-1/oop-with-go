package vehicle

// enum for vehicle category
type Category uint8

const (
	A Category = iota // Motorcycles, mopeds, scooters, and tricycles
	B                 // Cars, utility vehicles, SUVs, minivans, and pickup trucks
	C                 // Small trucks, pickup trucks, and cargo vans
	D                 // Buses and minibuses with more than 8 passenger seats
)

func (c Category) String() string {
	switch c {
	case A:
		return "category A"
	case B:
		return "category B"
	case C:
		return "category C"
	case D:
		return "category D"
	}
	return "unkown"
}
