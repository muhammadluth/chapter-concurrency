package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("START - %v\n", start.Format(time.RFC3339Nano))
	fmt.Println("------------------------------------------------------------------")
	increment := 0
	max := 10

	wg := sync.WaitGroup{}

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			increment += i
			worker(i)
		}(i)
	}
	wg.Wait()
	// 0+1+2+3+4+5+6+7+8+9 = 45

	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("TOTAL ID INCREMENT = %v\n", increment) // expected 45
	end := time.Now()
	fmt.Printf("END - %v. Time Since : %v\n", end.Format(time.RFC3339Nano), time.Since(start))
}

func worker(id int) {
	fmt.Printf("START - WORKER %v at %v\n", id, time.Now().Format(time.RFC3339Nano))
	time.Sleep(1 * time.Second)
	fmt.Printf("FINISH - WORKER %v at %v\n", id, time.Now().Format(time.RFC3339Nano))
}
