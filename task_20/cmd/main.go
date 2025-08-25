package main

import "fmt"

func reverse(arr []rune, l, r int) {
	for l < r {
		arr[r], arr[l] = arr[l], arr[r]
		l++
		r--
	}
}
func reverseStr(words string) string {
	arr := []rune(words)
	if len(arr) < 1 {
		return words
	}
	reverse(arr, 0, len(arr)-1)

	start := 0
	for i := 0; i <= len(arr); i++ {
		if i == len(arr) || arr[i] == ' ' {
			reverse(arr, start, i-1)
			start = i + 1
		}
	}

	return string(arr)
}

func main() {
	fmt.Println(reverseStr("hello i am eugene"))
}
