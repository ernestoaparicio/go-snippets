package concurrent

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func RunSimpleContext() {
	f1 := func(t int) {
		c1 := context.Background()
		c1, cancel := context.WithCancel(c1)
		defer cancel()

		go func() {
			time.Sleep(4 * time.Second)
			cancel()
		}()

		select {
		case <-c1.Done():
			fmt.Printf("f1(): %v\n", c1.Err())
			return
		case r := <-time.After(time.Duration(t) * time.Second):
			fmt.Printf("f1(): %v\n", r)
			return
		}

	}
	f2 := func(t int) {
		c2 := context.Background()
		c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
		defer cancel()

		go func() {
			time.Sleep(4 * time.Second)
			cancel()
		}()

		select {
		case <-c2.Done():
			fmt.Println("f2():", c2.Err())
			return
		case r := <-time.After(time.Duration(t) * time.Second):
			fmt.Println("f2():", r)
		}
		return
	}

	f3 := func(t int) {
		c3 := context.Background()
		deadline := time.Now().Add(time.Duration(2*t) * time.Second)
		c3, cancel := context.WithDeadline(c3, deadline)
		defer cancel()

		go func() {
			time.Sleep(4 * time.Second)
			cancel()
		}()

		select {
		case <-c3.Done():
			fmt.Println("f3():", c3.Err())
			return
		case r := <-time.After(time.Duration(t) * time.Second):
			fmt.Println("f3():", r)
		}
		return
	}

	if len(os.Args) != 2 {
		fmt.Printf("Need a delay!\n")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("Delay: %v\n", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}
