package mqtt

import (
	"fmt"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

type Factory func(config *Config) (*Client, error)

func NewClient(config *Config) (*Client, error) {
	if nil != config.Logger.Debug {
		pahoMqtt.DEBUG = config.Logger.Debug
	}
	if nil != config.Logger.Critical {
		pahoMqtt.CRITICAL = config.Logger.Critical
	}
	if nil != config.Logger.Error {
		pahoMqtt.ERROR = config.Logger.Error
	}

	ServerPort := 1883
	if 0 != config.Port {
		ServerPort = config.Port
	}
	ServerAddress := fmt.Sprintf("tcp://%s:%d", config.Server, ServerPort)

	opts := pahoMqtt.NewClientOptions()
	opts.AddBroker(ServerAddress)
	opts.SetClientID(config.Hostname)
	opts.SetKeepAlive(config.Options.KeepAlive)
	opts.SetPingTimeout(config.Options.PingTimeout)
	opts.SetDefaultPublishHandler(config.Options.Handler)

	return &Client{
		Client: pahoMqtt.NewClient(opts),
	}, nil
}
