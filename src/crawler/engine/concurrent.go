package engine

import (
	"log"
	"myproject/crawler/model"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//ConfigMasterWorkChan(chan Request)
	WorkChan() chan Request
	//WorkReady(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan RequestResult)
	//e.Scheduler.ConfigMasterWorkChan(in)
	e.Scheduler.Run()
	for i := 0; i < e.WorkCount; i++ {
		//createWork(in, out)
		createWork(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}
	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				//log.Printf("Got item is : %v", item)
				log.Printf("Got profile #%d , %v", profileCount, item)
				profileCount++
			}
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//func createWork(in chan Request, out chan RequestResult) {
//func createWork(out chan RequestResult, s Scheduler) {
func createWork(in chan Request, out chan RequestResult, s Scheduler) {

	//in := make(chan Request)
	go func() {
		for {
			s.WorkReady(in)
			request := <-in
			result, err := work(request)
			if err != nil {
				continue
			}
			out <- result
		}

	}()
}
