package vehicle

// Vehicle interface
type Vehicle interface {
	Accelerate()
	SlowDown()
	GetInfo() string
}

func Drive(v Vehicle) string {
	v.Accelerate()
	v.SlowDown()
	return v.GetInfo()
}
