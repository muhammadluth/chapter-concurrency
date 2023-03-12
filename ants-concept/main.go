package main

import (
	"ants-auto-tune-concept/model"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var runningTasks int32

func main() {
	start := time.Now()
	fmt.Printf("START - %v\n", start.Format(time.RFC3339Nano))
	fmt.Println("------------------------------------------------------------------")
	config := model.Config{
		PoolSize:         50,
		MinPoolSize:      10,
		MaxPoolSize:      200,
		PoolIncrement:    20,
		AutoTuneDuration: 10 * time.Millisecond,
	}

	runTimes := 1000

	pool, err := ants.NewPool(config.PoolSize)
	if err != nil {
		panic(err)
	}

	// Use the common pool.
	wg := sync.WaitGroup{}

	go AutoTune(pool, config)

	for i := 0; i < runTimes; i++ {
		pool.Submit(
			func() {
				sendMessage(i, &wg)
			},
		)
		defer pool.Release()
	}

	fmt.Println("------------------------------------------------------------------")
	end := time.Now()
	fmt.Printf("END - %v. Time Since : %v\n", end.Format(time.RFC3339Nano), time.Since(start))
	wg.Wait()
}

func AutoTune(pool *ants.Pool, config model.Config) {
	for {
		runTasks := atomic.LoadInt32(&runningTasks)
		if int(runTasks) > (pool.Cap()-1) &&
			pool.Cap()+config.PoolIncrement < (config.MaxPoolSize+1) {
			// tune pool
			newPool := pool.Cap() + config.PoolIncrement
			pool.Tune(newPool)
			fmt.Println("=================================================")
			fmt.Printf("TUNE CAPACITY POOL TO: %v\n", pool.Cap())
			fmt.Println("=================================================")
		} else if pool.Cap()-config.PoolIncrement > (config.MinPoolSize - 1) {
			// tune pool
			newPool := pool.Cap() - config.PoolIncrement
			pool.Tune(newPool)
			fmt.Println("=================================================")
			fmt.Printf("TUNE CAPACITY POOL TO: %v\n", pool.Cap())
			fmt.Println("=================================================")
		}
		fmt.Printf("RUNNING TASKS: %v\n", runTasks)
		time.Sleep(config.AutoTuneDuration)
	}
}

func sendMessage(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&runningTasks, +1)
	wg.Add(1)
	fmt.Printf("START - SEND MESSAGE '%v' at %v\n", id, time.Now().Format(time.RFC3339Nano))
	time.Sleep(1 * time.Second)
	fmt.Printf("FINISH - SEND MESSAGE '%v' at %v\n", id, time.Now().Format(time.RFC3339Nano))
	atomic.AddInt32(&runningTasks, -1)
}
