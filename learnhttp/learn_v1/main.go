package main

/*
Concurrency Level:      100
Time taken for tests:   0.082 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      207000 bytes
HTML transferred:       19000 bytes
Requests per second:    12211.35 [#/sec] (mean)
Time per request:       8.189 [ms] (mean)
Time per request:       0.082 [ms] (mean, across all concurrent requests)
Transfer rate:          2468.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   1.2      3       7
Processing:     1    4   1.4      4       9
Waiting:        0    3   1.3      3       7
Total:          4    8   1.5      7      12


*/

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/search/", CreaterUser)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := middleWareHandler{r: r}
	if err := http.ListenAndServe(":89", mh); err != nil {
		panic(err)
	}

}
