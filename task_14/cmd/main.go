package main

import (
	"fmt"
	"sync"
)

// Если нужно будет работать с большим количеством типов данных, то лучше перейти на reflect.TypeOf
func typeOfVar(wg *sync.WaitGroup, ch <-chan any) {
	defer wg.Done()
	for v := range ch {
		switch val := v.(type) {
		case int:
			fmt.Printf("Val %v is int", val)
		case string:
			fmt.Printf("Val %v is string", val)
		case bool:
			fmt.Printf("Val %v is bool", val)
		case chan int:
			fmt.Printf("Val %v is chan int", val)
		case chan bool:
			fmt.Printf("Val %v is chan bool", val)
		case chan string:
			fmt.Printf("Val %v is chan string", val)
		}
		fmt.Printf("\n")
	}
}

func main() {
	chInt := make(chan int)
	chStr := make(chan string)
	chBool := make(chan bool)

	mainCh := make(chan any)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go typeOfVar(&wg, mainCh)

	val := []any{1, "go", false, chInt, chStr, chBool}
	for _, v := range val {
		mainCh <- v
	}
	close(mainCh)
	wg.Wait()
}
