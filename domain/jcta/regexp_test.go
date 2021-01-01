/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/22 20:59
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp1(t *testing.T) {
	contentStr := `<li><div><input class="class1" id="CCTA200068" 
name="file_no" type="checkbox" value="CCTA200068" /></div> 
<b><a href="reader/view_abstract.aspx?file_no=CCTA200068&flag=1">局部新冠肺炎时滞模型及再生数的计算</a></b> 
<span>刘可伋,江渝,严阅,陈文斌</span><span><em>2020,37(3):453-460</em><em  class="zyao"><i class="zayao">
<a href="reader/view_abstract.aspx?file_no=CCTA200068&flag=1" target='_blank'>摘要</a>
(<font color="#FF0000">53</font>)</a></i><i class="pdf"><a href="reader/create_pdf.aspx?file_no=CCTA200068&year_id=2020&quarter_id=3&falg=1">PDF</a>(<font color="#FF0000">36</font>)</i></em></span><span style="display:none" class="yc">
2019年末的新型冠状病毒肺炎(简称: 新冠肺炎, 又称COVID-19, Novel Coronavirus Pneumonia, NCP, 2019-nCoV)疫情得到了全球的广泛关注.文献\cite{ChenArxiv2020, YanPreprint2020}提出
了一类新的时滞动力学系统的新冠肺炎传播模型 (A Time Delay Dynamic model for NCP, 
简称TDD-NCP模型)来描述疫情的传播过程. 本文将这个模型用于研究部分省市的疫情传播问题, 
通过增加模型的源项用于模拟外来潜伏感染者对于当地疫情的影响. 基于全国各级卫健委每日公布的累
计确诊数与治愈数, 我们有效地模拟并预测了各地疫情的发展. 我们提出了基于TDD模型的再生数的两种
计算方法, 并做了估计与分析. 我们发现疫情暴发初期再生数较大, 但随着各级政府防控力度的加大而
逐渐减小. 最后, 我们分析了返程潮对上海疫情发展的影响, 并建议上海市政府继续加大防控力度, 以
防疫情二次暴发.</span></li>`

	// 正则表达式匹配所有文章
	// <a href="reader/view_abstract.aspx?file_no=CCTA200068&flag=1">局部新冠肺炎时滞模型及再生数的计算</a>
	// reader/view_abstract.aspx?file_no=CCTA200068&flag=1
	// 局部新冠肺炎时滞模型及再生数的计算
	// 注意?是特殊字符，需要转义
	papersRegPattern := `<a href="(reader/view_abstract.aspx\?file_no=CCTA\d{6}&flag=1)">([^<]+)</a>`
	papersRegexp := regexp.MustCompile(papersRegPattern)
	matched := papersRegexp.FindAllStringSubmatch(contentStr, -1)
	fmt.Println(len(matched))
	for i, v := range matched {
		fmt.Println(i, v)
	}

	if len(matched) != 1 ||
		matched[0][0] != `<a href="reader/view_abstract.aspx?file_no=CCTA200068&flag=1">局部新冠肺炎时滞模型及再生数的计算</a>` ||
		matched[0][1] != `reader/view_abstract.aspx?file_no=CCTA200068&flag=1` ||
		matched[0][2] != `局部新冠肺炎时滞模型及再生数的计算` {
		t.Error("wrong")
	}
}


