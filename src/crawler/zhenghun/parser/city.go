package parser

import (
	"myproject/crawler/engine"
	"regexp"
)

const personReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func CityParser(content []byte) engine.RequestResult {

	regex := regexp.MustCompile(personReg)

	matchs := regex.FindAllSubmatch(content, -1)

	result := engine.RequestResult{}
	for _, m := range matchs {

		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: PersonParser,
		})

		break
	}

	return result

}
