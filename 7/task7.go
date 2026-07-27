package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	mar := make(map[string]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			mar[fmt.Sprintf("key%d", n)] = n
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	mu.Lock()
	for k, v := range mar {
		fmt.Printf("%s: %d\n", k, v)
	}
	mu.Unlock()
}
