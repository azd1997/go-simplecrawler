/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/15 19:54
* @Description: The file is for
***********************************************************************/

package engine

type Request struct {
	Url string
	ParserFunc func ([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

// 零解析器，主要用于占位，在开发阶段使程序编译通过
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
