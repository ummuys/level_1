package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		testNum  int
		arr      []int
		expected []int
	}{
		{1, []int{3, 2, 1}, []int{1, 2, 3}},
		{2, []int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
		{3, []int{1, 1, 1}, []int{1, 1, 1}},
		{4, []int{}, []int{}},
		{5, []int{42}, []int{42}},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		tc := tc
		wg.Add(1)
		go func() {
			defer wg.Done()
			quickSort(tc.arr, 0, len(tc.arr)-1)
			if !reflect.DeepEqual(tc.arr, tc.expected) {
				t.Errorf("testNum = %d: res = %v, expected = %v", tc.testNum, tc.arr, tc.expected)
			}
		}()
	}

	wg.Wait()
}
