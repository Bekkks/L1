package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 80, 8, 9, 11}
	arr2 := [5]int{10, 80, 8, 6, 11}
	m := make(map[int]bool)
	for _, v := range arr {
		m[v] = true
	}
	for _, v := range arr2 {
		if m[v] {
			fmt.Println(v)
		}
	}
}
