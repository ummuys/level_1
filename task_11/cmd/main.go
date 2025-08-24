package main

import (
	"fmt"
	"sort"
)

func SameNums(arr1, arr2 []int) []int {

	map1 := make(map[int]struct{})
	sNums := make(map[int]struct{})
	for _, num := range arr1 {
		map1[num] = struct{}{}
	}

	for _, num := range arr2 {
		if _, ok := map1[num]; ok {
			sNums[num] = struct{}{}
		}
	}

	res := make([]int, 0, len(sNums))

	for key := range sNums {
		res = append(res, key)
	}

	sort.Ints(res)

	return res
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 3, 4}
	fmt.Println(SameNums(arr1, arr2))
}
