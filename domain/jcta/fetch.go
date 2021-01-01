/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/23 8:59
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
	"github.com/azd1997/go-crawler2/fetcher"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func fetch(url string) (string, error) {
	// 实测使用检测网页编码方式，会检测错误。具体见该网页源代码，其中标记了两种编码格式
	//
	//content, err := fetcher.Get(indexPage, fetcher.DefaultHeader)
	//if err != nil {
	//	return nil, err
	//}

	// 模拟浏览器发起请求
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.New("fetcher: error request")
	}
	// 设置请求头，模拟浏览器
	req.Header = fetcher.DefaultHeader

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("fetcher: error response")
	}
	defer resp.Body.Close()


	if resp.StatusCode >= 300 && resp.StatusCode <= 500 {
		return "", fmt.Errorf("fetcher: error status code: %d", resp.StatusCode)
	}

	// 读取响应内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
