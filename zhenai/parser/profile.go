/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/16 9:35
* @Description: Parse user profile
***********************************************************************/

package parser

import (
	"github.com/azd1997/go-crawler2/model"
	"github.com/azd1997/go-crawler2/zhenai/engine"
	"regexp"
	"strconv"
)

const (
	// `([\d]+)`即为岁数
	ageRegPattern = `<td><span class="label">年龄：</span>([\d]+)岁</td>`

	heightRegPattern = `<td><span class="label">身高：</span>([\d]+)CM</td>`
	weightRegPattern = `<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`
	incomeRegPattern = `<td><span class="label">月收入：</span>([^<]+)</td>`
	genderRegPattern = `<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`
	xingzuoRegPattern = `<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`
	marriageRegPattern = `<td><span class="label">婚况：</span>([^<]+)</td>`
	educationRegPattern = `<td><span class="label">学历：</span>([^<]+)</td>`
	occupationRegPattern = `<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`
	hukouRegPattern = `<td><span class="label">籍贯：</span>([^<]+)</td>`
	houseRegPattern = `<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`
	carRegPattern = `<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`
)

var (
	ageRegexp = regexp.MustCompile(ageRegPattern)
	heightRegexp = regexp.MustCompile(heightRegPattern)
	weightRegexp = regexp.MustCompile(weightRegPattern)
	incomeRegexp = regexp.MustCompile(incomeRegPattern)
	genderRegexp = regexp.MustCompile(genderRegPattern)
	xingzuoRegexp = regexp.MustCompile(xingzuoRegPattern)
	marriageRegexp = regexp.MustCompile(marriageRegPattern)
	educationRegexp = regexp.MustCompile(educationRegPattern)
	occupationRegexp = regexp.MustCompile(occupationRegPattern)
	hukouRegexp = regexp.MustCompile(hukouRegPattern)
	houseRegexp = regexp.MustCompile(houseRegPattern)
	carRegexp = regexp.MustCompile(carRegPattern)
)

// ParseProfile 解析用户信息。 name用户名
func ParseProfile(content []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	// name 从ParseCity得到的结果中取。如果因此修改ParseFunc签名，那么其实之前的很多解析函数都不需要
	// 这里可以采取修改ParseProfile签名，然后在其调用方ParseCity中去通过函数闭包来传入这个name
	profile.Name = name

	// age
	if age, err := strconv.Atoi(extractString(content, ageRegexp, 1)); err == nil {
		profile.Age = age
	}

	// height
	if height, err := strconv.Atoi(extractString(content, heightRegexp, 1)); err == nil {
		profile.Height = height
	}

	// weight
	if weight, err := strconv.Atoi(extractString(content, weightRegexp, 1)); err == nil {
		profile.Weight = weight
	}

	// income
	profile.Income = extractString(content, incomeRegexp, 1)

	// gender
	profile.Gender = extractString(content, genderRegexp, 1)

	// xingzuo
	profile.Xingzuo = extractString(content, xingzuoRegexp, 1)

	// marriage
	profile.Marriage = extractString(content, marriageRegexp, 1)

	// education
	profile.Education = extractString(content, educationRegexp, 1)

	// occupation
	profile.Occupation = extractString(content, occupationRegexp, 1)

	// hukou
	profile.Hukou = extractString(content, hukouRegexp, 1)

	// house
	profile.House = extractString(content, houseRegexp, 1)

	// car
	profile.Car = extractString(content, carRegexp, 1)

	return engine.ParseResult{
		Items: []interface{}{profile},
		// 没有新的request。当然以后如果有“猜你喜欢”，那么就会产生新的请求
	}
}

// index指目的字符串在匹配结果中的下标，这应当是在编写正则表达式时自己明确的
func extractString(content []byte, reg *regexp.Regexp, index int) string {
	match := reg.FindSubmatch(content)
	if len(match) > index {				// index应自己清楚
		return string(match[index])
	}
	return ""
}
