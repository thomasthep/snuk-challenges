package mqtt

import (
	"fmt"

	"github.com/eclipse/paho.mqtt.golang"
)

type MqttClientFactory func(mqttClient *MqttClient) (*MqttClient, error)

func NewMqttClient(mqttClient *MqttClient) (*MqttClient, error) {
	if nil != mqttClient.Logger.Debug {
		mqtt.DEBUG = mqttClient.Logger.Debug
	}
	if nil != mqttClient.Logger.Critical {
		mqtt.CRITICAL = mqttClient.Logger.Critical
	}
	if nil != mqttClient.Logger.Error {
		mqtt.ERROR = mqttClient.Logger.Error
	}

	ServerPort := 1883
	if 0 != mqttClient.Port {
		ServerPort = mqttClient.Port
	}
	ServerAddress := fmt.Sprintf("tcp://%s:%d", mqttClient.Server, ServerPort)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(ServerAddress)
	opts.SetClientID(mqttClient.Hostname)
	opts.SetKeepAlive(mqttClient.Options.KeepAlive)
	opts.SetPingTimeout(mqttClient.Options.PingTimeout)
	opts.SetDefaultPublishHandler(mqttClient.Options.Handler)

	return &MqttClient{
		Client: mqtt.NewClient(opts),
	}, nil
}
