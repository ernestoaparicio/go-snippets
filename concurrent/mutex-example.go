package concurrent

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

var increment = func() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Printf("Incrementing: %d\n", count)
}

var decrement = func() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Printf("Decrementing: %d\n", count)
}

func RunMutexExample() {
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete")
}
