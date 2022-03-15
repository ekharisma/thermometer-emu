package main

import (
	"fmt"

	"github.com/ekharisma/thermometer-emu/thermometer"
)

func main() {
	// random antara 20 - 30
	// push ke mqtt tiap 2 detik
	// ada 2 thermo -> proram yg berjaan di raspi
	thermo := thermometer.NewThermometer()
	fmt.Println(thermo.GetTemperature())
}
