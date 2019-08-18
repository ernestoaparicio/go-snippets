package concurrent

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunParallelLoop() {
	urls := []string{"http://www.google.com", "http://www.optum.com", "http://www.espn.com", "http://www.golang.org", "http://www.asu.edu", "http://www.cnn.com", "http://www.fox.com"}
	ch := make(chan struct{})

	for _, u := range urls {
		go func(u string) { // pass u explictly to ensure correct u is evaluated
			res, err := http.Get(u)
			if err != nil {
				log.Fatal("Err %v", err)
			}
			time.Sleep(2 * time.Second)
			fmt.Println(u + " : " + res.Status)
			ch <- struct{}{}
		}(u)
	}

	// wait for goroutines to complete
	for range urls {
		<-ch
	}
}
