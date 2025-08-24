package main

import (
	"fmt"
	"sync"
	"time"
)

func WriteData(wg *sync.WaitGroup, arr []int, in chan<- int) {
	defer wg.Done()

	for _, num := range arr {
		in <- num
		time.Sleep(100 * time.Millisecond)
	}
	close(in)
}

func ReadData(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()

	for data := range in {
		out <- data * 2
	}
	close(out)
}

func main() {
	wg := sync.WaitGroup{}
	in := make(chan int)
	out := make(chan int)
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	wg.Add(1)
	go WriteData(&wg, arr, in)

	wg.Add(1)
	go ReadData(&wg, in, out)

	for data := range out {
		fmt.Println(data)
	}

	wg.Wait()
}
