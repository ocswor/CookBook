package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		//total.Lock()
		total.value += 1
		//total.Unlock()
	}
}

func main() {
	fmt.Println(total.value)
	var wg sync.WaitGroup
	wg.Add(100)
	for i:=1;i<=100 ;i++  {
		go worker(&wg)
	}
	wg.Wait()

	fmt.Println(total.value)
}

