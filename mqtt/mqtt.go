package mqtt

import (
	"log"
	"time"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

type (
	ClientInterface interface {
		Connect()
		Disconnect()
		Subscribe(string)
		SubscribeHandler(string, pahoMqtt.MessageHandler)
		Unsubscribe(string)
		Publish(string, interface{})
	}
	Client struct {
		Client   pahoMqtt.Client
	}
	Config struct {
		Server   string
		Port     int
		Hostname string
		Options  Options
		Logger   Logger
	}
	Options struct {
		KeepAlive   time.Duration
		PingTimeout time.Duration
		Handler     pahoMqtt.MessageHandler
	}
	Logger struct {
		Debug    *log.Logger
		Critical *log.Logger
		Error    *log.Logger
	}
)

func (m *Client) Connect() {
	if token := m.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *Client) Disconnect() {
	m.Client.Disconnect(250)
}

func (m *Client) Subscribe(topic string) {
	if token := m.Client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *Client) SubscribeHandler(topic string, handler pahoMqtt.MessageHandler) {
	if token := m.Client.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *Client) Unsubscribe(topic string) {
	if token := m.Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *Client) Publish(topic string, payload interface{}) {
	if token := m.Client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
