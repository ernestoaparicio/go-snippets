package concurrent

import "fmt"

func RunBetweenTwo() {
	nameChannel := make(chan string)
	doneChannel := make(chan string)

	go func() {
		names := []string{"John", "Paul", "George", "Ringo"}

		for _, name := range names {
			fmt.Println("Processing first stage of: " + name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	go func() {
		for name := range nameChannel {
			fmt.Println("Processing the second stage of: " + name)
		}
		doneChannel <- ""
	}()

	<-doneChannel
}
