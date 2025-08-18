/*
	Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var a = []int{2, 4, 6, 8, 10}
	for _, x := range a {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			fmt.Println(a * a)
		}(x)
	}
	wg.Wait()
}
