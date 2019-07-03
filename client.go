package main

import (
	"fmt"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var SUB_TOPIC = "cmd"

var mqttClient mqtt.Client

// StartClient starts the local mqtt client
func StartClient(port string) {
	time.Sleep(time.Millisecond * 1)

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://0.0.0.0:" + port)
	opts.SetClientID("local-client")
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println(" ===> ready")

	mqttClient.Subscribe(SUB_TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {

		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		topic := msg.Topic()
		if string([]rune(topic)[0]) == "/" {
			topic = topic[1:]
		}

		// msg.Payload()

		// PUB TO MQTT
		// token := client.Publish(topic, 0, false, string(out))
		// token.Wait()
		// if token.Error() != nil {
		// 	fmt.Println(token.Error())
		// }
	})

	//go loopSend()
}

var notes = []int{60, 62, 64, 67, 69}
var idx = 0

func loopSend() {
	if idx == 4 {
		idx = 0
	} else {
		idx = idx + 1
	}
	msg := "1:" + strconv.Itoa(notes[idx]) + ":100"
	token := mqttClient.Publish("asdf", 0, false, msg)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Println(msg)
	time.Sleep(time.Millisecond * 100)
	loopSend()
}
