package ch3

import "fmt"

// ChanDeadLock causes a deadlock, because
// main goroutine is waiting for a value to be placed
// onto the stringStream, but this will never happen
func ChanDeadLock() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		stringStream <- "Write something"
	}()
	fmt.Println(<-stringStream)
}

// ChanNilReadDeadLock shows that attempt to read
// from nil chan blocks (not necessarily deadlock) a program
func ChanNilReadDeadLock() {
	var dataStream chan interface{}
	<-dataStream
}

// ChanNilWriteDeadLock shows that attempt to write
// to nil chan blocks (not necessarily deadlock) a program
func ChanNilWriteDeadLock() {
	var dataStream chan interface{}
	dataStream <- struct{}{}
}

// ChanNilClosePanic shows that attempt to close
// nil chan causes a panic
func ChanNilClosePanic() {
	var dataStream chan interface{}
	close(dataStream)
}
