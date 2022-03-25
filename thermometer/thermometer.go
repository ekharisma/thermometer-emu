package thermometer

import "time"

type Thermometer struct {
	temperature float32
}

func NewThermometer() *Thermometer {
	return &Thermometer{
		temperature: 20.0,
	}
}

func (t *Thermometer) GetTemperature() float32 {
	//t.temperature = 20 + rand.Float32()*(30-20)
	if t.temperature <= 30 {
		t.temperature++
	} else {
		t.temperature = 20.0
	}
	time.Sleep(2 * time.Second)
	return t.temperature
}
