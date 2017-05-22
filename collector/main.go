package main

import (
	"fmt"
	"os"

	"github.com/thomasthep/snuk-challenges/mqtt"
)

func main() {
	fmt.Println("Collector")

	sigs := make(chan os.Signal, 1)

	mqtt.Connect()
	defer mqtt.Disconnect()

	mqtt.Subscribe("random")

	fmt.Println("started")

	sig := <-sigs
	fmt.Printf("event: %s\n", sig)
	fmt.Println("exiting")
}
