package concurrent

import "fmt"

func RunChannelRange() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Println("%v ", integer)
	}
}
