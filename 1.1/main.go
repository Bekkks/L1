package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	play string
	Human
}

func (h *Human) hello() {
	fmt.Println("hello ", h.name)
}

func (h *Human) per() {
	fmt.Println("age ", h.age)
}

func main() {
	a := Action{
		play: "hello",
		Human: Human{
			name: "Bek",
			age:  20,
		},
	}
	a.hello()
	a.per()
}
