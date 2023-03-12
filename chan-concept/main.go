package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("START - %v\n", start.Format(time.RFC3339Nano))
	fmt.Println("------------------------------------------------------------------")
	increment := 0
	max := 10
	done := make(chan int)

	for i := 0; i < max; i++ {
		go func(c chan int, i int) {
			increment += i
			worker(i)
			c <- 1
		}(done, i)
	}
	for i := 0; i < max; i++ {
		<-done
	}

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
