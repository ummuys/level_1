package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestSameNums(t *testing.T) {
	tests := []struct {
		testNum  int
		arr1     []int
		arr2     []int
		expected []int
	}{
		{1, []int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{2, []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{3, []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{4, []int{}, []int{1, 2, 3}, []int{}},
		{5, []int{}, []int{}, []int{}},
		{6, []int{1, 4, 6}, []int{2, 3, 6}, []int{6}},
		{7, []int{1, 1, 2, 2, 3}, []int{2, 2, 3, 3, 4}, []int{2, 3}},
		{8, []int{-1, -2, 0, 3}, []int{-2, 3, 5}, []int{-2, 3}},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		tc := tc
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := sameNums(tc.arr1, tc.arr2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("testNum = %d: res %v, want %v", tc.testNum, res, tc.expected)
			}
		}()
	}

	wg.Wait()
}
