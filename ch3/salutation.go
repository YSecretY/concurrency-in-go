package ch3

import (
	"fmt"
	"sync"
)

// SalutationWrong will print good day 3 times
// because salutation has been assigned to the next string value earlier
// than goroutine started printing it
func SalutationWrong() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

// Salutation function created in the right way
// every word was taken and passed to the goroutine
func Salutation() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
