package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
)

func someFunc(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	wg := sync.WaitGroup{}
	defer cancel()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go someFunc(ctx, &wg)
	}
	wg.Wait()
}
