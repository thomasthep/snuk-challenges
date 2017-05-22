package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/robfig/cron"
	"github.com/thomasthep/snuk-challenges/mqtt"
)

func main() {
	fmt.Println("Random Generator")

	r1 := rand.New(rand.NewSource(time.Now().Unix()))

	sigs := make(chan os.Signal, 1)

	mqtt.Connect()
	defer mqtt.Disconnect()

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		randomValue := r1.Intn(10000-1) + 1
		fmt.Printf("-> %d\n", randomValue)
		mqtt.Publish("random", randomValue)
	})
	c.Start()
	defer c.Stop()

	fmt.Println("started")

	sig := <-sigs
	fmt.Printf("event: %s\n", sig)
	fmt.Println("exiting")
}
