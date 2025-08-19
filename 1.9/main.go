package main

import (
	"fmt"
)

func work1(ch chan int, arr [10]int) {
	for _, v := range arr {
		ch <- v
	}
	close(ch)
}

func work2(ch chan int, out chan int) {
	for v := range ch {
		out <- v * 2
	}
	close(out)
}

func main() {
	ch := make(chan int)
	out := make(chan int)
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	go work1(ch, arr)
	go work2(ch, out)
	for v := range out {
		fmt.Println(v)
	}
}

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается. */
