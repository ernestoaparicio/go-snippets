package concurrent

import (
	"fmt"
	"time"
)

func RunLeakExampleWithFix() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil) // passing a nil chan will cause go routine above to run and wait forever

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}
