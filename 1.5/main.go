package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for v := range 100 {
			select {
			case ch <- v:
				time.Sleep(500 * time.Millisecond)
			case <-done:
				return
			}
		}
		close(ch)
	}()
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	select {
	case <-time.After(time.Duration(num) * time.Second):
		fmt.Println("timeout")
		close(done)
		return
	}
}
