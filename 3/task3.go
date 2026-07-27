package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer close(ch)
		for i := 0; i < 11; i++ {
			ch <- i
		}
		defer wg.Done()
	}()
	for worker := 1; worker <= 3; worker++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for msg := range ch { // Цикл завершится, когда канал закроется
				fmt.Printf("работник %d засек %d\n", id, msg)
			}
		}(worker)

	}
	wg.Wait()
}
