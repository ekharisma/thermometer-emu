package thermometer

import (
	"math/rand"
	"time"
)

type Thermometer struct {
	temperature float32
}

func NewThermometer() Thermometer {
	return Thermometer{}
}

func (t Thermometer) GetTemperature() float32 {
	t.temperature = 20 + rand.Float32()*(30-20)
	time.Sleep(2 * time.Second)
	return t.temperature
}
