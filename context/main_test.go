package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	// create a channel
	ch := make(chan int)

	go func() {
		// close the channel when the function is done
		defer close(ch)
		counter := 1

		for {
			// select is like switch but for channel
			select {
			// if the context is done
			case <-ctx.Done():
				return
			default:
				// send the counter to the channel
				ch <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return ch
}

func TestContextWithCancel(t *testing.T) {
	// call context background
	parentCtx := context.Background()
	// cal the context with cancel
	ctx, cancel := context.WithCancel(parentCtx)

	// print the number of goroutine
	fmt.Println("Total Goroutine first", runtime.NumGoroutine())

	// parsing the context to the function
	dst := CreateCounter(ctx)

	for n := range dst {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	// cancel the context
	cancel()

	time.Sleep(2 * time.Second)

	// the number of goroutine is still the same
	fmt.Println("Total Goroutine last", runtime.NumGoroutine())
}