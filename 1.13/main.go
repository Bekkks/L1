package main

import "fmt"

func main() {
	var a, b = 1, 2
	fmt.Println(a, b)
	a, b = b, a

	a = a ^ b
	b = a ^ b
	a = a ^ b

	a = a + b
	b = a - b
	a = a - b

	//здесь представлено 3 способа выполнения задачи
	fmt.Println(a, b)
}
