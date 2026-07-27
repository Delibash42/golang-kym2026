package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	wg.Add(1)
	go func() {
		i := 1
		defer close(ch)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("даватель закончился")
				return
			default:
				ch <- i
				i++
				time.Sleep(1 * (time.Second))
			}

		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("читатель закочился")
				return
			case i := <-ch:
				fmt.Println("засек", i)
			}
		}
	}()
	select {
	case <-time.After(3 * time.Second):
		cancel()
	}
	wg.Wait()
}
