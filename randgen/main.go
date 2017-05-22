package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/robfig/cron"
)

func main() {
	fmt.Println("Random Generator")

	r1 := rand.New(rand.NewSource(time.Now().Unix()))

	sigs := make(chan os.Signal, 1)

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		fmt.Printf("-> %d\n", r1.Intn(10000-1)+1)
	})
	c.Start()
	defer c.Stop()

	fmt.Println("started")

	sig := <-sigs
	fmt.Printf("event: %s\n", sig)
	fmt.Println("exiting")
}
