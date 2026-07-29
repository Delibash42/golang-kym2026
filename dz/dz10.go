package main

import "fmt"

func typei(inter interface{}) string {
	switch inter.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		return "hz"
	}
}

func main() {
	fmt.Println(typei(42))
	fmt.Println(typei("42"))
	fmt.Println(typei(true))
	fmt.Println(typei(make(chan int)))
	fmt.Println(typei(make(chan string)))
	fmt.Println(typei(make(chan bool)))

	fmt.Println(typei(42.0))

}
