/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/15 22:17
* @Description: The file is for
***********************************************************************/

package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 有网络连接时可以使用Fetch在线获取网页内容
	//content, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", content)

	// 也可以直接使用本地保存的该网页
	content, err := ioutil.ReadFile("./view-source_https___www.zhenai.com_zhenghun.html")
	if err != nil {
		t.Error(err)
	}
	// fmt.Print(content)

	const resSize = 470
	res := ParseCityList(content)
	if len(res.Requests) != resSize || len(res.Items) != resSize {
		t.Errorf("result should havae %d requests and %d items but got %d requests and %d items\n",
			resSize, resSize, len(res.Requests), len(res.Items))
	}
}
