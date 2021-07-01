package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/micheam/wiseman/cache"
	"github.com/micheam/wiseman/scrumwise"
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

	chache := &cache.DataGateway{
		Interval:       time.Duration(*interval) * time.Second,
		GetDataVersion: scrumwise.GetDataVersion,
		GetData: func(ctx context.Context) (*scrumwise.Data, scrumwise.DataVersion, error) {
			param := scrumwise.NewGetDataParam("196595-12667-68")
			r, err := scrumwise.GetData(ctx, *param)
			if err != nil {
				return nil, 0, err
			}
			return r.Data, r.DataVersion, nil
		},
	}
	chache.StartTick(ctx)

	<-ctx.Done()
}
