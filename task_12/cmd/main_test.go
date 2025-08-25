package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestUniqueWord(t *testing.T) {
	tests := []struct {
		testNum  int
		words    []string
		expected []string
	}{
		{1, []string{"cat", "cat", "dog", "cat", "tree"}, []string{"cat", "dog", "tree"}},
		{2, []string{"a", "b", "c", "a", "b"}, []string{"a", "b", "c"}},
		{3, []string{"one"}, []string{"one"}},
		{4, []string{}, []string{}},
		{5, []string{"x", "x", "x", "x"}, []string{"x"}},
		{6, []string{"go", "java", "python", "go"}, []string{"go", "java", "python"}},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		tc := tc
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := uniqueWord(tc.words)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("testNum = %d: res %v, want %v", tc.testNum, res, tc.expected)
			}
		}()
	}

	wg.Wait()

}
