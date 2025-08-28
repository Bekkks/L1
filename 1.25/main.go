package main

import (
	"time"
)

func Sleep(t time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= t {
			break
		}
	}
}

func main() {
	Sleep(3 * time.Second)
}
