package ch3

import (
	"fmt"
	"sync"
	"time"
)

// WG is an example how to use sync.WaitGroup
func WG() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine is running...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine is running...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")
}
