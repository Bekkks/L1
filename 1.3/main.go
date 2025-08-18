/*Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров. */

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(ch <-chan int, id int) {
	for n := range ch {
		fmt.Printf("Worker %d got %d\n", id, n)
	}
}

func main() {
	ch := make(chan int)
	count := 1
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < num; i++ {
		go worker(ch, i)
	}

	for {
		ch <- count
		count++
		time.Sleep(1 * time.Second)
	}
}
