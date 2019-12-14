package scheduler

import "myproject/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.WorkerChan
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.WorkerChan <- r }()
}
