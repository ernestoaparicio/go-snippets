package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

func RunNilChannel() {

	add := func(c chan int) {
		sum := 0
		t := time.NewTimer(time.Second)

		for {
			select {
			case input := <-c:
				sum = sum + input
			case <-t.C:
				c = nil
				fmt.Println(sum)
			}
		}
	}

	send := func(c chan int) {
		for {
			c <- rand.Intn(10)
		}
	}

	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(3 * time.Second)

}
