package ch4

import (
	"fmt"
	"math/rand"
	"time"
)

// Clean shows how to close goroutine by another one
func Clean() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWord exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		// Will cause select case <-done
		close(done)
	}()

	// Join goroutines
	<-terminated
	fmt.Println("Done")
}

// CleanWriteChan shows how to clean up closure goroutine with
func CleanWriteChan() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				// Will close closure func
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	// Will cause select case <-done
	close(done)

	// Simulate work
	time.Sleep(1 * time.Second)
}

// Good convention
// If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine
