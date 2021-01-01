/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/17 7:11
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
	"github.com/azd1997/go-crawler2/domain/cnki/cli"
	"time"
)

const timeFormat = "%s %s %d %d:%d:%d %d GMT+0800 (中国标准时间)"

func main() {
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
	fmt.Println(nowStr)

	args := make(map[string]string)
	cli.GetUserInput(args)
}
