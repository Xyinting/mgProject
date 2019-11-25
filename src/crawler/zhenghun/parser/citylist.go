package parser

import (
	"crawler/engine"
	"regexp"
)

const regexstr = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func CityListParser(content []byte) engine.RequestResult {

	match := regexp.MustCompile(regexstr)

	matches := match.FindAllSubmatch(content, -1)

	result := engine.RequestResult{}
	for _, m := range matches {

		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(m[1]),
				ParseFunc: CityParser,
			})

		break
	}

	return result
}
