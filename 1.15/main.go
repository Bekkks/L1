package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createHugeString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	result := make([]rune, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v[:100]))
}

func main() {
	someFunc()
	for v := range justString {
		fmt.Print(string(v))
	}
}

/*в том коде была представлена одна из частых ошибок в го, так как в го строки неизменяемы, но при присваевании juststring
мы ссылаемся на начальную строку и сборщик мусора не может очистить память, и в памяти так и остается вся
 строка сохраняется хотя нам нужно только первые 100 символов*/
