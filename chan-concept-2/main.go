package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("START - %v\n", start.Format(time.RFC3339Nano))
	fmt.Println("------------------------------------------------------------------")

	// init pool worker
	initPool := 1000

	// read csv or source data
	messages := readMessages()

	wg := sync.WaitGroup{}

	chMessages := make(chan string)

	for i := 0; i < initPool; i++ {
		go worker(&wg, chMessages)
	}

	wg.Add(len(messages))
	for _, msg := range messages {
		chMessages <- msg
	}

	close(chMessages)

	wg.Wait()

	// resource limit :
	// 100 milicore cpu
	// 100 mb memory

	// expectation:
	// please using worker pool and the execution should n time faster than serial process

	// if not using concurrency : 1s * 10000 = 10000s -> 2,7777777778h
	fmt.Println("------------------------------------------------------------------")
	end := time.Now()
	fmt.Printf("END - %v. Time Since : %v\n", end.Format(time.RFC3339Nano), time.Since(start))
}

func worker(wg *sync.WaitGroup, chMessage chan string) {
	for msg := range chMessage {
		defer wg.Done()
		sendMessage(msg)
	}
}

func readMessages() []string {
	messages := []string{}
	max := 10000
	for i := 0; i < max; i++ {
		messages = append(messages, strings.Replace("hi {:username}xx!", "{:username}", fmt.Sprintf("username%d", i), 1))
	}
	return messages
}

func sendMessage(message string) error {
	fmt.Printf("START - SEND MESSAGE '%v' at %v\n", message, time.Now().Format(time.RFC3339Nano))
	time.Sleep(1 * time.Second)
	fmt.Printf("FINISH - SEND MESSAGE '%v' at %v\n", message, time.Now().Format(time.RFC3339Nano))
	return nil
}
