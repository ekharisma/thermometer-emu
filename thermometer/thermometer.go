package thermometer

type thermometer struct {
	temperature float32
}

func NewThermometer() thermometer {
	return thermometer{}
}

func (t thermometer) GetTemperature() float32 {
	return 1.0
}
