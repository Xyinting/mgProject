package parser

import (
	"crawler/engine"
	"regexp"
)

const infoRegx = `<div class="m-btn purple"[^>]*>([^<]*)</div>`

func PersonParser(content []byte) engine.RequestResult {

	regex := regexp.MustCompile(infoRegx)
	matches := regex.FindAllSubmatch(content, -1)

	result := engine.RequestResult{}
	for _, m := range matches {

		result.Items = append(result.Items, m[1])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       "",
				ParseFunc: engine.NilParser,
			})
	}

	return result
}
