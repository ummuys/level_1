package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(wg *sync.WaitGroup, out <-chan struct{}) {
	defer wg.Done()

	for range out {
		fmt.Printf("read out\n")
	}

}

func writer(wg *sync.WaitGroup, in chan<- struct{}, seconds time.Duration) {
	defer wg.Done()
	timer := time.After(time.Second * seconds)

	for {
		select {
		case <-timer:
			close(in)
			return
		case in <- struct{}{}:
			fmt.Printf("write in\n")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{})

	wg.Add(1)
	go reader(&wg, ch)

	wg.Add(1)
	go writer(&wg, ch, 5)

	wg.Wait()
}
