package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(wg *sync.WaitGroup, id int, ctx context.Context, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		case n, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Worker %d получил %d\n", id, n)
		}
	}
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go worker(&wg, i, ctx, ch)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt, shutting down...")
		cancel()
		close(ch)
	}()

	count := 1
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Все воркеры завершили работу. Выход.")
			return
		case ch <- count:
			count++
			time.Sleep(1 * time.Second)
		}
	}
}
