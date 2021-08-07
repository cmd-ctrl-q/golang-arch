package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Printf("GOROUTINES RUNNING = %d\n", runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	// launch 100 go routines
	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println("Launching goroutine:", n)
			// infinite for loop
			for {

				// do work
				time.Sleep(50 * time.Millisecond)

				// go routine either exits or does work
				select {
				case <-ctx.Done():
					runtime.Goexit() // exit go routine, similar to return
				default:
					fmt.Printf("goroutine %d doing work\n", n)
					time.Sleep(50 * time.Millisecond)
				}
			}
		}(i)
	}

	// goroutines finish in < 1 millisecond
	time.Sleep(time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING AFTER ONE MILLISECOND = %d\n", runtime.NumGoroutine())

	cancel()
	// give goroutines time to exit
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING AFTER CANCEL() = %d\n", runtime.NumGoroutine())
}
