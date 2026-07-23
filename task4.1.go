package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d остановлен\n", id)
			return
		default:
			fmt.Printf("Воркер %d: засек %d\n", id, i)
			i++
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	fmt.Println("Программа запущена")

	<-sigChan
	fmt.Println("\nПолучен сигнал")

	cancel()
	wg.Wait()

}
