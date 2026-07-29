package main

import "fmt"

func describeAll(items []any) {
	for _, v := range items {
		switch val := v.(type) {
		case int:
			fmt.Printf("int: %d\n", val)
		case string:
			fmt.Printf("string: %q\n", val)
		case bool:
			fmt.Printf("bool: %t\n", val)
		case float64:
			fmt.Printf("float64: %f\n", val)
		default:
			fmt.Printf("другой тип: %v\n", val)
		}
	}
}

func main() {
	data := []any{
		42,
		"hello",
		true,
		make(chan int),
		3.14,
		"world",
		100,
		false,
		[]int{1, 2, 3},
	}
	describeAll(data)
}
