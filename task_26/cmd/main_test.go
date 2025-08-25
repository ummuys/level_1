package main

import (
	"sync"
	"testing"
)

func TestUniqueSymb(t *testing.T) {
	tests := []struct {
		testNum  int
		input    string
		expected bool
	}{
		{1, "", true},
		{2, "abcd", true},
		{3, "abCdefAaf", false},
		{4, "aabcd", false},
		{5, "Aa", false},
		{6, "GoLang", false},
		{7, "abcdefg", true},
		{8, "TheQuickBrownFox", false},
		{9, "HelloWorld", false},
		{10, "no duplicates!", true},
		{11, "1234567890", true},
		{12, "12341234", false},
		{13, "!@#$%^&*()", true},
		{14, "!@#!!@#", false},
	}

	wg := sync.WaitGroup{}

	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			res := uniqueSymb(tc.input)
			if res != tc.expected {
				t.Errorf("testNum = %d: got = %v, expected = %v",
					tc.testNum, res, tc.expected)
			}
		}()
	}

	wg.Wait()
}
