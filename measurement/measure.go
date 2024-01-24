package measurement

type Measure struct {
	Value float64
	Unit  Unit
}

func NewMeasure(value float64, unit Unit) Measure {
	return Measure{Value: value, Unit: unit}
}

func (measure Measure) isCompatibleWith(other Measure) bool {
	return measure.Unit == other.Unit
}

func (measure Measure) convertTo(targetUnit Unit) Measure {
	if measure.Unit != targetUnit {
		panic("Cannot add measures with differnt units")
	}
	baseValue := measure.Unit.ConvertToBaseUnit(measure.value)
	resultValue := targetUnit.convertFromBaseUnit(baseValue)
	return Measure{Value: resultValue, Unit: targetUnit}
}
