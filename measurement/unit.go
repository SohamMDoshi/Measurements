package measurement

type Unit interface {
	convertToBaseUnit(value float64) float64
	convertFromBaseUnit(baseValue float64) float64
}
