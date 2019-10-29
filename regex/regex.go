package main

import (
	"fmt"
	"regexp"
)

const text = "My email is hh@163.com"

func main() {

	compile := regexp.MustCompile(`.+@.+\..+`)

	match := compile.FindString(text)
	fmt.Println(match)

}
