/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/5/12 15:20
* @Description: The file is for
***********************************************************************/

package main

import (
	"golang.org/x/text/width"
)


// 将字符串中所有全角符号转为半角符号
func DBC2SBC(s string) string {
	return width.Narrow.String(s)
}

// 将字符串中所有半角符号转为全角符号
func SBC2DBC(s string) string {
	return width.Widen.String(s)
}

// 由于半角符号中也有句号，而期刊摘要要求使用.号分隔而非句号
func Period2Dot(s string) string {
	s1 := []rune(s)
	for i, v := range s1 {
		if v == '｡' {
			s1[i] = '.'
		}
	}
	return string(s1)
}

// 处理符号和汉字后的空格：
// 符号后有且仅有一个空格相连
// 汉字后不能有空格
func CheckSpaces(s string) string {
	s1, s2 := []rune(s), []rune(s)
	p1, p2 := 0, 0
	for p1 < len(s1) {
		//fmt.Println(p1, p2, string(s1), " || ", string(s2))

		switch s1[p1] {
		case '.', ';', ',', '?', ':':	// 查看右边有没有空格，没有则加上(p1不动，p2加1，s2添加)
			//fmt.Println(p1, p2)
			if s1[p1] == '.' && (p1+1 < len(s1) && IsDigit(s1[p1+1]) ) {	// 处理小数
				p1++;p2++
			} else if p1+1 < len(s1) && s1[p1+1] != ' ' {
				//fmt.Println("s2", string(s2))
				s2T := append([]rune{}, s2[:p2+1]...)
				s2T = append(s2T, ' ')
				// 注意！如果是直接s2T = append(s2[:p2+1], ' ')
				// 修改了原数组地址上的内容！！！好坑！！！
				//fmt.Println("s2T", string(s2T))
				s2 = append(s2T, s2[p2+1:]...)
				//fmt.Println("s2", string(s2))
				p1++; p2 = p2 + 2
			} else {
				p1++; p2++
			}


		case ' ':	// 如果空格的前一个字符不是符号，则将当前空格删去（p1右移,p2不动，s2去掉当前空格）
			if p1-1 >= 0 && !IsSymbol(s1[p1-1]) {
				p1++
				s2 = append(s2[:p2], s2[p2+1:]...)
				continue
			}
			p1++; p2++

		case '\r', '\n':	// 去除换行符，不管是哪个系统下，把这两个符号删掉
			s2 = append(s2[:p2], s2[p2+1:]...)
			p1++

		default:	// 普通汉字或英文字符
			p1++	// p1++意味着检查下一个字符
			p2++
		}

	}

	return string(s2)
}


func IsSymbol(r rune) bool {
	switch r {
	case '.', ';', ',', '?', ':':
		return true
	default:
		return false
	}
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

