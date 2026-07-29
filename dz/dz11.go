package main

import "fmt"

func sumInts(items []any) int {
	sum := 0
	for _, v := range items {
		if val, ok := v.(int); ok {
			sum += val
		}
	}
	return sum
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
	fmt.Println(sumInts(data))
}
