/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/15 20:04
* @Description: The file is for
***********************************************************************/

package engine

import (
	"github.com/azd1997/go-crawler2/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request	// 请求队列

	for _, req := range seeds {
		requests = append(requests, req)
	}

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s\n", req.Url)

		body, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Printf("Fetcher: error " + "fetching url %s: %v", req.Url, err)
			continue
		}

		parseResult := req.ParserFunc(body)		// 调用每个请求自定义的解析函数
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v\n", item)
		}

	}
}
