package main

import (
	"fmt"
)

// Использовал bool т.к хотел именно бит, а bool принимает лишь 0 или 1

func SetBit(num int64, pos uint64, bit bool) int64 {
	var res int64

	if bit {
		res = num | (1 << pos)
	} else {
		res = num &^ (1 << pos)
	}
	return res
}

func main() {
	// 15 = 01111, 7 = 00111
	fmt.Println(SetBit(15, 3, false)) // ожидаем 7

	// 5 = 0101, 4 = 0100
	fmt.Println(SetBit(5, 0, false)) // ожидаем 4

	// 5 = 0101
	fmt.Println(SetBit(5, 2, true)) // ожидаем 5

	//1 = 0001, 5 = 0101
	fmt.Println(SetBit(1, 2, true)) // ожидаем 5

	//0 = 00000, 4 = 10000
	fmt.Println(SetBit(0, 4, true)) // ожидаем 16
}
