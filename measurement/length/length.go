package length

import "measurements/measurement"

type LengthUnit int

const (
	Kilometer LengthUnit = iota
	Meter
	Centimeter
	Millimeter
)

type Length struct {
	scale float64
}

func NewLengthUnit(scale float64) measurement.Unit {
	return &Length{scale: scale}
}

func (u *Length) ConvertToBaseUnit(value float64) float64 {
	return value * u.scale
}

func (u *Length) ConvertFromBaseUnit(baseValue float64) float64 {
	return baseValue / u.scale
}
