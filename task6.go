package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("завершена по условию")
	}()

	wg.Add(1)
	die1 := make(chan bool)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-die1:
				fmt.Println("завершена через канал")
				return
			default:
				fmt.Println(i)
				i++
				if i >= 2 {
					die1 <- true
				}
			}
		}
	}()

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("завершена через контекст")
				return
			default:
				fmt.Println(i)
				i++
				if i >= 2 {
					cancel()
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("defer")
		for i := 0; i < 3; i++ {
			if i == 1 {
				runtime.Goexit()
			}
			fmt.Println(i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			fmt.Println(i)
		}
		fmt.Println("return")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			recover()
		}()
		for i := 0; i < 2; i++ {
			if i == 1 {
				panic("стоп")
			}
			fmt.Println(i)
		}
	}()

	wg.Add(1)
	die2 := make(chan bool)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-die2:
				fmt.Println("завершена")
				return
			default:
				fmt.Println(i)
				i++
				if i >= 2 {
					close(die2)
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("таймаут")
				return
			default:
				fmt.Println(i)
				i++
				if i >= 2 {
					cancel()
				}
			}
		}
	}()

	wg.Wait()
}
