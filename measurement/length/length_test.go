package length

import (
	"testing"
)

func TestLengthConvertToBaseUnit(t *testing.T) {
	testCases := []struct {
		value    float64
		unit     LengthUnit
		expected float64
	}{
		{1, Kilometer, 1000},
		{1, Meter, 1},
		{1, Centimeter, 0.01},
		{1, Millimeter, 0.001},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			length := NewLengthUnit(tc.unit)
			got := length.convertToBaseUnit(tc.value)
			if got != tc.expected {
				t.Error("ConvertToBaseUnit(%f, %v) = %f, want %f", tc.value, tc.unit, tc.expected)
			}
		})
	}
}
