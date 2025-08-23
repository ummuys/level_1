package main

import (
	"fmt"
	"strconv"
	"sync"
)

type MutexMap struct {
	mutex sync.Mutex
	info  map[string]string
}

func NewMutexMap() *MutexMap {
	return &MutexMap{
		info: make(map[string]string),
	}
}

func (m *MutexMap) Set(key, value string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.info[key] = value
}

func main() {
	m := NewMutexMap()
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
