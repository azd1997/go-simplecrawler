/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/16 9:24
* @Description: Parse city, list users.
***********************************************************************/

package parser

import (
	"github.com/azd1997/go-crawler2/domain/zhenai/engine"
	"log"
	"regexp"
)

// 用户ID为数字串`[0-9]+`; `[^>]*`表示匹配任意字符直至遇到第一个`>`; `[^<]+`匹配<a>xxx</a>中的xxx，这是用户名
const cityRegPattern = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var cityRegexp = regexp.MustCompile(cityRegPattern)

// 解析城市，获取城市下面的所有用户的列表
func ParseCity(content []byte) engine.ParseResult {
	matches := cityRegexp.FindAllSubmatch(content, -1)	// 返回值类型 [][][]byte

	res := engine.ParseResult{}
	for _, m := range matches {
		username := string(m[2])
		res.Items = append(res.Items, "User: " + username)
		res.Requests = append(res.Requests,
			engine.Request{
				Url:string(m[1]),
				// 闭包传递name参数
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, username)
			}})
		log.Printf("User: %s, Url: %s\n", m[2], m[1])
	}
	log.Printf("Matches found: %d\n", len(matches))

	return res
}
