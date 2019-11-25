package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const text = `My email is hh@163.com
email is xilovele@qq.com
email2 is ntoap@nantian.com
`

func main() {

	//compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	//
	//match := compile.FindAllStringSubmatch(text, -1)
	////fmt.Println(match)
	//for _, m := range match {
	//	fmt.Println(m)
	//}

	resp, err := http.Get("http://album.zhenai.com/u/1202417524")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s", body)

}
