package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	a_big := new(big.Int)
	b_big := new(big.Int)
	a_big.SetString(a, 10)
	b_big.SetString(b, 10)

	sum := new(big.Int).Add(a_big, b_big)
	fmt.Printf("Сумма:%s\n", sum)
	raz := new(big.Int).Sub(a_big, b_big)
	fmt.Printf("Разность:%s\n", raz)
	prod := new(big.Int).Mul(a_big, b_big)
	fmt.Printf("Произведение:%s\n", prod)
	r := new(big.Rat)
	r.SetFrac(a_big, b_big)
	fmt.Printf("Деление: %s\n", r)

}
