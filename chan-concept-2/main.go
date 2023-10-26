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

	wg := sync.WaitGroup{}

	chMessages := make(chan string)

	messages := readMessages()

	poolSize := len(messages)

	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go worker(&wg, chMessages)
	}

	go func() {
		for _, msg := range messages {
			chMessages <- msg
		}
		close(chMessages)
	}()

	wg.Wait()

	fmt.Println("------------------------------------------------------------------")
	end := time.Now()
	fmt.Printf("END - %v. Time Since: %v\n", end.Format(time.RFC3339Nano), time.Since(start))
}

func worker(wg *sync.WaitGroup, chMessage chan string) {
	for msg := range chMessage {
		sendMessage(msg)
		wg.Done()
	}
}

func readMessages() []string {
	messages := []string{}
	max := 20
	for i := 0; i < max; i++ {
		messages = append(messages, strings.Replace("hi {:username}!", "{:username}", fmt.Sprintf("username-%d", i), 1))
	}
	return messages
}

func sendMessage(message string) {
	fmt.Printf("START - SEND MESSAGE '%v' at %v\n", message, time.Now().Format(time.RFC3339Nano))
	time.Sleep(10 * time.Second)
	fmt.Printf("FINISH - SEND MESSAGE '%v' at %v\n", message, time.Now().Format(time.RFC3339Nano))
}
