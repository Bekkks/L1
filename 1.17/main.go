package main

import (
	"fmt"
	"sort"
)

func Get(arr []int, a int) int {
	i := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= a
	})
	if i < len(arr) && arr[i] == a {
		return i
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Get(arr, 5))  // 4
	fmt.Println(Get(arr, 11)) // -1
}
