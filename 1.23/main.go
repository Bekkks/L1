package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s int
	fmt.Scan(&s)
	copy(arr[s:], arr[s+1:])
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}
