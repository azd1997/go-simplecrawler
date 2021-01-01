/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/16 16:38
* @Description: The file is for
***********************************************************************/

package engine

import (
	"encoding/json"
	"fmt"
	"github.com/azd1997/go-crawler2/domain/cnki/cli"
	"github.com/azd1997/go-crawler2/domain/cnki/config"
	"github.com/azd1997/go-crawler2/fetcher"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// 获取cookie
	BASIC_URL = "http://kns.cnki.net/kns/brief/result.aspx"
	// 利用post请求先行注册一次
	SEARCH_HANDLE_URL = "http://kns.cnki.net/kns/request/SearchHandler.ashx"
	// 发送get请求获得文献资源
	GET_PAGE_URL = "http://kns.cnki.net/kns/brief/brief.aspx?pagename="
	// 下载的基础链接
	DOWNLOAD_URL = "http://kns.cnki.net/kns/"
	// 切换页面基础链接
	CHANGE_PAGE_URL = "http://kns.cnki.net/kns/brief/brief.aspx"

	// 每页展示20条结果
	resultNumPerPage = 20
)

// 配置参数
var (
	conf    = config.ParseConfig()
	dir     = "./data_" + strconv.Itoa(int(time.Now().Unix()))
	curPage = 1
)

// POST请求提交的数据，这些是静态的，加上从用户输入的部分，组装成完整的提交数据
var postData = url.Values {
	"action":      []string{""},
	"NaviCode":    []string{"*"},
	"ua":          []string{"1.21"},
	"isinEn":      []string{"1"},
	"PageName":    []string{"ASP.brief_default_result_aspx"},
	"DbPrefix":    []string{"SCDB"},
	"DbCatalog":   []string{"中国学术期刊网络出版总库"},
	"ConfigFile":  []string{"CJFQ.xml"},
	"db_opt":      []string{"CJFQ,CDFD,CMFD,CPFD,IPFD,CCND,CCJD"}, // 搜索类别（CNKI右侧的）
	"db_value":    []string{"中国学术期刊网络出版总库"},
	"year_type":   []string{"echar"},
	"his":         []string{"0"},
	"db_cjfqview": []string{"中国学术期刊网络出版总库,WWJD"},
	"db_cflqview": []string{"中国学术期刊网络出版总库"},
	"__":          []string{now()},
}

func Run(seed ...Request) {

	var requests []Request // 请求队列
	requests = append(requests, seed...)

	h := getHeader()       // 构造请求头
	fmt.Println(postData.Encode())     // 调试

	// 第一次发送POST请求，需要携带数据
	firstPostResult := firstPost(SEARCH_HANDLE_URL, postData, h)

	// 第二次则直接GET，但是需要附上第一个检索条件的值
	keyV := strconv.Quote(postData.Get("text_1_value1"))
	getResultUrl := GET_PAGE_URL + string(firstPostResult) +
		"&t=1544249384932&keyValue=" + keyV + "&S=1&sorttype="

	// 检索结果的第一个页面
	secondGetResult, err := fetcher.Get(getResultUrl, h)
	if err != nil {
		log.Printf("second get failed: %v\n", err)
	}
	fmt.Println("secondGetResult: ", string(secondGetResult)) // 调试

	// 下一页按钮的匹配
	changePageRegexp := regexp.MustCompile(
		`.*?pagerTitleCell.*?<a href="(.*?)".*`)
	match := changePageRegexp.FindSubmatch(secondGetResult)
	if len(match) < 2 {		// 说明没有匹配到
		return
	}

	changePageUrl := match[1]
	fmt.Println(changePageUrl)

}

