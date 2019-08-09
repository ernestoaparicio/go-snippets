package concurrent

import (
	"fmt"
	"time"
)

func RunBasic() {
	nameChannel := make(chan string)
	ageChannel := make(chan string)

	go func() {
		names := []string{"John", "Paul", "George", "Ringo"}

		for _, name := range names {
			time.Sleep(0 * time.Second)
			fmt.Println(name)
		}
		nameChannel <- ""
	}()

	go func() {
		ages := []int{23, 67, 11, 89}

		for _, age := range ages {
			time.Sleep(0 * time.Second)
			fmt.Println(age)
		}

		ageChannel <- ""
	}()

	<-nameChannel
	<-ageChannel

	//time.Sleep(10 * time.Second)

}
