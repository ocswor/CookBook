package engine

import (
	"fmt"
	"log"
)

//演示  一个 channel 循环等待的bug
type ConcurrentEngine_One struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)

}

func (e *ConcurrentEngine_One) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out  //  这里 收 解析的结果
		for _, item := range result.Items {
			fmt.Printf("Got item:%v", item)
		}
		i := 0
		for _, request := range result.Requests {
			log.Printf("submit %d",i)
			e.Scheduler.Submit(request) //  这里将request 送走,但是没有足够的worker 需要等待worker 处理,但是worker
			i++
		}
	}
}


func createWorker(in chan Request, out chan ParseResult) {

	go func() {
		for {
			request := <-in // work 开始读
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result // 发
		}
	}()

}
