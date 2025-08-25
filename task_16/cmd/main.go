package main

import "fmt"

func quickSort(arr []int, left, right int) {
	if right < left {
		return
	}

	mid := arr[(left+right)/2]

	i, j := left, right

	for i <= j {
		for arr[j] > mid {
			j--
		}
		for arr[i] < mid {
			i++
		}
		if j >= i {
			arr[i], arr[j] = arr[j], arr[i]
		}
		j--
		i++
	}

	quickSort(arr, left, j)
	quickSort(arr, i, right)
}

func main() {
	arr := []int{7, 2, 3}
	quickSort(arr, 0, 2)
	fmt.Println(arr)
}
