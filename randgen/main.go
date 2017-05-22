package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/thomasthep/snuk-challenges/mqtt"
)

func main() {
	fmt.Println("Random Generator")

	r1 := rand.New(rand.NewSource(time.Now().Unix()))

	sigs := make(chan os.Signal, 1)

	mqtt.Connect()
	defer mqtt.Disconnect()

	ticker := time.NewTicker(100 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				randomValue := r1.Intn(10000-1) + 1
				// fmt.Printf("-> %d\n", randomValue)
				mqtt.Publish("random", strconv.Itoa(randomValue))
			}
		}
	}()

	defer ticker.Stop()

	fmt.Println("started")

	sig := <-sigs
	fmt.Printf("event: %s\n", sig)
	fmt.Println("exiting")
}
