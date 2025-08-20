package main

import (
	"fmt"
	"sync"
)

type Inc struct {
	i int
}

var I Inc

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			I.i++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(I.i)
}
