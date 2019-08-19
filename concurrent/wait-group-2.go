package concurrent

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func RunWaitGroup2() {
	urls := []string{"http://www.google.com", "http://www.optum.com", "http://www.espn.com", "http://www.golang.org", "http://www.asu.edu", "http://www.cnn.com", "http://www.fox.com"}
	ch := make(chan struct{})
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go func(u string) { // pass u explictly to ensure correct u is evaluated
			defer wg.Done()
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
	wg.Wait()
	close(ch)
}
