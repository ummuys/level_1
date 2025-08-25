package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestSameNums(t *testing.T) {
	tests := []struct {
		arr1     []int
		arr2     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 4, 6}, []int{2, 3, 6}, []int{6}},
		{[]int{1, 1, 2, 2, 3}, []int{2, 2, 3, 3, 4}, []int{2, 3}},
		{[]int{-1, -2, 0, 3}, []int{-2, 3, 5}, []int{-2, 3}},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		tc := tc
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := SameNums(tc.arr1, tc.arr2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("arrs not equal: res %v, want %v", res, tc.expected)
			}
		}()
	}

	wg.Wait()
}
