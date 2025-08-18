package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

var stop int32

func worker1() {
	defer wg.Done()
	for {
		if atomic.LoadInt32(&stop) == 1 {
			fmt.Printf("WORKER1 STOPPED\n")
			return
		}
		fmt.Println("Worker1 working...")
		time.Sleep(500 * time.Millisecond)
	}
}

func worker2(done <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("WORKER2 STOPPED\n")
			return
		default:
			fmt.Println("Worker2 working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker3(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WORKER3 STOPPED\n")
			return
		default:
			fmt.Println("Worker3 working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker4(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WORKER4 STOPPED\n")
			return
		default:
			fmt.Println("Worker4 working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker5(jobs <-chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Worker5 working\n", job)
	}
	fmt.Println("WORKER5 STOPPED\n")
}

func worker6() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("WORKER6 STOPPED\n")
			runtime.Goexit() // мгновенно завершает эту горутину
		}
		fmt.Println("Worker6 working")
		time.Sleep(500 * time.Millisecond)
	}
}

func worker7() {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("WORKER7 recovered:", r)
		}
	}()
	for i := 0; i < 5; i++ {
		if i == 2 {
			panic("force stop goroutine") // завершение горутины
		}
		fmt.Println("Worker7 working:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	done := make(chan struct{})
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	ctx4, cancel4 := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel4()
	defer cancel()

	wg.Add(7)

	go worker1()
	go worker2(done)
	go worker3(ctx)
	go worker4(ctx4)
	go worker5(ch)
	go worker6()
	go worker7()
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)

	time.Sleep(2 * time.Second)

	atomic.AddInt32(&stop, 1)
	close(done)

	cancel()
	time.Sleep(time.Second)

	wg.Wait()
}
