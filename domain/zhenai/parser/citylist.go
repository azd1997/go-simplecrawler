/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/15 19:54
* @Description: Parse citylist of www.zhenai.com
***********************************************************************/

package parser

import (
	"github.com/azd1997/go-crawler2/domain/zhenai/engine"
	"log"
	"regexp"
)

const cityListRegPattern = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

const limitCity = 5	// 限制只爬取10个城市，避免在开发过程中爬取过多页面，影响调试

// 这是为了避免正则表达式的重复编译（相当于字符串匹配中的模式串预处理）
var cityListRegexp = regexp.MustCompile(cityListRegPattern)

func ParseCityList(content []byte) engine.ParseResult {
	matches := cityListRegexp.FindAllSubmatch(content, -1)	// 返回值类型 [][][]byte

	res := engine.ParseResult{}
	limit := limitCity
	for _, m := range matches {
		cityname := string(m[2])
		res.Items = append(res.Items, "City: " + cityname)
		res.Requests = append(res.Requests,
			engine.Request{Url: string(m[1]), ParserFunc: ParseCity}) // 进而解析城市中的用户列表
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