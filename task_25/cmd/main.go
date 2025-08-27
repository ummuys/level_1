package main

import "time"

func sleep(second uint) {
	timer := time.NewTimer(time.Second * time.Duration(second))
	<-timer.C
}

func main() {
	sleep(5)
}