func TestRegexp2(t *testing.T) {
	contentStr := `<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>基于多搜索器优化算法的含可再生能源协同优化调度-Collaborative optimization scheduling with renewable energy based on multi-searcher optimization algorithm</title>
<meta http-equiv="Content-Language" content="zh-cn">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="keywords" content="风力发电厂；光伏发电厂；热电联产系统；混沌理论；多搜索器算法"/>
<meta name="HW.ad-path" content="http://jcta.alljournals.ac.cn/cta_cn/ch/reader/view_abstract.aspx?file_no=CCTA190150"/>
<meta name="HW.identifier" content="http://jcta.alljournals.ac.cn/cta_cn/ch/reader/view_abstract.aspx?file_no=CCTA190150"/>
<meta name="DC.Format" content="text/html"/>
<meta name="DC.Language" content="cn"/>
<meta name="DC.Title" content="基于多搜索器优化算法的含可再生能源协同优化调度"/>
<meta name="DC.Identifier" content="10.7641/CTA.2019.90150"/>
<meta name="DC.Contributor" content="唐建林"/>
<meta name="DC.Contributor" content="余 涛"/>
<meta name="DC.Contributor" content="张孝顺"/>
<meta name="DC.Contributor" content="李卓环"/>
<meta name="DC.Contributor" content="陈俊斌"/>
<meta name="DC.Date" content=""/>
<meta name="DC.Keywords" content="风力发电厂；光伏发电厂；热电联产系统；混沌理论；多搜索器算法"/>
<meta name="DC.Keywords" content="Wind power plant; photovoltaic power plant; combined heat and power system; chaos theory; multi-searcher algorithm"/>
<meta name="Description" content="考虑到风光不确定性给能源系统运行带来的影响，本文建立了含有风力、光伏的热电数学模型，并提出了多搜索器优化算法，用于解决热电能量管理优化问题。混沌理论的引入能够不断缩小优化变量的搜索空间，并不断提高搜索精度，从而有较高的搜索效率。算法中的双层搜索器的设置既能保证快速收敛到近似最优解，也能及时避免陷入局部最优解。通过标准函数和27机组热电系统的仿真结果表明，本方法具有较高收敛稳定性且收敛速度快，能够有效解决热电能量管理系统中的高度非线性、非光滑、非凸问题。;Considering the influence of uncertainty of wind and solar on the operation of energy system, this paper establishes a combined heat and power mathematical model containing wind and photovoltaic, and proposes a multi-searcher optimization algorithm to solve the thermoelectric energy management optimization problem. The introduction of chaos theory can continuously reduce the search space of optimized variables, and continuously improve the search accuracy, thus having higher search efficiency. The setting of the two-layer searcher in the algorithm can ensure fast convergence to the approximate optimal solution, and can avoid falling into the local optimal solution in time. The standard function and simulation results show that the proposed method has high convergence stability and fast convergence speed, and can effectively solve the highly nonlinear, non-smooth and non-convex problems in the thermoelectric energy management system."/>
<meta name="Keywords" content="风力发电厂；光伏发电厂；热电联产系统；混沌理论；多搜索器算法;Wind power plant; photovoltaic power plant; combined heat and power system; chaos theory; multi-searcher algorithm"/>
<meta name="citation_title" content="基于多搜索器优化算法的含可再生能源协同优化调度"/>
<meta name="citation_journal_title" content="控制理论与应用"/>
<meta name="citation_author" content="唐建林"/>
<meta name="citation_author" content="余 涛"/>
<meta name="citation_author" content="张孝顺"/>
<meta name="citation_author" content="李卓环"/>
<meta name="citation_author" content="陈俊斌"/>
<meta name="citation_volume" content="37"/>
<meta name="citation_issue" content="3"/>
<meta name="citation_date" content="20200420"/>
<meta name="citation_firstpage" content="492"/>
<meta name="citation_lastpage" content="504"/>`

	// 正则表达式匹配所有文章
	titleRegPattern := `<meta name="DC\.Title" content="([^{/>}]+)"/>`
	titleRegexp := regexp.MustCompile(titleRegPattern)
	matched := titleRegexp.FindAllStringSubmatch(contentStr, -1)
	fmt.Println(len(matched))
	for i, v := range matched {
		fmt.Println(i, v)
	}

	if len(matched) != 1 ||
		matched[0][0] != `<meta name="DC.Title" content="基于多搜索器优化算法的含可再生能源协同优化调度"/>` ||
		matched[0][1] != `基于多搜索器优化算法的含可再生能源协同优化调度` {
		t.Error("wrong")
	}
}


