/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/5/12 15:26
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
	"golang.org/x/text/width"
	"testing"
)

func TestDBC2SBC(t *testing.T) {
	s := `。，（）-1！@234567890abc１２３４５６７８９ａｂｃ`
	// 全角转半角
	fmt.Println(width.Narrow.String(s))
	// 半角转全角
	fmt.Println(width.Widen.String(s))
}

func TestCheckSpaces(t *testing.T) {
	s := "天色渐晚,回家收衣服;好不好?   好。"
	// 应该改成 "天色渐晚, 回家收衣服; 好不好? 好。"
	fmt.Println(CheckSpaces(s))
}

func TestProcessAbstract(t *testing.T) {
	s := `摘要:本文针对一类具有未建模动态和预设性能的输出反馈非线性切换系统, 提出基于公共~Lyapunov~函数法的自适应输出反馈动态面控制方案.
通过设计 K 滤波器和观测器估计不可测量的状态. 引入动态信号处理动态不确定性. 利用 Nussbaum 函数解决增益符号未知的问题.
神经网络用于逼近由设计过程和理论分析所产生的未知连续函数. 引入性能函数和误差转换器将预设性能控制问题转换为稳定性问题.
通过适当选取切换子系统的初值, 并利用动态面控制系统证明的特点, 证明了闭环切换系统所有信号半全局一致终结有界. 仿真例子验证了所提方案的有效性.`
	fmt.Println(CheckSpaces(Period2Dot(DBC2SBC(s))))
}
