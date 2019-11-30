package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkChan(chan Request)
}

func (e ConcurrentEngine) run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan RequestResult)
	e.Scheduler.ConfigMasterWorkChan(in)

	for i := 0; i < e.WorkCount; i++ {
		createWork(in, out)
	}
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item is : %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWork(in chan Request, out chan RequestResult) {
	go func() {
		for {
			request := <-in
			result, err := work(request)
			if err != nil {
				continue
			}
			out <- result
		}

	}()
}
