package main

import (
	"fmt"
	"sync"
)

func arrSqrt(arr []int) []int {
	wg := sync.WaitGroup{}
	res := make([]int, len(arr)) // Для предсказуемости, но можно мутировать и исходный массив
	for i, num := range arr {
		wg.Add(1)
		go func(i, num int) {
			defer wg.Done()
			res[i] = num * num
		}(i, num)
	}
	wg.Wait()
	return res
}

func main() {
	fmt.Println(arrSqrt([]int{2, 4, 6, 8, 10}))
}
