package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (a *Dog) Woof() {
	fmt.Println("Woof")
}

func MakeSpeak(s Speaker) {
	s.Speak()
}

type DogAdapter struct {
	dog *Dog
}

func (a *DogAdapter) Speak() {
	a.dog.Woof()
}

func main() {
	d := &Dog{}
	a := &DogAdapter{d}
	MakeSpeak(a)
}
