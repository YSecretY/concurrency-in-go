package ch3

import "sync"

// OnceDeadlock shows an example of deadlock caused by
// calling sync.Once several times that depends on each other
func OnceDeadlock() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
}
