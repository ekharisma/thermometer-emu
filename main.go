package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ekharisma/thermometer-emu/raspberry"
	"github.com/ekharisma/thermometer-emu/thermometer"
)

type Payload struct {
	Timestamp   time.Time  `json:"timestamp"`
	Temperature [2]float32 `json:"temperature"`
}

// backend terima json
// subscribe topic yg sama
// connect ke db, ditaruh ke global variable dibackend (array of temperature)

func main() {

	// random antara 20 - 30
	// push ke mqtt tiap 2 detik
	// ada 2 thermo -> proram yg berjaan di raspi
	rand.Seed(time.Now().Unix())
	var temperature1 float32 = -1.0
	var temperature2 float32 = -1.0
	usb1 := make(chan float32)
	usb2 := make(chan float32)
	raspi := raspberry.NewRaspi()
	raspiClient := raspi.InitBroker()
	thermometer1 := thermometer.NewThermometer()
	thermometer2 := thermometer.NewThermometer()
	go raspi.GetData(usb1, *thermometer1)
	go raspi.GetData(usb2, *thermometer2)
	for {
		select {
		case temperature1 = <-usb1:
			fmt.Println("1 : ", temperature1)
		case temperature2 = <-usb2:
			fmt.Println("2 : ", temperature2)
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("Heartbeat")
		}
		payload := Payload{
			Timestamp:   time.Now(),
			Temperature: [2]float32{temperature1, temperature2},
		}
		message, err := json.Marshal(payload)
		if err != nil {
			log.Panic(err)
		}
		// temperature1 := <-usb1
		// temperature2 := <-usb2
		// fmt.Println("Thermometer 1 : ", temperature1)
		// fmt.Println("Thermometer 2 : ", temperature2)
		fmt.Println()
		fmt.Println(string(message))
		raspi.Publish(raspiClient, "/demo/pp/3", string(message))
	}
}
