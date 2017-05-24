package mqtt

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

func TestNewMqttClient(t *testing.T) {
	handler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	}

	logger := Logger{
		Debug: log.New(os.Stdout, "DEBUG: ", 0),
		Critical: log.New(os.Stdout, "CRITICAL: ", 0),
		Error: log.New(os.Stdout, "ERROR: ", 0),
	}

	options := Options{
		KeepAlive: 2 * time.Second,
		PingTimeout: 1 * time.Second,
		Handler: handler,
	}

	client, err := NewMqttClient(&Config{
		Hostname: "myLocal",
		Server: "localhost",
		Port: 1883,
		Options: options,
		Logger: logger,
	})

	if nil != err {
		t.Error(err.Error())
	}

	if nil == client {
		t.Errorf("Expected MQTTClient, got: nil")
	}
}
