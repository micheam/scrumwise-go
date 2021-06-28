package wiseman

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/micheam/wiseman/scrumwise"
)

type Datasource struct {
	data *scrumwise.Data
	mux  sync.Mutex
}

func Work(ctx context.Context) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("done from ctx")
			return nil
		case t := <-ticker.C:
			fmt.Println(t.Format(time.RFC3339))
			continue
		}
	}

}
