package main

import (
	"fmt"
	"sync"
)

func kvadr(i int) {
	fmt.Println(i * i)
}
func main() {
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}
	for i, value := range arr {
		fmt.Println(i, value)
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			kvadr(g)
		}(value)
	}
	wg.Wait()
}
