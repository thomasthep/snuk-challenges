package mqtt

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type (
	Mqtt interface {
		Connect()
		Disconnect()
		Subscribe(string)
		SubscribeHandler(string, mqtt.MessageHandler)
		Unsubscribe(string)
		Publish(string, interface{})
	}
	MqttClient struct {
		Client   mqtt.Client
		Server   string
		Port     int
		Hostname string
		Options  MqttOptions
		Logger   MqttLogger
	}
	MqttOptions struct {
		KeepAlive   time.Duration
		PingTimeout time.Duration
		Handler     mqtt.MessageHandler
	}
	MqttLogger struct {
		Debug    *log.Logger
		Critical *log.Logger
		Error    *log.Logger
	}
)

func (m *MqttClient) Connect() {
	if token := m.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *MqttClient) Disconnect() {
	m.Client.Disconnect(250)
}

func (m *MqttClient) Subscribe(topic string) {
	if token := m.Client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *MqttClient) SubscribeHandler(topic string, handler mqtt.MessageHandler) {
	if token := m.Client.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *MqttClient) Unsubscribe(topic string) {
	if token := m.Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *MqttClient) Publish(topic string, payload interface{}) {
	if token := m.Client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
