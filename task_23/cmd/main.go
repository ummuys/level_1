package main

import "fmt"

func removeAt[T any](slice []T, i int) ([]T, error) {
	if i >= len(slice) || i < 0 {
		// out of range
		return slice, fmt.Errorf("index out of range")
	}

	if slice == nil {
		return slice, fmt.Errorf("slice is nil")
	}

	copy(slice[i:], slice[i+1:])

	// Важно!!! Если у нас будет слайс указателей, то
	var empty T
	slice[len(slice)-1] = empty

	return slice[:len(slice)-1], nil
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(removeAt(slice, 1))
}
