package main

import (
	"fmt"
	"sync"
)

type sqrtNum struct {
	num, pos int
}

func worker(wg *sync.WaitGroup, in <-chan sqrtNum, out chan<- sqrtNum) {
	defer wg.Done()
	for s := range in {
		out <- sqrtNum{s.num * s.num, s.pos}
	}
}

func arrsqrt(arr []int, cWorkers int) []int {
	in := make(chan sqrtNum)
	out := make(chan sqrtNum)
	wg := sync.WaitGroup{}
	res := make([]int, len(arr)) // Для предсказуемости, но можно мутировать и исходный массив

	for i := 0; i < cWorkers; i++ {
		wg.Add(1)
		go worker(&wg, in, out)
	}

	// Нужно запустить в рутине, т.к без нее схватим deadlock: зависним в Worker в ожидании close(in)
	go func() {
		for i, num := range arr {
			in <- sqrtNum{num, i}
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
	fmt.Println(arrsqrt([]int{2, 4, 6, 8, 10}, 3))
}
