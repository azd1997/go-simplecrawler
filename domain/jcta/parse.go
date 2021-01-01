/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/23 9:02
* @Description: The file is for
***********************************************************************/

package main

import "fmt"

// 第一次爬取列表页，首页.
// 返回值为文章超链接列表和错误
func parseIndexPage() ([]string, error) {

	var detailPages []string
	var err error

	// 抓取响应内容
	content, err := fetch(indexPage)
	if err != nil {
		return nil, err
	}

	// 正则表达式匹配所有文章
	matched := papersRegexp.FindAllStringSubmatch(content, -1)
	for i, v := range matched {
		fmt.Println(i, v)
		detailPages = append(detailPages, v[1])
	}

	return detailPages, nil
}



func parseDetailPage(urlSuffix string) (*Paper, error) {

	url := detailPagePrefix + urlSuffix

	// 抓取响应内容
	content, err := fetch(url)
	if err != nil {
		return nil, err
	}

	p := &Paper{}
	p.AbstractUrl = url
	p.setTitle(content)
	p.setAuthors(content)
	p.setAbstract(content)
	p.setKeywords(content)
	p.setPDFLink(content)
	p.setPosition(content)

	return p, nil
}
