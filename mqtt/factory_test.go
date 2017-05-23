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

	mqttLogger := MqttLogger{
		Debug: log.New(os.Stdout, "DEBUG: ", 0),
		Critical: log.New(os.Stdout, "CRITICAL: ", 0),
		Error: log.New(os.Stdout, "ERROR: ", 0),
	}

	mqttOptions := MqttOptions{
		KeepAlive: 2 * time.Second,
		PingTimeout: 1 * time.Second,
		Handler: handler,
	}

	mqttClient, err := NewMqttClient(&MqttClient{
		Hostname: "myLocal",
		Server: "localhost",
		Port: 1883,
		Options: mqttOptions,
		Logger: mqttLogger,
	})

	if nil != err {
		t.Error(err.Error())
	}

	if nil == mqttClient {
		t.Errorf("Expected MQTTClient, got: nil")
	}
}
