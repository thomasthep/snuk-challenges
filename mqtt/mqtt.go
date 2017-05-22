package mqtt

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	client mqtt.Client

	handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	}
)

func init() {
	if false {
		mqtt.DEBUG = log.New(os.Stdout, "", 0)
	}
	mqtt.CRITICAL = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	id, _ := os.Hostname()

	opts := mqtt.NewClientOptions().AddBroker("tcp://mosquitto:1883").SetClientID(id)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetDefaultPublishHandler(handler)

	client = mqtt.NewClient(opts)
}

func Connect() {
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Disconnect() {
	client.Disconnect(250)
}

func Subscribe(topic string) {
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func SubscribeHandler(topic string, handler mqtt.MessageHandler) {
	if token := client.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Unsubscribe(topic string) {
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Publish(topic string, payload interface{}) {
	if token := client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
