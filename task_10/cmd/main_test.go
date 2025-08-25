package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestGroupTemp(t *testing.T) {
	tests := []struct {
		testNum  int
		temp     []float64
		expected map[int][]float64
	}{
		{
			testNum: 1,
			temp:    []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5},
			expected: map[int][]float64{
				-20: {-25.4, -27.0, -21.0},
				10:  {13.0, 19.0, 15.5},
				20:  {24.5},
				30:  {32.5},
			},
		},
		{
			testNum: 2,
			temp:    []float64{1.1, 5.5, 9.9},
			expected: map[int][]float64{
				0: {1.1, 5.5, 9.9},
			},
		},
		{
			testNum: 3,
			temp:    []float64{-5.0, -15.9, 0.0, 7.7},
			expected: map[int][]float64{
				-10: {-15.9},
				0:   {-5.0, 0.0, 7.7},
			},
		},
		{
			testNum:  4,
			temp:     []float64{},
			expected: map[int][]float64{},
		},
		{
			testNum: 5,
			temp:    []float64{10.0, 19.9, -20.0, -29.9},
			expected: map[int][]float64{
				10:  {10.0, 19.9},
				-20: {-20.0, -29.9},
			},
		},
	}

	wg := sync.WaitGroup{}

	for _, tc := range tests {
		wg.Add(1)
		tc := tc // Это нужно для того, чтобы каждая рутина брала
		// именно свой тест. Затемняем в общем
		go func() {
			defer wg.Done()
			res := groupTemp(tc.temp)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("testNum = %d: res %v, want %v", tc.testNum, res, tc.expected)
			}
		}()
	}

	wg.Wait()

}
