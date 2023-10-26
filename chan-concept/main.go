package main

import (
	"fmt"
)

func main() {
	numTasks := 11
	results := make(chan int, numTasks)

	for i := 0; i < numTasks; i++ {
		go func(id int) {
			square := id * 2
			results <- square
		}(i)
	}

	for i := 0; i < numTasks; i++ {
		result := <-results
		fmt.Printf("Hasil ke-%d: %d\n", i, result)
	}
	close(results)
}
