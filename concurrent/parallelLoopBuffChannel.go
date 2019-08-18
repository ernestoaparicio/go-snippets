package concurrent

import (
	"fmt"
	"net/http"
	"time"
)

func RunParallelLoopBuffChannel() {
	type responseObj struct {
		res *http.Response
		url string
		err error
	}

	urls := []string{"http://www.google.com", "http://www.optumBADURL.com", "http://www.espn.com", "http://www.golang.org", "http://www.asu.edu", "http://www.cnn.com", "http://www.fox.com"}
	ch := make(chan responseObj, len(urls))

	for _, u := range urls {
		go func(u string) { // pass u explictly to ensure correct u is evaluated
			var resObj responseObj
			resObj.url = u
			resObj.res, resObj.err = http.Get(u)
			time.Sleep(2 * time.Second)
			ch <- resObj
		}(u)
	}

	for range urls {
		resObj := <-ch

		if resObj.err != nil {
			fmt.Println(resObj.url, resObj.err.Error())
		} else {
			fmt.Println(resObj.url, resObj.res.Status)
		}
	}
}
