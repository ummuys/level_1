package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func someFunc(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-done
}

func main() {
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, syscall.SIGINT)

	done := make(chan struct{})
	wg := sync.WaitGroup{}

	go func() {
		<-sCh
		close(done)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go someFunc(done, &wg)
	}
	wg.Wait()
}
