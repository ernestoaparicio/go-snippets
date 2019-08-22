package concurrent

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func sayHello() {
	time.Sleep(10 * time.Second)
	fmt.Println("Done.")
	ch <- 1
}

func RunBasic2() {
	go sayHello()

	<-ch
}
