package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}
func generator() chan int {
	out := make(chan int)

	go func() {
		i := 0
		//fmt.Println(time.Duration(rand.Intn(1500)) *time.Millisecond)
		for {
			//fmt.Println(time.Duration(rand.Intn(1500)) *time.Millisecond)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i // 异步会阻塞 //同步会报错
			i++
		}
	}()
	return out
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator() // 两个生产者
	var worker = createWorker(0) //  一个消费者

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select { // select 机制 可以监听多个chanel状态,可以通选接受其他通道的消息
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]

		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println(
				"queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
