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
	raspi := raspberry.NewRaspi()
	raspiClient := raspi.InitBroker()
	thermometer1 := thermometer.NewThermometer()
	thermometer2 := thermometer.NewThermometer()
	go raspi.GetData(usb1, thermometer1)
	go raspi.GetData(usb2, thermometer2)
	for {
		fmt.Println("Thermometer 1 : ", <-usb1)
		fmt.Println("Thermometer 2 : ", <-usb2)
		payload := fmt.Sprintf("Thermometer 1 : %v, Thermometer 2 : %v", <-usb1, <-usb2)
		fmt.Println(payload)
		raspi.Publish(raspiClient, "/demo/pp/2", payload)
	}
}
