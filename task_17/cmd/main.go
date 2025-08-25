package main

import "fmt"

func binSearch(arr []int, item int) int {
	left := 0
	right := len(arr) - 1
	for right >= left {
		mid := (right + left) / 2
		if arr[mid] == item {
			return mid
		} else if arr[mid] > item {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
