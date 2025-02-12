package vehicle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const MOCK_LICENSE_PLATE = "XYZ-1234"

type testCarBehavior struct {
	name    string
	car     Car
	wantErr bool
	asserts func()
}

func buildTestCarInitialization(t *testing.T) testCarBehavior {
	name := "Test Car Initialization"

	car := NewCar(MOCK_LICENSE_PLATE, 5, 4)

	return testCarBehavior{
		name: name,
		car:  car,
		asserts: func() {
			assert.Equal(t, MOCK_LICENSE_PLATE, car.licensePlate)
			assert.Equal(t, 5, car.numberOfSeats)
			assert.Equal(t, 4, car.numberOfDoors)
			assert.Equal(t, 4, car.numberOfWheels)
			assert.Equal(t, B, car.category)
		},
		wantErr: false,
	}
}

func buildTestCarAccelerate(t *testing.T) testCarBehavior {
	name := "Test Car Accelerate"

	car := NewCar(MOCK_LICENSE_PLATE, 5, 4)

	return testCarBehavior{
		name: name,
		car:  car,
		asserts: func() {
			car.Accelerate()
		},
		wantErr: false,
	}
}

func buildTestCarSlowDown(t *testing.T) testCarBehavior {
	name := "Test Car SlowDown"

	car := NewCar(MOCK_LICENSE_PLATE, 5, 4)

	return testCarBehavior{
		name: name,
		car:  car,
		asserts: func() {
			car.SlowDown()
		},
		wantErr: false,
	}
}

func buildTestCarGetInfo(t *testing.T) testCarBehavior {
	name := "Test Car GetInfo"

	car := NewCar(MOCK_LICENSE_PLATE, 5, 4)

	return testCarBehavior{
		name: name,
		car:  car,
		asserts: func() {
			info := car.GetInfo()
			assert.Contains(t, info, MOCK_LICENSE_PLATE)
			assert.Contains(t, info, "category B")
		},
		wantErr: false,
	}
}

func TestCarBehavior(t *testing.T) {
	tests := []testCarBehavior{
		buildTestCarInitialization(t),
		buildTestCarAccelerate(t),
		buildTestCarSlowDown(t),
		buildTestCarGetInfo(t),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.asserts()
		})
	}
}
