package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thomasthep/snuk-challenges/collector/aggregator"
	"github.com/thomasthep/snuk-challenges/collector/data"
	"github.com/thomasthep/snuk-challenges/collector/http"
	"github.com/thomasthep/snuk-challenges/mqtt"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("Collector")

	storage := &data.Storage{
		Value: 0,
	}

	bucket := &data.Bucket{
		Name:   "default",
		Values: &[]float64{},
	}

	aggregator := aggregator.Aggregator{
		Storage: storage,
		Bucket:  bucket,
	}

	aggregator.Start(5 * time.Second)
	defer aggregator.Stop()

	sigs := make(chan os.Signal, 1)

	mqtt.Connect()
	defer mqtt.Disconnect()

	var onMessage pahoMqtt.MessageHandler = func(client pahoMqtt.Client, msg pahoMqtt.Message) {
		// fmt.Printf("Topic: %s -> %s\n", msg.Topic(), msg.Payload())
		value, _ := strconv.ParseFloat(string(msg.Payload()), 64)
		bucket.Append(value)
	}
	mqtt.SubscribeHandler("random", onMessage)

	var onRequest = func(ctx *fasthttp.RequestCtx) {
		avg := storage.Get()
		ctx.WriteString(strconv.FormatFloat(avg, 'f', 12, 64))
	}
	http.AddRoute("GET", "/average", onRequest)
	http.Start()

	fmt.Println("started")

	sig := <-sigs
	fmt.Printf("event: %s\n", sig)
	fmt.Println("exiting")
}
