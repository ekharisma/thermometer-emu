package thermometer_test

import (
	"testing"

	"github.com/ekharisma/thermometer-emu/thermometer"
	"github.com/stretchr/testify/assert"
)

func TestGetTemperature(t *testing.T) {
	tests := []struct {
		name         string
		expected_min float32
		expected_max float32
	}{
		{
			name:         "Test get temperature normal",
			expected_min: 20.0,
			expected_max: 30.0,
		},
	}

	// thinking out-loud
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialization
			th := thermometer.NewThermometer()
			temp := th.GetTemperature()
			assert.True(t, temp > tt.expected_min && temp < tt.expected_max)
			// finalization
		})
	}
}
