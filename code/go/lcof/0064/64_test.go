/*
LCR 189.中 设计机械累加器

请设计一个机械累加器，计算从 1、2... 一直累加到目标数值 target 的总和。注意这是一个只能进行加法操作的程序，不具备乘除、if-else、switch-case、for 循环、while 循环，及条件判断语句等高级功能。

示例 1：

输入: target = 5
输出: 15
示例 2：

输入: target = 7
输出: 28

提示：

1 <= target <= 10000
*/
package main

import (
	"fmt"
	"testing"
)

/*
other: 快速乘
*/
func TestMechanicalAccumulator(t *testing.T) {
	n := 5
	fmt.Println(mechanicalAccumulator(n) == 15)
}

func mechanicalAccumulator(target int) int {
	ans := 0
	var sum func(int) bool
	sum = func(target int) bool {
		ans += target
		return target > 0 && sum(target-1)
	}
	sum(target)
	return ans
}
