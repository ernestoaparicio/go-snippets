package concurrent

import (
	"fmt"
	"sync"
)

func RunWaitForFinish() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello world")
			wg.Done()
		}()
	}

	wg.Wait()
	//time.Sleep(5 * time.Second)
}
