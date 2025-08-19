package main

import "fmt"

func Setbit(i int, bit int, num int64) int64 {
	if bit == 1 {
		return num | (1 << i)
	} else {
		return num &^ (1 << i)
	}
}

func main() {
	var i, bit int
	var num int64
	fmt.Scan(&num, &i, &bit)
	fmt.Println(Setbit(i, bit, num))
}
