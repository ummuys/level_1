package main

import (
	"fmt"
)

func groupTemp(temps []float64) map[int][]float64 {
	group := make(map[int][]float64)
	for _, temp := range temps {
		key := int(temp) / 10 * 10
		group[key] = append(group[key], temp)
	}
	return group
}
func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(groupTemp(temp))
}
