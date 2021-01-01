/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/23 9:16
* @Description: The file is for
***********************************************************************/

package main

import (
	"os"
	"time"

	_ "github.com/russross/blackfriday"		// 这是用于markdown解析的一个库，这里暂时用不到
)

func createMarkdownFile() (*os.File, error) {
	// 创建新文件，用时间戳命名
	filename := "./" + time.Now().Format(time.RFC3339)[:13] + ".md"	// "2006-01-02T15:04:05Z07:00" => "2006-01-02T15"
	return os.Create(filename)
}

// 将内容追加写到markdown文件
// 可以直接使用文件读写方式实现
func writeAppendToMarkdown(file *os.File, content string) error {
	// file追加写内容
	_, err := file.WriteString(content)
	return err
}

func writeAppendToMarkdownViaBlackFriday() {

}
