/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/15 19:47
* @Description: Fetcher: --input: url; --output: text in utf-8 encoding
***********************************************************************/

package fetcher

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"net/url"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

)

var DefaultHeader = defaultHeader()

func defaultHeader() http.Header {
	h := make(http.Header)
	// 模拟浏览器
	h.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")
	return h
}

func Get(urlS string, header http.Header) ([]byte, error) {
	// 直接调用http.Get容易403 Forbidden
	//resp, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	// 模拟浏览器发起请求
	client := http.DefaultClient
	req, err := http.NewRequest("GET", urlS, nil)
	if err != nil {
		return nil, errors.New("fetcher: error request")
	}
	// 设置请求头，模拟浏览器
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("fetcher: error response")
	}
	defer resp.Body.Close()


	if resp.StatusCode >= 300 && resp.StatusCode <= 500 {
		return nil, fmt.Errorf("fetcher: error status code: %d", resp.StatusCode)
	}

	// 检测html编码方式并转为utf-8编码
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func Post(urlS, body string, header http.Header) ([]byte, error) {

	// 模拟浏览器发起请求
	client := http.DefaultClient
	req, err := http.NewRequest("POST", urlS, strings.NewReader(body))
	if err != nil {
		return nil, errors.New("fetcher: error request")
	}
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("fetcher: error response")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode <= 500 {
		return nil, fmt.Errorf("fetcher: error status code: %d", resp.StatusCode)
	}

	// 检测html编码方式并转为utf-8编码
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func PostForm(urlS string, data url.Values, header http.Header) ([]byte, error) {
	// 注意在header中必须设置"Content-Type" = "application/x-www-form-urlencoded"
	body := data.Encode()	// 编码成这种url参数形式 "bar=baz&foo=quux"
	return Post(urlS, body, header)
}

// 推断html编码方式(依赖前1024字节)
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
