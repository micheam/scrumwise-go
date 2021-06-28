package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/micheam/wiseman"
)

var (
	interval = flag.Int("interval", 5, "the interval to check for updates in seconds")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		s := <-c
		log.Println("Got Signal: ", s)
		cancel()
	}()

	// Server start
	err := wiseman.Work(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
