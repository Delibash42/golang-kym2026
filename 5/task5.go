package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	die := make(chan bool)
	go func() {
		i := 1
		defer close(ch)
		for {
			select {
			case <-die:
				return
			default:
				ch <- i
				i++
				time.Sleep(1 * (time.Second))
			}

		}
	}()
	go func() {
		for {
			select {
			case <-die:
				return
			case i := <-ch:
				fmt.Println("засек", i)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	die <- true
}
