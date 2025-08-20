package main

import (
	"fmt"
	"sync"
)

type SqrtNum struct {
	num, pos int
}

func Worker(wg *sync.WaitGroup, in <-chan SqrtNum, out chan<- SqrtNum) {
	defer wg.Done()
	for s := range in {
		out <- SqrtNum{s.num * s.num, s.pos}
	}
}

func ArrSqrt(arr []int, cWorkers int) []int {
	in := make(chan SqrtNum)
	out := make(chan SqrtNum)
	wg := sync.WaitGroup{}
	res := make([]int, len(arr)) // Для предсказуемости, но можно мутировать и исходный массив

	for i := 0; i < cWorkers; i++ {
		wg.Add(1)
		go Worker(&wg, in, out)
	}

	// Нужно запустить в рутине, т.к без нее схватим deadlock: зависним в Worker в ожидании close(in)
	go func() {
		for i, num := range arr {
			in <- SqrtNum{num, i}
		}
		close(in)
	}()

	// Нужно запустить в рутине, т.к без нее схватим deadlock: зависним в цикле range out в ожидании close(out)
	go func() {
		wg.Wait()
		close(out)
	}()

	for s := range out {
		res[s.pos] = s.num
	}

	return res
}

func main() {
	fmt.Println(ArrSqrt([]int{2, 4, 6, 8, 10}, 3))
}
