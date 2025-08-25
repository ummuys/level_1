package main

import (
	"fmt"
	"sync"
	"time"
)

func writeData(wg *sync.WaitGroup, arr []int, in chan<- int) {
	defer wg.Done()

	for _, num := range arr {
		in <- num
		time.Sleep(100 * time.Millisecond)
	}
	close(in)
}

func readData(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
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
	go writeData(&wg, arr, in)

	wg.Add(1)
	go readData(&wg, in, out)

	for data := range out {
		fmt.Println(data)
	}

	wg.Wait()
}
