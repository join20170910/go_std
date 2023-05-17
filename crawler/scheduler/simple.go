package scheduler

import "awesomeProject/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}
