package main

import "fmt"

func reverse(arr []rune) {
	if len(arr) < 2 {
		return
	}

	l := 0
	r := len(arr) - 1
	for l < r {
		arr[r], arr[l] = arr[l], arr[r]
		l++
		r--
	}
}

func main() {
	str := []rune("hello")
	reverse(str)
	fmt.Println(string(str))
}
