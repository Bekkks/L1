package main

import (
	"fmt"
)

func reverse(a []rune) []rune {
	var res []rune
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}
	return res
}

func main() {
	s := "главрыба 😊🚀"
	runes := []rune(s)
	fmt.Println(string(reverse(runes))) // вывод: "🚀😊 абырвалг"
}
