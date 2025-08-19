package main

import "fmt"

func Test(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("Invalid type")
	}
}

func main() {
	var intVar interface{} = 42
	var strVar interface{} = "Hello"
	var boolVar interface{} = true
	var intChanVar interface{} = make(chan int)
	var unknownVar interface{} = float64(3.14)
	Test(intVar)
	Test(strVar)
	Test(boolVar)
	Test(intChanVar)
	Test(unknownVar)
}
