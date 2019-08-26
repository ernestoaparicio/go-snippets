package concurrent

import (
	"fmt"
	"sync"
)

func RunOnceDo() {
	var count int

	fmt.Println("First: Count: %d\n", count)

	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Println("Last: Count: %d\n", count)
}
