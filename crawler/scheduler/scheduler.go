package scheduler

import "Cookbook/crawler/engine"

//
//type SimpleSch
//￿eduler struct {
//workerChan chan engine.Request
//}

type SimpleScheduler struct {
	workerChan chan engine.Request
}


func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
