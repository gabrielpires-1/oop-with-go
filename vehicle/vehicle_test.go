package vehicle

import (
	"testing"

	"github.com/gabrielpires-1/oop-with-go.git/mocks"
	"github.com/go-playground/assert/v2"
)

type testVehicleBehavior struct {
	name    string
	mock    *mocks.Vehicle
	wantErr bool
	asserts func()
}

func buildTestVehicleDrive(t *testing.T) testVehicleBehavior {
	name := "Test Vehicle Drive"

	mockVehicle := new(mocks.Vehicle)
	mockVehicle.On("Accelerate").Return()
	mockVehicle.On("SlowDown").Return()
	mockVehicle.On("GetInfo").Return("Mocked Vehicle Info")

	return testVehicleBehavior{
		name: name,
		mock: mockVehicle,
		asserts: func() {
			Drive(mockVehicle)

			mockVehicle.AssertCalled(t, "Accelerate")
			mockVehicle.AssertCalled(t, "SlowDown")
			mockVehicle.AssertCalled(t, "GetInfo")

			info := mockVehicle.GetInfo()
			assert.Equal(t, "Mocked Vehicle Info", info)
		},
	}
}

func TestVehicleBehavior(t *testing.T) {
	tests := []testVehicleBehavior{
		buildTestVehicleDrive(t),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.asserts()
		})
	}
}
