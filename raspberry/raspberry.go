package raspberry

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/thermometer-emu/thermometer"
)

type MqttService interface {
	// Initialize connection()
	InitBroker() mqtt.Client
	Publish(client mqtt.Client, topic string, payload string)
	Subscribe(client mqtt.Client, topic string)
}

type Raspberry struct {
	thermometer.Thermometer
}

func NewRaspi() Raspberry {
	return Raspberry{}
}

func (raspi Raspberry) GetData(usb chan float32, thermometer thermometer.Thermometer) {
	for {
		usb <- thermometer.GetTemperature()
	}
}

var messagePubHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("Received message: %s from topic %s", m.Topic(), m.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func (raspi Raspberry) InitBroker() mqtt.Client {
	broker := "broker.emqx.io"
	port := 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("mqtt-client")
	opts.OnConnect = connectHandler
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func (raspi Raspberry) Publish(client mqtt.Client, topic string, payload string) {
	token := client.Publish(topic, 2, false, payload)
	token.Wait()
}

func (raspi Raspberry) Subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 2, nil)
	token.Wait()
}
