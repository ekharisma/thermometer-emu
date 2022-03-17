package raspberry

import "github.com/ekharisma/thermometer-emu/thermometer"

type RaspberryData struct {
	thermometer.Thermometer
}

func NewRaspi() RaspberryData {
	return RaspberryData{}
}

func (raspi RaspberryData) GetData(usb chan float32, thermometer thermometer.Thermometer) {
	for {
		usb <- thermometer.GetTemperature()
	}
}
