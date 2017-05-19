package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Printf("event: %s\n", sig)
		ticker.Stop()
		done <- true
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Broker")
			}
		}
	}()

	fmt.Println("started")
	<-done
	fmt.Println("exiting")
}
