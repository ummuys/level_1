package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestUniqueWord(t *testing.T) {
	tests := []struct {
		words    []string
		expected []string
	}{
		{[]string{"cat", "cat", "dog", "cat", "tree"}, []string{"cat", "dog", "tree"}},
		{[]string{"a", "b", "c", "a", "b"}, []string{"a", "b", "c"}},
		{[]string{"one"}, []string{"one"}},
		{[]string{}, []string{}},
		{[]string{"x", "x", "x", "x"}, []string{"x"}},
		{[]string{"go", "java", "python", "go"}, []string{"go", "java", "python"}},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		tc := tc
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := UniqueWord(tc.words)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("arrs not equal: res %v, want %v", res, tc.expected)
			}
		}()
	}

	wg.Wait()

}
