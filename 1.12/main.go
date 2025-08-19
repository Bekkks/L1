package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree", "dog"}
	m := make(map[string]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for l := range m {
		fmt.Println(l)
	}
}

/*Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

*/
