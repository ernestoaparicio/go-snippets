package concurrent

import (
	"fmt"
	"time"
)

func RunDefineOrder2() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	one := func() {
		val := 1
		time.Sleep(2 * time.Second)
		fmt.Println("one =", val)
		a <- val
	}

	two := func() {
		time.Sleep(2 * time.Second)
		val := <-a + 1
		fmt.Println("two =", val)
		b <- val
	}

	three := func() {
		time.Sleep(2 * time.Second)
		val := <-b + 1
		fmt.Println("three =", val)
		c <- val
	}

	go three()
	go one()
	go two()

	<-c
	fmt.Println("done")
}
