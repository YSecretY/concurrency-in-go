package ch3

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

// CloseChan shows it is possible to read from
// close channels. In addition, operator <-
// can return additional bool value
func CloseChan() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Printf("(%v): %v", ok, integer)
}

// RangeChan shows it is possible
// to iterate using range through a chan
func RangeChan() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}

	fmt.Printf("\n")
}

// SignalCloseChan shows if n goroutines waiting
// on one chan, it is possible to unblock them
// by closing this chan
func SignalCloseChan() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

// BuffChan is an example of little
// optimization using buffered channels
func BuffChan() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

// ChanExample gives an example of good practice using channels
// 1. Instantiate a channel
// 2. Perform writes, or pass ownership to another goroutine
// 3. Close channel
// 4. Encapsulate the previous three things in this list and expose them via a reader channel.
func ChanExample() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
