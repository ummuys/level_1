package main

import (
	"fmt"
	"sort"
)

func uniqueWord(words []string) []string {
	uw := make(map[string]struct{})
	for _, w := range words {
		if _, ok := uw[w]; !ok {
			uw[w] = struct{}{}
		}
	}

	res := []string{}

	for key := range uw {
		res = append(res, key)
	}

	sort.Strings(res)
	return res
}

func main() {
	fmt.Println(uniqueWord([]string{"cat", "cat", "dog", "cat", "tree"}))
}
