package main

/*
ab -n 2000 -c 200 http://127.0.0.1:89/search/

Document Path:          /search/
Document Length:        3 bytes

Concurrency Level:      100
Time taken for tests:   0.143 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      161000 bytes
HTML transferred:       3000 bytes
Requests per second:    7000.50 [#/sec] (mean)
Time per request:       14.285 [ms] (mean)
Time per request:       0.143 [ms] (mean, across all concurrent requests)
Transfer rate:          1100.66 [Kbytes/sec] received


Server Software:
Server Hostname:        127.0.0.1
Server Port:            89

Document Path:          /search/
Document Length:        3 bytes

Concurrency Level:      200
Time taken for tests:   0.588 seconds
Complete requests:      2000
Failed requests:        0
Non-2xx responses:      2000
Total transferred:      322000 bytes
HTML transferred:       6000 bytes
Requests per second:    3401.88 [#/sec] (mean)
Time per request:       58.791 [ms] (mean)
Time per request:       0.294 [ms] (mean, across all concurrent requests)
Transfer rate:          534.87 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   18  16.4     10      70
Processing:     2   39  23.9     32      99
Waiting:        0   30  19.0     23      94
Total:          7   57  30.4     63     129

*/
import (
	"Cookbook/learnhttp/learn_v2/controller"
	"log"
	"net/http"
)

func main() {

	http.Handle("/search/",controller.App_User{})

	port := ":89"
	log.Printf("linsten:%s",port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic(err)
	}
}
