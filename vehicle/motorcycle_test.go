package vehicle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testMotorcycleBehavior struct {
	name    string
	car     Motorcycle
	wantErr bool
	asserts func()
}

func buildTestMotorcycleInitialization(t *testing.T) testMotorcycleBehavior {
	name := "Test Car Initialization"

	moto := NewMotorcycle(MOCK_LICENSE_PLATE, false)

	return testMotorcycleBehavior{
		name: name,
		car:  moto,
		asserts: func() {
			assert.Equal(t, MOCK_LICENSE_PLATE, moto.licensePlate)
			assert.Equal(t, false, moto.HasCargo)
			assert.Equal(t, 2, moto.numberOfWheels)
			assert.Equal(t, A, moto.category)
		},
		wantErr: false,
	}
}

func TestMotorcycleBehavior(t *testing.T) {
	tests := []testMotorcycleBehavior{
		buildTestMotorcycleInitialization(t),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.asserts()
		})
	}
}
