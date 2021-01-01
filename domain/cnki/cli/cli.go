/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/17 1:18
* @Description: provide a simple user commandline interaction
***********************************************************************/

package cli

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// 向cnki POST时转换成的查询条件字段
var conditionValues = map[string]string{
	"a": "SU$%=|",
	"b": "KY",
	"c": "TI",
	"d": "AB",
	"e": "FT",
	"f": "RF",
	"g": "CLC$=|??",
}

// 查询条件的中文名
var conditionNames = map[string]string{
	"a": "主题",
	"b": "关键词",
	"c": "篇名",
	"d": "摘要",
	"e": "全文",
	"f": "被引文献",
	"g": "中图分类号",
}

// 条件类型：与或非
var conditionTypes = map[string]string{
	"a": "and",
	"b": "or",
	"c": "not",
}

// 获取命令行输入
func GetUserInput(args url.Values) {
	setSearchCondition(args)
	searchYearRange(args)
	searchMagazineSource(args)
	//fmt.Println(args)	// 调试

	fmt.Println("正在检索中......")
	fmt.Println("-------------------------------------------------------")
}

// 获取检索条件
func setSearchCondition(args url.Values) {
	fmt.Println(`
	##################################################
	#	cnki searcher v0.1 ----by Eiger				 #
	#												 #
	#	请选择检索条件：（可多选）					   	 #
	#	(a) 主题		(b) 关键词	(c) 篇名	    (d) 摘要  #
	#   (e) 全文		(f) 被引文献	(g) 中图分类号		 #
	#												 #
	#	提示：若条件为空，则任选一项再直接回车				 #
	##################################################`)
	fmt.Println()

	// 获取搜索条件的类型列表
	var selectCondition string
	fmt.Print("> 请选择检索条件(英文逗号分割，如a,c): ")
	_, err := fmt.Scan(&selectCondition)
	if err != nil {
		log.Fatalf("scan condition failed: %v\n", err)
	}
	//fmt.Println(selectCondition)
	selectCondition = strings.TrimSpace(selectCondition)	// 去除两边空格

	var conds []string
	conds = strings.Split(selectCondition, ",")

	// 打印用户的选择
	var inputCheck string
	inputCheck = strings.Join(conds, " | ")
	fmt.Printf(`-------------------------------------------------------
您的选择是： %s
-------------------------------------------------------
`, inputCheck)

	// 获取各项参数具体的值。conds长度至少为1
	for idx, term := range conds {
		// 读取当前term的参数
		conditionVal := ""
		fmt.Printf("> 请输入【%s】：", conditionNames[term])
		_, err := fmt.Scan(&conditionVal)
		if err != nil {
			log.Fatalf("scan %s failed: %v\n", term, err)
		}

		prefix := "txt_" + strconv.Itoa(idx+1)

		// * 视作空或者任意. 	// 特殊情况：不需要搜索条件时，任选一个条件，输入*
		// 长度>1时不允许使用"*"值
		if conditionVal == "*" && len(conds) == 1 {
			// 填充args
			args.Set(prefix + "_sel", conditionValues[term])
			args.Set(prefix + "_value1", "")
			args.Set(prefix + "_relation", "#CNKI_AND")
			args.Set(prefix + "_special1", "=")
			break
		}

		if conditionVal == "*" && len(conds) > 1 {
			log.Fatalln("value * is not accept when condition number larger than 1 !")
		}

		// 条件类型（与或非）
		if idx != 0 {
			fmt.Printf("> 请输入【%s】条件类型：(a) 并且 (b) 或者 (c) 不含：", conditionNames[term])
			conditionType := ""
			_, err = fmt.Scan(&conditionType)
			if err != nil {
				log.Fatalf("scan condition %s type failed: %v\n", term, err)
			}
			// 填充args
			args.Set(prefix + "_logical", conditionTypes[conditionType])
		}

		// 填充args
		args.Set(prefix + "_sel", conditionValues[term])
		args.Set(prefix + "_value1", conditionVal)
		args.Set(prefix + "_relation", "#CNKI_AND")
		args.Set(prefix + "_special1", "=")
	}
}

// 获取检索时间范围
func searchYearRange(args url.Values) {
	fmt.Println("-------------------------------------------------------")
	// 获取输入
	searchYear := ""
	fmt.Print("> 是否限定搜索年限范围（y/n）：")
	_, err := fmt.Scan(&searchYear)
	if err != nil {
		log.Fatalf("scan search_year failed: %v\n", err)
	}
	switch searchYear {
	case "n", "N":
		return
	case "y", "Y":
		goto YEARRANGE
	default:
		log.Fatalln("incorrect search_year!")
	}

YEARRANGE:
	yearRange := ""
	fmt.Print("> 请输入搜索年限范围（例如2018, 2018-2020）：")
	_, err = fmt.Scan(&yearRange)
	if err != nil {
		log.Fatalf("scan year_range failed: %v\n", err)
	}

	// 处理输入
	yearFrom, yearTo := "", ""
	yearRange = strings.TrimSpace(yearRange)
	// 2018 (4) 2018-2020 (9)
	switch len(yearRange) {
	case 4:
		yearFrom, yearTo = yearRange, yearRange
	case 9:
		yearFrom, yearTo = yearRange[:4], yearRange[5:]
	default:
		log.Fatalln("incorrect year_range!")
	}

	args.Set("year_from", yearFrom)
	args.Set("year_to", yearTo)
}

// 获取检索期刊来源
func searchMagazineSource(args url.Values) {
	fmt.Println("-------------------------------------------------------")
	// 获取输入
	searchSource := ""
	fmt.Print("> 是否限定搜索期刊来源（y/n）：")
	_, err := fmt.Scan(&searchSource)
	if err != nil {
		log.Fatalf("scan search_source failed: %v\n", err)
	}
	switch searchSource {
	case "n", "N":
		return
	case "y", "Y":
		goto SEARCHSOURCE
	default:
		log.Fatalln("incorrect search_source!")
	}

SEARCHSOURCE:
	source := ""
	fmt.Print("> 请输入搜索期刊全称（例如 控制理论与应用）：")
	_, err = fmt.Scan(&source)
	if err != nil {
		log.Fatalf("scan magazine_source failed: %v\n", err)
	}

	// 处理输入
	args.Set("magazine_value1", source)
	args.Set("magazine_special1", "%")
}
