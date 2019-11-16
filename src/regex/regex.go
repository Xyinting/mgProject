package main

import (
	"fmt"
	"regexp"
)

const text = `My email is hh@163.com
email is xilovele@qq.com
email2 is ntoap@nantian.com
`

func main() {

	compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)

	match := compile.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}
