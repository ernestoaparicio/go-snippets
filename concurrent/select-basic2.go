package concurrent

import (
	"fmt"
	"time"
)

func RunSelectBasic2() {
	done := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Waiting for goroutine to be done.")
		}
	}
}
