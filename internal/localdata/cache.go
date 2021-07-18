package localdata

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/micheam/wiseman/scrumwise"
)

type DataCache struct {
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

func (r *DataCache) Current() *scrumwise.Data {
	r.mux.RLock()
	defer r.mux.RUnlock()
	if r.currData == nil {
		return nil
	}
	return r.currData
}

func (r *DataCache) CurrentVersion() scrumwise.DataVersion {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.currVersion
}

// Init load initial data from scrumwise, then return result.
//
// This will block gorutine until initialization is complete.
// If the data has already been loaded, it will skip
// the process and return the flow to the caller.
func (r *DataCache) Init(ctx context.Context) error {
	log.Println("Init...")
	if v := r.CurrentVersion(); v > 0 {
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

func (r *DataCache) StartTick(ctx context.Context) {
	panic("Not Implemented Yet")
	// log.Println("start...")
	// ticker := time.NewTicker(r.Interval)
	// go func() {
	// 	err := r.Init(ctx)
	// 	if err != nil {
	// 		panic(err) // TODO: broadcasts error occurrence
	// 	}
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			ticker.Stop()
	// 			return
	// 		case <-ticker.C:
	// 			v, err := r.GetDataVersion(ctx)
	// 			if err != nil {
	// 				log.Println("[Cacher] failed to get-version: ", err)
	// 				continue
	// 			}
	// 			if v <= r.currVersion {
	// 				continue
	// 			}
	// 			log.Println("try to refresh data ...")
	// 			r.mux.Lock()
	// 			d, v, err := r.GetData(ctx)
	// 			if err != nil {
	// 				log.Println("failed to get-data: ", err)
	// 				continue
	// 			}
	// 			r.currVersion, r.currData = v, d
	// 			log.Println("data refreshed with data-version: ", r.currVersion)
	// 			r.mux.Unlock()
	// 			continue
	// 		}
	// 	}
	// }()
}
