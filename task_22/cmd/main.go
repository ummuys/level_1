package main

import (
	"fmt"
	"math/big"
)

func addBI(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subBI(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mulBI(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divBI(a, b *big.Int) (*big.Int, error) {
	if b.Sign() == 0 {
		return nil, fmt.Errorf("divide by zero")
	}
	return new(big.Int).Quo(a, b), nil
}

func main() {
	a := new(big.Int).Lsh(big.NewInt(1), 128)
	b, ok := new(big.Int).SetString("12345678900198273409876123948761978236498716234987612398476918236749871263498761234", 10)
	if !ok {
		panic("can't create a bigInt froms tr")
	}

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	fmt.Println("a + b =", addBI(a, b))
	fmt.Println("a - b =", subBI(a, b))
	fmt.Println("a * b =", mulBI(a, b))

	z, _ := divBI(a, b)
	fmt.Println("a / b =", z)
}
