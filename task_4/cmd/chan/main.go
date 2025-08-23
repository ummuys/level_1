package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
)

func SomeFunc(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-done
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})
	wg := sync.WaitGroup{}
	defer cancel()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go SomeFunc(done, &wg)
	}

	<-ctx.Done()
	close(done)
	wg.Wait()
}
