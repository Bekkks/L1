package main

import (
	"fmt"
)

func reverseWordsInPlace(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			for j := start; j < (start+i)/2; j++ {
				s[j], s[start+i-1-(j-start)] = s[start+i-1-(j-start)], s[j]
			}
			start = i + 1
		}
	}
}

func main() {
	str := "snow dog sun"
	runes := []rune(str)
	reverseWordsInPlace(runes)
	fmt.Println(string(runes)) // Вывод: "sun dog snow"
}
