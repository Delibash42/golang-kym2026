package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan os.Signal, 1)
	die := make(chan string)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-ch:
			println("cntrlc")
			die <- "die"
		}
	}()
	wg.Add(1)
	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-die:
				println("otschet okonchen")
				wg.Done()
				break
			default:
				fmt.Println("time: ", i)
				i++
			}
		}
	}()
	wg.Wait()
}
