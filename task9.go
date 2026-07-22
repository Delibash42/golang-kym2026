package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for _, n := range nums {
			ch1 <- n
		}
	}()

	go func() {
		defer close(ch2)
		for n := range ch1 {
			ch2 <- n * 2
		}
	}()

	for n := range ch2 {
		fmt.Println(n)
	}
}
