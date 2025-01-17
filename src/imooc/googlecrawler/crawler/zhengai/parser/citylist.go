package parser

import (
	"googlecrawler/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})

		// fmt.Printf("City：%s,URL：%s\n", m[2], m[1])
		// for _, subMatch := range m {
		// 	fmt.Printf("%s ", subMatch)
		// }
	}
	// fmt.Printf("%d\n", len(matches))
	return result
}
