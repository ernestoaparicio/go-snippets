package concurrent

import (
	"fmt"
	"time"
)

func RunDefineOrder() {
	a := func(a, b chan struct{}) {
		<-a
		fmt.Println("A()!")
		time.Sleep(time.Second)
		close(b)
	}

	b := func(a, b chan struct{}) {
		<-a
		fmt.Println("B()!")
		close(b)
	}

	c := func(a chan struct{}) {
		<-a
		fmt.Println("C()!")
	}

	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go c(z)
	go a(x, y)
	go c(z)
	go b(y, z)
	go c(z)

	close(x)
	time.Sleep(3 * time.Second)

}