func TestRegexp3(t *testing.T) {
	contentStr := `<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>基于多搜索器优化算法的含可再生能源协同优化调度-Collaborative optimization scheduling with renewable energy based on multi-searcher optimization algorithm</title>
<meta http-equiv="Content-Language" content="zh-cn">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="keywords" content="风力发电厂；光伏发电厂；热电联产系统；混沌理论；多搜索器算法"/>
<meta name="HW.ad-path" content="http://jcta.alljournals.ac.cn/cta_cn/ch/reader/view_abstract.aspx?file_no=CCTA190150"/>
<meta name="HW.identifier" content="http://jcta.alljournals.ac.cn/cta_cn/ch/reader/view_abstract.aspx?file_no=CCTA190150"/>
<meta name="DC.Format" content="text/html"/>
<meta name="DC.Language" content="cn"/>
<meta name="DC.Title" content="基于多搜索器优化算法的含可再生能源协同优化调度"/>
<meta name="DC.Identifier" content="10.7641/CTA.2019.90150"/>
<meta name="DC.Contributor" content="唐建林"/>
<meta name="DC.Contributor" content="余 涛"/>
<meta name="DC.Contributor" content="张孝顺"/>
<meta name="DC.Contributor" content="李卓环"/>
<meta name="DC.Contributor" content="陈俊斌"/>
<meta name="DC.Date" content=""/>
<meta name="DC.Keywords" content="风力发电厂；光伏发电厂；热电联产系统；混沌理论；多搜索器算法"/>
<meta name="DC.Keywords" content="Wind power plant; photovoltaic power plant; combined heat and power system; chaos theory; multi-searcher algorithm"/>
<meta name="Description" content="考虑到风光不确定性给能源系统运行带来的影响，本文建立了含有风力、光伏的热电数学模型，并提出了多搜索器优化算法，用于解决热电能量管理优化问题。混沌理论的引入能够不断缩小优化变量的搜索空间，并不断提高搜索精度，从而有较高的搜索效率。算法中的双层搜索器的设置既能保证快速收敛到近似最优解，也能及时避免陷入局部最优解。通过标准函数和27机组热电系统的仿真结果表明，本方法具有较高收敛稳定性且收敛速度快，能够有效解决热电能量管理系统中的高度非线性、非光滑、非凸问题。;Considering the influence of uncertainty of wind and solar on the operation of energy system, this paper establishes a combined heat and power mathematical model containing wind and photovoltaic, and proposes a multi-searcher optimization algorithm to solve the thermoelectric energy management optimization problem. The introduction of chaos theory can continuously reduce the search space of optimized variables, and continuously improve the search accuracy, thus having higher search efficiency. The setting of the two-layer searcher in the algorithm can ensure fast convergence to the approximate optimal solution, and can avoid falling into the local optimal solution in time. The standard function and simulation results show that the proposed method has high convergence stability and fast convergence speed, and can effectively solve the highly nonlinear, non-smooth and non-convex problems in the thermoelectric energy management system."/>
<meta name="Keywords" content="风力发电厂；光伏发电厂；热电联产系统；混沌理论；多搜索器算法;Wind power plant; photovoltaic power plant; combined heat and power system; chaos theory; multi-searcher algorithm"/>
<meta name="citation_title" content="基于多搜索器优化算法的含可再生能源协同优化调度"/>
<meta name="citation_journal_title" content="控制理论与应用"/>
<meta name="citation_author" content="唐建林"/>
<meta name="citation_author" content="余 涛"/>
<meta name="citation_author" content="张孝顺"/>
<meta name="citation_author" content="李卓环"/>
<meta name="citation_author" content="陈俊斌"/>
<meta name="citation_volume" content="37"/>
<meta name="citation_issue" content="3"/>
<meta name="citation_date" content="20200420"/>
<meta name="citation_firstpage" content="492"/>
<meta name="citation_lastpage" content="504"/>`

	authorsReg := `<meta name="DC\.Contributor" content="([^"]+)"/>`
	authorsRegexp := regexp.MustCompile(authorsReg)
	matched := authorsRegexp.FindAllStringSubmatch(contentStr, -1)
	fmt.Println(len(matched))
	for i, v := range matched {
		fmt.Println(i, v)
	}

	if len(matched) == 0 {
		t.Error("wrong")
	}
}