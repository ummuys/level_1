package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu sync.Mutex
	c  int
}

// Или просто atomic.AddInt64(&c, 1)

func (co *counter) Inc() {
	co.mu.Lock()
	defer co.mu.Unlock()
	co.c++
}

func main() {
	co := counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			co.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(co.c)
}