// 检查总结果数，并根据用户输入决定应该解析多少页
// 注意这里的序号都是从1开始，0代表没有
func preParsePage(pageSource []byte) (pageStart, pageEnd int) {
	// 获取检索到的总结果数
	resultNumRegPattern := `.*?找到&nbsp;(.*?)&nbsp;`	// 网页上检索到的结果总数
	resultNumRegexp := regexp.MustCompile(resultNumRegPattern)
	match := resultNumRegexp.FindSubmatch(pageSource)
	if len(match) < 2 {		// 匹配失败
		return 0, 0
	}
	resultNumStr := strings.ReplaceAll(string(match[1]), ",", "")
	resultNum, err := strconv.Atoi(resultNumStr)
	if err != nil {
		return 0, 0
	}
	needTime := time.Duration(resultNum) * time.Duration(conf.StepWaitSecond) * time.Second
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("检索到%d条结果，全部下载约需要%s。\n", resultNum,
		needTime.String())

	// 让用户决定是否全部下载
	fmt.Println("-------------------------------------------------------")
	downloadAll := ""
	fmt.Print("> 是否全部下载（y/n）：")
	_, err = fmt.Scan(&downloadAll)
	if err != nil {
		log.Fatalf("scan download_all failed: %v\n", err)
	}
	switch downloadAll {
	case "y", "Y":
		pageEnd = resultNum / resultNumPerPage
		if resultNum % resultNumPerPage != 0 {
			pageEnd++
		}
		return 1, pageEnd
	case "n", "N":
		// 不下载全部的话，1.就指定范围，指定该范围后，程序根据范围计算整张的页数
		// 2.指定数量，前n条结果
		arg := ""
		fmt.Println("-------------------------------------------------------")
		fmt.Print("> 请指定下载前n条记录或者m-n范围（例如33 33-56）：")
		_, err = fmt.Scan(&arg)
		if err != nil {
			log.Fatalf("scan result_arg failed: %v\n", err)
		}
		// 按"-"切分
		args := strings.Split(arg, "-")
		switch len(args) {
		case 2:		// 这是想要的
			left, _ := strconv.Atoi(args[0])
			right, _ := strconv.Atoi(args[1])
			if left <= 0 || left > resultNum || right <= 0 || right > resultNum || left > right {
				log.Fatalln("incorrect result_arg!")
			}
			// 否则计算left处于第几页，right处于第几页
			if left % resultNumPerPage == 0 {
				pageStart = left / resultNumPerPage
			} else {
				pageStart = left / resultNumPerPage + 1
			}
			if right % resultNumPerPage == 0 {
				pageEnd = right / resultNumPerPage
			} else {
				pageEnd = right / resultNumPerPage + 1
			}
			return pageStart, pageEnd
		case 1:		// 说明没有破折号，尝试按规则2解析
			right, _ := strconv.Atoi(args[0])
			if right <= 0 || right > resultNum {
				log.Fatalln("incorrect result_arg!")
			}
			// 否则计算，right处于第几页
			if right % resultNumPerPage == 0 {
				pageEnd = right / resultNumPerPage
			} else {
				pageEnd = right / resultNumPerPage + 1
			}
			return 1, pageEnd

		default:
			log.Fatalln("incorrect result_arg!")
		}

	default:
		log.Fatalln("incorrect download_all!")
	}
	return 0, 0
}

// 保存页面信息，解析每一页的下载地址
func parsePage(downloadPageLeft, pageSource) {

}

// 第一次POST请求进行一次临时注册
func firstPost(url string, postData url.Values, h http.Header) []byte {
	firstPostResult, err := fetcher.PostForm(SEARCH_HANDLE_URL, postData, h)
	if err != nil {
		log.Printf("first post failed: %v\n", err)
	}
	fmt.Println("firstPostResult: ", string(firstPostResult)) // 调试
	return firstPostResult
}

func getHeader() http.Header {
	h := fetcher.DefaultHeader
	h.Set("Host", "kns.cnki.net")
	h.Set("Connection", "keep-alive")
	h.Set("Cache-Control", "max-age=0")
	h.Set("Referer", "https://kns.cnki.net/kns/brief/result.aspx?dbprefix=CJFQ")
	// cnki网站提交表单使用的是x-www-form-urlencoded。通过fiddler抓包可以看到
	h.Set("Content-Type", "application/x-www-form-urlencoded")
	return h
}

func getPostData() string {
	// 获取完整提交数据的map
	cli.GetUserInput(postData)
	// 转成json字符串
	res, err := json.Marshal(postData)
	if err != nil {
		return "{}" // 这会导致请求失败
	}
	return string(res)
}

//////////////////////////////////////////////////////////

// 准备好存放数据的目录
func prepareDataDir() {
	// 检测dir是否存在
	if info, err := os.Stat(dir); err != nil {
		if err != os.ErrNotExist {
			log.Fatalf("get %s info failed: %v\n", dir, err)
		} else { // 目录不存在，则创建
			os.MkdirAll(dir, os.ModePerm)
		}
	} else {
		if info.IsDir() { // 目录已存在则递归删除
			os.RemoveAll(dir)
		}
	}
}

/////////////////////////////////////////////////////////////

const timeFormat = "%s %s %d %d:%d:%d %d GMT+0800 (中国标准时间)"

func now() string {
	now := time.Now().Local()
	nowStr := fmt.Sprintf(timeFormat,
		now.Weekday().String()[:3],
		now.Month().String()[:3],
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Year(),
	)
	return nowStr
}
