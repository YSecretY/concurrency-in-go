package ch4

import (
	"bytes"
	"fmt"
	"sync"
)

// Confinement is an idea of ensuring information is only
// ever available from one concurrent process
func Confinement() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

// LexicalConfinement involves using lexical scope to expose only the correct data and
// concurrency primitives for multiple concurrent processes to use.
// It makes it impossible to do the wrong thing.
func LexicalConfinement() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving")
	}

	results := chanOwner()
	consumer(results)
}

// BufConfinement is an example of using confinement with buffer
// in this case it is impossible to do something wrong because
// printData takes a different subsets of the slice thus
// goroutines starts to only the part is passed in
func BufConfinement() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
