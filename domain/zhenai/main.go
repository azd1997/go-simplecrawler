/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/15 13:56
* @Description: The file is for
***********************************************************************/

package main

import (
	"github.com/azd1997/go-crawler2/domain/zhenai/engine"
	"github.com/azd1997/go-crawler2/domain/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}

