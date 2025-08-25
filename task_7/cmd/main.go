package main

import (
	"fmt"
	"strconv"
	"sync"
)

type mutexMap struct {
	mutex sync.Mutex
	info  map[string]string
}

func newMutexMap() *mutexMap {
	return &mutexMap{
		info: make(map[string]string),
	}
}

func (m *mutexMap) Set(key, value string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.info[key] = value
}

func main() {
	m := newMutexMap()
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := strconv.Itoa(i)
			value := fmt.Sprintf("value number %d", i)
			m.Set(key, value)
		}(i)
	}
	wg.Wait()
}
