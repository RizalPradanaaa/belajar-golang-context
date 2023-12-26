package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			// Akan Berhenti bila ada informasi Cancel
			case <-ctx.Done():
				return
			case destination <- counter:
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Jumlah Goroutine Berjalan ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Jumlah Goroutine Berjalan ", runtime.NumGoroutine())
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("Jumlah Goroutine Berjalan ", runtime.NumGoroutine())
}
