package engine

import (
	"log"
	"myproject/crawler/fetcher"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...Request) {

	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {

		r := requests[0]
		requests = requests[1:]

		parseResult, err := work(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, items := range parseResult.Items {
			log.Printf("Got item %v", items)
		}
	}
}

func work(r Request) (RequestResult, error) {

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s : %v", r.Url, err)
		return RequestResult{}, err
	}

	return r.ParseFunc(body), nil
}
