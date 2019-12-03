package parser

import (
	"myproject/crawler/engine"
	"myproject/crawler/model"
	"regexp"
)

const infoRegx = `<div class="m-btn purple"[^>]*>([^<]*)</div>`

func PersonParser(content []byte, name string) engine.RequestResult {

	profile := model.Profile{}
	profile.Name = name
	var info []string

	regex := regexp.MustCompile(infoRegx)
	matches := regex.FindAllSubmatch(content, -1)

	for _, m := range matches {
		info = append(info, string(m[1]))
	}
	profile.PersonalInfo = info

	result := engine.RequestResult{
		Items: []interface{}{profile},
	}

	return result
}
