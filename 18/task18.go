package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	count int
}

var wg sync.WaitGroup

func main() {
	C := counter{count: 0}

	for i := 1; i <= 5; i++ {
		go func(b int) {
			wg.Add(1)
			defer wg.Done()
			for j := 1; j <= b; j++ {
				C.count += b
				fmt.Println(C.count)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println(C.count)
}
