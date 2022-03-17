package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ekharisma/thermometer-emu/raspberry"
	"github.com/ekharisma/thermometer-emu/thermometer"
)

func main() {
	// random antara 20 - 30
	// push ke mqtt tiap 2 detik
	// ada 2 thermo -> proram yg berjaan di raspi
	rand.Seed(time.Now().Unix())
	usb1 := make(chan float32)
	usb2 := make(chan float32)
	raspi1 := raspberry.NewRaspi()
	raspi2 := raspberry.NewRaspi()
	thermometer1 := thermometer.NewThermometer()
	thermometer2 := thermometer.NewThermometer()
	go raspi1.GetData(usb1, thermometer1)
	go raspi2.GetData(usb2, thermometer2)
	for {
		fmt.Println("Raspi 1 : ", <-usb1)
		fmt.Println("Raspi 2 : ", <-usb2)
	}
}
