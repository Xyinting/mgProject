package main

import (
	"fmt"
	"myproject/crawler/engine"
	"myproject/crawler/zhenghun/parser"
	"regexp"
)

func main() {

	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.CityListParser,
	})

	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", resp.StatusCode)
	//	return
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//processCity(body)

}

func processCity(bytes []byte) {

	match := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

	matches := match.FindAllSubmatch(bytes, -1)

	for _, m := range matches {
		//for _, submatch := range m {
		//	fmt.Printf("%s ", submatch)
		//}

		fmt.Printf("City: %s, URL: %s", m[2], m[1])
		fmt.Println()
	}

	fmt.Printf("Matches city number: %d \n ", len(matches))
}
