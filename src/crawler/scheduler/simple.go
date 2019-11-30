package scheduler

import "myproject/crawler/engine"

type simpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *simpleScheduler) ConfigMasterWorkChan(c chan engine.Request) {
	s.WorkChan = c
}

func (s *simpleScheduler) Submit(r engine.Request) {
	s.WorkChan <- r
}
