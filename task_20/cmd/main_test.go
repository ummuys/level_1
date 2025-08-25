package main

import (
	"sync"
	"testing"
)

func TestReverseStr(t *testing.T) {
	tests := []struct {
		testNum  int
		input    string
		expected string
	}{
		{1, "", ""},
		{2, "word", "word"},
		{3, "snow dog sun", "sun dog snow"},
		{4, "go is fun", "fun is go"},
		{5, "a b c d", "d c b a"},
		{6, "hello, world!", "world! hello,"},
		{7, "привет мир", "мир привет"},
		{8, "🙂🙃 👋", "👋 🙂🙃"},
		{9, "one two", "two one"},
		{10, "a longword", "longword a"},
	}

	wg := sync.WaitGroup{}

	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			res := reverseStr(tc.input)
			if res != tc.expected {
				t.Errorf("testNum = %d: got = %q, expected = %q",
					tc.testNum, res, tc.expected)
			}
		}()
	}

	wg.Wait()
}
