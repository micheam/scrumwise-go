package cache

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/micheam/wiseman/scrumwise"
)

type DataGateway struct {
	Interval       time.Duration
	GetDataVersion GetDataVersionFn
	GetData        GetDataFn

	mux         sync.RWMutex
	currVersion scrumwise.DataVersion
	currData    *scrumwise.Data
}

type (
	GetDataVersionFn func(ctx context.Context) (scrumwise.DataVersion, error)
	GetDataFn        func(ctx context.Context) (*scrumwise.Data, scrumwise.DataVersion, error)
)

func (r *DataGateway) Current() (scrumwise.DataVersion, *scrumwise.Data) {
	r.mux.RLock()
	defer r.mux.RUnlock()
	if r.currData == nil {
		return -1, nil
	}
	return r.currVersion, r.currData
}

// Init load initial data from scrumwise, then return result.
//
// This will block gorutine until initialization is complete.
// If the data has already been loaded, it will skip
// the process and return the flow to the caller.
func (r *DataGateway) Init(ctx context.Context) error {
	log.Println("Init...")
	if v, _ := r.Current(); v > 0 {
		log.Println("already initialized. Skip Init()")
		return nil
	}
	r.mux.Lock()
	defer r.mux.Unlock()

	d, v, err := r.GetData(ctx)
	if err != nil {
		return err
	}
	r.currVersion, r.currData = v, d
	log.Println("initialized.")
	return nil
}

func (r *DataGateway) StartTick(ctx context.Context) {
	log.Println("start...")
	ticker := time.NewTicker(r.Interval)
	go func() {
		err := r.Init(ctx)
		if err != nil {
			panic(err) // TODO: broadcasts error occurrence
		}
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				v, err := r.GetDataVersion(ctx)
				if err != nil {
					log.Println("[Cacher] failed to get-version: ", err)
					continue
				}
				if v <= r.currVersion {
					continue
				}
				log.Println("try to refresh data ...")
				r.mux.Lock()
				d, v, err := r.GetData(ctx)
				if err != nil {
					log.Println("failed to get-data: ", err)
					continue
				}
				r.currVersion, r.currData = v, d
				log.Println("data refreshed with data-version: ", r.currVersion)
				r.mux.Unlock()
				continue
			}
		}
	}()
}
