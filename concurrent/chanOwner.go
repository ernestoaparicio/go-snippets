package concurrent

import (
	"fmt"
	"time"
)

func getResults() <-chan int {
	results := make(chan int)

	go func() {
		defer close(results)
		for i := 0; i < 5; i++ {
			time.Sleep(5 * time.Second)
			results <- i
		}
	}()

	return results
}

func RunChanOwner() {
	results := getResults() // func returns a read only channel, main routine waits while ranging through results

	for val := range results {
		fmt.Println("Val", val)
	}
}
