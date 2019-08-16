package functional

import (
	"fmt"
	"time"
)

func namedGreeting(name string) {
	fmt.Printf("Hey %s!\n", name)
}

func anonymousGreeting() func(string) {
	return func(name string) {
		fmt.Printf("Hey %s!\n", name)
	}
}

func RunFunctionalAnonymous() {
	namedGreeting("Alice")

	greet := anonymousGreeting()
	greet("Bob")

	func(name string) {
		fmt.Printf("Hello %s!\n", name)
	}("Cindy")
}

func GreetingClosure(name string) func() {
	msg := name + fmt.Sprintf(" (at %v)", time.Now().String())

	closure := func() {
		fmt.Printf("Hey %s!\n", msg)
	}

	return closure
}
