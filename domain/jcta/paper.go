/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/23 8:55
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
	"strings"
)

const mdFormat1 = `

## %s

**作者:** %s

**摘要:** %s

**关键词:** %s

**引用格式:** %s

**全文链接:** <%s>
`

const mdFormat2 = `

## %s

**作者** %s

**摘要** %s

**关键词** %s

**引用格式** %s

**全文链接** <%s>
`

type Paper struct {
	AbstractUrl string
	Title       string
	Authors []string
	Abstract string
	Keywords string
	Position string	// 第几卷第几期
	Reference string	// 引用格式按规则拼接
	PdfUrl string
}

func (p *Paper) MDString() string {

	// 检查引用格式是否设置
	if p.Reference == "" {
		// 拼接引用格式
		p.setReference()
	}

	return fmt.Sprintf(mdFormat2, p.Title, strings.Join(p.Authors, ", "), p.Abstract, p.Keywords, p.Reference, p.PdfUrl)
}

// 设置标题
func (p *Paper) setTitle(content string) {
	matched := titleRegexp.FindAllStringSubmatch(content, -1)
	if len(matched) > 0 {
		p.Title = matched[0][1]
	}
}

// 设置作者
func (p *Paper) setAuthors(content string) {
	matched := authorRegexp.FindAllStringSubmatch(content, -1)
	authors := make([]string, 0)
	if len(matched) > 0 {
		for _, v := range matched {
			// 对于中文名，希望把中文名间的空格去掉，而英文名的空格则保留
			// 中文名一般就全是中文，英文名则是英文字母
			// 这里简单按照第一个rune去判断是否是中英字符
			name := []rune(v[1])
			if (name[0] >= 'a' && name[0] <= 'z') || (name[0] >= 'A' && name[0] <= 'Z') {
				// 英文名不作处理
				authors = append(authors, v[1])
			} else {	// 视作中文名，要去除空格
				tmp := strings.Split(v[1], " ")
				authors = append(authors, strings.Join(tmp, ""))
			}
		}
	}
	p.Authors = authors
}

// 设置摘要
func (p *Paper) setAbstract(content string) {
	matched := abstractRegexp.FindAllStringSubmatch(content, -1)
	if len(matched) == 1 {
		p.Abstract = CheckSpaces(Period2Dot(DBC2SBC(matched[0][1])))
	}
}

// 设置关键词
func (p *Paper) setKeywords(content string) {
	matched := keywordsRegexp.FindAllStringSubmatch(content, -1)
	if len(matched) == 1 {
		p.Keywords = matched[0][1]
	}
}

// 设置pdf链接
func (p *Paper) setPDFLink(content string) {
	matched := pdflinkRegexp.FindAllStringSubmatch(content, -1)
	if len(matched) == 1 {
		p.PdfUrl = pdfPagePrefix + matched[0][1]
	}
}

// 设置引用格式
func (p *Paper) setPosition(content string) {
	matched := positionRegexp.FindAllStringSubmatch(content, -1)
	if len(matched) == 1 {
		p.Position = matched[0][1]
	}
}

// 设置引用格式
func (p *Paper) setReference() {

	var authors string
	if len(p.Authors) > 3 {
		authors = strings.Join(p.Authors[:3], ", ") + ", 等. "
	} else {
		authors = strings.Join(p.Authors, ", ") + ". "
	}

	p.Reference = authors + p.Title + ". " + journal + ", " + p.Position
}

