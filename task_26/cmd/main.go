package main

import "fmt"

func uniqueSymb(word string) bool {
	uniq := make(map[byte]struct{})
	for i := 0; i < len(word); i++ {
		s := word[i]
		if _, ok := uniq[s]; ok {
			return false
		}
		uniq[s] = struct{}{}
		if s >= 'A' && s <= 'Z' {
			if _, ok := uniq[s+32]; ok {
				return false
			}
			uniq[s+32] = struct{}{}
		} else if s >= 'a' && s <= 'z' {
			if _, ok := uniq[s-32]; ok {
				return false
			}
			uniq[s-32] = struct{}{}
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueSymb("hello"))
}
