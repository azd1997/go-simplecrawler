/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/15 19:54
* @Description: Parse sheet of papers info
***********************************************************************/

package parser

import (
	"log"
	"regexp"

	"github.com/azd1997/go-crawler2/domain/zhenai/engine"
)


const (
	// 检索结果总数
	numRegPattern = `找到$nbsp;([0-9]+)$nbsp;条结果`
	// 检索表格
	sheetRegPattern = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	// 翻页按钮匹配
	nextRegPattern = `<a href="(?curpage=2&amp;RecordsPerPage=20&amp;QueryID=0&amp;ID=&amp;turnpage=1&amp;tpagemode=L&amp;dbPrefix=CJFQ&amp;Fields=&amp;DisplayMode=listmode&amp;PageName=ASP.brief_result_aspx&amp;isinEn=1&amp;)" title="键盘的“← →”可以实现快速翻页">[]</a>`
)


var (
	numRegexp = regexp.MustCompile(numRegPattern)
	sheetRegexp = regexp.MustCompile(sheetRegPattern)
	nextRegexp = regexp.MustCompile(nextRegPattern)
)

func ParseSheet(content []byte) engine.ParseResult {
	matches := sheetRegexp.FindAllSubmatch(content, -1)	// 返回值类型 [][][]byte

	res := engine.ParseResult{}




	for _, m := range matches {
		cityname := string(m[2])
		res.Items = append(res.Items, "City: " + cityname)
		res.Requests = append(res.Requests,
			engine.Request{Url: string(m[1]), ParserFunc:ParseCity}) // 进而解析城市中的用户列表
		log.Printf("City: %s, Url: %s\n", m[2], m[1])

		limit--
		if limit == 0 {
			break
		}
	}
	log.Printf("Matches found: %d\n", len(matches))

	return res
}


//// 使用正则表达式匹配。 （正式做最好用CSS选择器）
//func printCityList(content []byte) {
//	re := regexp.MustCompile(
//		`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
//	matches := re.FindAll(content, -1)
//	for _, v := range matches {
//		fmt.Printf("%s\n", v)
//	}
//	fmt.Printf("Matches found: %d\n", len(matches))
//}
//
//func printCityListSubMatch(content []byte) {
//	re := regexp.MustCompile(
//		`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
//	matches := re.FindAllSubmatch(content, -1)	// 返回值类型 [][][]byte
//	for _, m := range matches {
//		fmt.Printf("City: %s, Url: %s\n", m[2], m[1])
//	}
//	fmt.Printf("Matches found: %d\n", len(matches))
//}