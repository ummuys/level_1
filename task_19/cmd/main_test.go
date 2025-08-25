package main

import (
	"sync"
	"testing"
)

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		testNum  int
		input    string
		expected string
	}{
		{1, "", ""},
		{2, "a", "a"},
		{3, "ab", "ba"},
		{4, "abcde", "edcba"},
		{5, "abcd", "dcba"},
		{6, "a b c", "c b a"},
		{7, "привет", "тевирп"},
		{8, "🙂🙃", "🙃🙂"},
		{9, "hi 👋", "👋 ih"},
		{10, "⚽🔥🏆", "🏆🔥⚽"},
	}

	wg := sync.WaitGroup{}

	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			r := []rune(tc.input)
			reverse(r)
			res := string(r)
			if res != tc.expected {
				t.Errorf("testNum = %d: got = %q, expected = %q",
					tc.testNum, res, tc.expected)
			}
		}()
	}

	wg.Wait()
}
