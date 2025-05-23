/*
LCR 187.简 破冰游戏

社团共有 num 位成员参与破冰游戏，编号为 0 ~ num-1。成员们按照编号顺序围绕圆桌而坐。社长抽取一个数字 target，从 0 号成员起开始计数，排在第 target 位的成员离开圆桌，且成员离开后从下一个成员开始计数。请返回游戏结束时最后一位成员的编号。

示例 1：

输入：num = 7, target = 4
输出：1
示例 2：

输入：num = 12, target = 5
输出：0

提示：

1 <= num <= 10^5
1 <= target <= 10^6
*/
package main

import (
	"fmt"
	"testing"
)

/*
other:
1. 环形链表
*/
// 约瑟夫环: f（n） = (f(n-1）+m)%n
func TestReverseWords(t *testing.T) {
	num := 7
	target := 4

	fmt.Println(iceBreakingGame(num, target) == 1)
	fmt.Println(iceBreakingGame2(num, target) == 1)
}

// 递归
// https://github.com/francistao/LearningNotes/blob/master/Part3/Algorithm/%E5%89%91%E6%8C%87Offer/%E9%9D%A2%E8%AF%95%E9%A2%9845%EF%BC%9A%E5%9C%86%E5%9C%88%E4%B8%AD%E6%9C%80%E5%90%8E%E5%89%A9%E4%B8%8B%E7%9A%84%E6%95%B0%E5%AD%97.md
func iceBreakingGame2(num int, target int) int {
	if num == 1 {
		return 0
	}
	return (iceBreakingGame2(num-1, target) + target) % num
}

// 时间复杂度：O(num)，需要求解的函数值有 num 个
// iceBreakingGame2的迭代版
func iceBreakingGame(num int, target int) int {
	f := 0
	for i := 2; i != num+1; i++ {
		f = (target + f) % i
	}
	return f
}
