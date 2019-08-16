package concurrent

import "fmt"

func mirroredQuery() {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()

	for x := range <-responses {
		fmt.Println(x)
	}
}

func request(hostname string) string {
	return hostname
}

func RunBufferedChannel() {
	mirroredQuery()
}
