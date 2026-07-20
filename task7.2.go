package main

import (
	"fmt"
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func main() {
	m := cmap.New[int]()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			m.Set(key, n*10)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := m.Get(key); ok {
			fmt.Printf("%s = %d\n", key, val)
		}
	}
}
