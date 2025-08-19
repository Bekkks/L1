package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func work1(p map[string]int) map[string]int {
	mu.Lock()
	defer mu.Unlock()
	p["a"] = 2
	p["b"] = 3
	return p
}

func work2(p map[string]int) map[string]int {
	mu.Lock()
	defer mu.Unlock()
	p["c"] = 4
	p["d"] = 5
	return p
}

func main() {
	pers := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		work1(pers)
	}()
	go func() {
		defer wg.Done()
		work2(pers)
	}()
	wg.Wait()
	for k, v := range pers {
		fmt.Println(k, v)
	}
}
