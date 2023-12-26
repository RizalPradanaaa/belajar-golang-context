package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterTimeOut(ctx context.Context) chan int {
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
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Jumlah Goroutine Berjalan ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterTimeOut(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Jumlah Goroutine Berjalan ", runtime.NumGoroutine())
}
