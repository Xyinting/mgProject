package scheduler

import "myproject/crawler/engine"

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkChan(c chan engine.Request) {
	s.WorkChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.WorkChan <- r }()
}
