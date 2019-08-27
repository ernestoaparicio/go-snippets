package concurrent

import "fmt"

func RunLeakExample() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings {
				// do something interesting
				fmt.Println(s)
			}
		}()

		return completed
	}

	doWork(nil) // passing a nil chan will cause go routine above to run and wait forever
	// perhaps more work is done here
	fmt.Println("Done.")
}
