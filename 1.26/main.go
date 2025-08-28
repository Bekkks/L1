package main

import (
	"fmt"
	"strings"
)

func work(str string) bool {
	str = strings.ToLower(str)
	ms := make(map[rune]bool)
	for _, r := range str {
		if ms[r] {
			return false
		}
		ms[r] = true
	}
	return true
}

func main() {
	fmt.Println(work("helo WiRD"))
}
