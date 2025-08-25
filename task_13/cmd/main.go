package main

import "fmt"

// Моя функция работает по такому принципу
// 1) Получаем основу из двух чисел с помощью основ: 5 ^ 1 = 4
// 2) Чтобы получить в первом числе, где было 5, нам надо: 4 ^ 5 = 1
// 3) Чтобы получить в первом числе, где было 1, нам надо: 4 ^ 1 = 5

func XORSwap(num1, num2 int) (int, int) {
	num1 = num2 ^ num1
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	return num1, num2
}

func main() {
	fmt.Println(XORSwap(5, 1))
}
