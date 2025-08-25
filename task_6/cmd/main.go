package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 1
func exitCicle(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Cicle gorutine start")
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Cicle gorutine end")
}

// 2
func exitDoneChan(wg *sync.WaitGroup, exit <-chan struct{}) {
	defer wg.Done()
	fmt.Println("Done chan gorutine start")
	for {
		select {
		case <-exit:
			fmt.Println("Done chan gorutine end")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 3
func exitCloseChan(wg *sync.WaitGroup, ch <-chan struct{}) {
	defer wg.Done()
	fmt.Println("Close chan gorutine start")
	for range ch {
	}
	fmt.Println("Close chan gorutine end")
}

// 4
func exitCtx(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Ctx gorutine start")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx gorutine end")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 5
func exitTime(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Time gorutine start")
	timer := time.After(time.Second * 2)
	for {
		select {
		case <-timer:
			fmt.Println("Time gorutine end")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 6
func exitRuntime(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Runtime gorutine start")
	<-time.After(2 * time.Second)
	fmt.Println("Runtime gorutine end")
	runtime.Goexit()
}

// 7
func exitPanic(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Panic gorutine start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic gorutine end")
		}
	}()
	panic("hello!")
}

func main() {

	wg := sync.WaitGroup{}
	//1
	wg.Add(1)
	go exitCicle(&wg)

	//2
	done := make(chan struct{})
	wg.Add(1)
	go exitDoneChan(&wg, done)
	done <- struct{}{}

	//3
	ch := make(chan struct{})
	wg.Add(1)
	go exitCloseChan(&wg, ch)
	close(ch)

	//4
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go exitCtx(ctx, &wg)
	cancel()

	//5
	wg.Add(1)
	go exitTime(&wg)

	//6
	wg.Add(1)
	go exitRuntime(&wg)

	//7
	wg.Add(1)
	go exitPanic(&wg)

	wg.Wait()
}
