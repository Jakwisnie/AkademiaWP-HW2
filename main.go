package main

import (
	"homework2/scripts"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		scripts.Web()

	}()
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		clientStart := time.Now()
		scripts.Client()
		clientElapsed := time.Since(clientStart)
		log.Printf("Client used for: %s\n", clientElapsed)
	}()

	wg.Wait()
}
