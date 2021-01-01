/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/20 23:13
* @Description: crawl http://jcta.alljournals.ac.cn/cta_cn/ch/index.aspx
***********************************************************************/

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

// jcta网站上的论文爬取步骤，首先访问论文列表页，获取本期文章表格
// 通过正则表达式获取所有文章的超链接
// 紧接着根据超链接去爬取所有文章详情，写入到papers.md文件（主要用于我快速编辑推文）




const (
	indexPage = "http://jcta.alljournals.ac.cn/cta_cn/ch/index.aspx?year_id=2020&quarter_id=11&Submit=提交"
	detailPagePrefix = "http://jcta.alljournals.ac.cn/cta_cn/ch/"
	pdfPagePrefix = "http://jcta.alljournals.ac.cn/cta_cn/ch/reader/"
	journal = "控制理论与应用"

	// <a href="reader/view_abstract.aspx?file_no=CCTA200068&flag=1">局部新冠肺炎时滞模型及再生数的计算</a>
	// reader/view_abstract.aspx?file_no=CCTA200068&flag=1
	// 局部新冠肺炎时滞模型及再生数的计算
	papersRegPattern = `<a href="(reader/view_abstract.aspx\?file_no=CCTA\d{6}&flag=1)">([^<]+)</a>`

	// .?是特殊字符
	titleRegPattern = `<meta name="DC\.Title" content="([^{/>}]+)"/>`
	authorRegPattern = `<meta name="DC\.Contributor" content="([^"]+)"/>`
	abstractRegPattern = `<span id="Abstract">([^<]+)</span>`
	keywordsRegPattern = `<meta name="keywords" content="([^"]+)"/>`
	positionRegPattern = `<span id="Position">([^<]+)</span>`
	pdflinkRegPattern = `<a href='(create_pdf\.aspx\?file_no=CCTA\d{6}&flag=1&journal_id=cta_cn&year_id=\d{4})' target='_blank'><u>查看全文</u></a>`
)

var (
	papersRegexp = regexp.MustCompile(papersRegPattern)

	titleRegexp = regexp.MustCompile(titleRegPattern)
	authorRegexp = regexp.MustCompile(authorRegPattern)
	abstractRegexp = regexp.MustCompile(abstractRegPattern)
	keywordsRegexp = regexp.MustCompile(keywordsRegPattern)
	positionRegexp = regexp.MustCompile(positionRegPattern)
	pdflinkRegexp = regexp.MustCompile(pdflinkRegPattern)
)

func main() {
	var detailPages []string
	var err error
	var p *Paper
	var md *os.File

	// 1. 解析表格页
	if detailPages, err = parseIndexPage(); err != nil {
		log.Fatalln(err)
	}

	// 创建markdown文件
	if md, err = createMarkdownFile(); err != nil {
		log.Fatalln(err)
	}

	// 2. 解析详情页
	for i, v := range detailPages {
		// 获取页面内容并生成结构化数据
		p, err = parseDetailPage(v)
		if err != nil {
			log.Printf("parse [%d | %s] error: %v\n", i, v, err)
			continue
		}

		// 将p内容追加到文件
		content := p.MDString()
		fmt.Println(content)
		if err = writeAppendToMarkdown(md, content); err != nil {
			log.Printf("append [%d | %s] error: %v\n", i, v, err)
			continue
		}
	}
}





