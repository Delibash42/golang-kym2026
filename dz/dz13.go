package main

import "fmt"

func countTypes(items []any) map[string]int {
	result := make(map[string]int)
	for _, v := range items {
		switch v.(type) {
		case int:
			result["int"]++
		case string:
			result["string"]++
		default:
			result["other"]++
		}
	}
	return result
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
	c := countTypes(data)
	for i, v := range c {
		fmt.Println(i, v)
	}
}
