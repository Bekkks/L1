package main

import (
	"fmt"
	"strings"
)

func work(str string) bool {
	str = strings.ToLower(str)
	for _, i := range str {
		if strings.Count(str, string(i)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(work("helo WiRD"))
}
