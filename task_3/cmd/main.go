package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

func Worker(wg *sync.WaitGroup, out <-chan string, id int) {
	defer wg.Done()
	for d := range out {
		fmt.Printf("Worker #%d received: %s\n", id, d)
	}

}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cWorkers := flag.Int("workers", 3, "worker count")
	flag.Parse()

	info := make(chan string)
	wg := sync.WaitGroup{}
	for i := 0; i < *cWorkers; i++ {
		wg.Add(1)
		go Worker(&wg, info, i)
	}

	for {
		select {
		case <-ctx.Done():
			close(info)
			wg.Wait()
			return
		case info <- "go!":
		}

	}

}
