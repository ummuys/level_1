package main

import (
	"sync"
	"testing"
)

func TestBinSearch(t *testing.T) {
	tests := []struct {
		testNum  int
		arr      []int
		item     int
		expected int
	}{
		{1, []int{1, 3, 5, 7, 9}, 5, 2},
		{2, []int{1, 3, 5, 7, 9}, 1, 0},
		{3, []int{1, 3, 5, 7, 9}, 9, 4},
		{4, []int{1, 3, 5, 7, 9}, 2, -1},
		{5, []int{}, 10, -1},
	}

	wg := sync.WaitGroup{}

	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			res := binSearch(tc.arr, tc.item)
			if res != tc.expected {
				t.Errorf("testNum = %d: res = %d, expected = %d", tc.testNum, res, tc.expected)
			}
		}()
	}

	wg.Wait()
}
