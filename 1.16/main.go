package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	var less, greater []int
	for i, v := range arr {
		if i == len(arr)/2 {
			continue
		}
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	quickLess := quickSort(less)
	quickGreater := quickSort(greater)
	result := make([]int, 0, len(quickLess)+1+len(quickGreater))
	result = append(result, quickLess...)
	result = append(result, pivot)
	result = append(result, quickGreater...)
	return result
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4, 3}
	fmt.Println(quickSort(arr))
}
