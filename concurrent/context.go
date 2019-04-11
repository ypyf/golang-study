package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func worker(ctx context.Context, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d exit.\n", id)
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("worker %d is running...\n", id)
		}
	}
}

func main() {
	wg.Add(3)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	go worker(ctx, 1)
	go worker(ctx, 2)
	go worker(ctx, 3)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("All tasks are finished.")
}
