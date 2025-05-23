/*
LCR 186.简 文物朝代判断

展览馆展出来自 13 个朝代的文物，每排展柜展出 5 个文物。某排文物的摆放情况记录于数组 places，其中 places[i] 表示处于第 i 位文物的所属朝代编号。其中，编号为 0 的朝代表示未知朝代。请判断并返回这排文物的所属朝代编号是否能够视为连续的五个朝代（如遇未知朝代可算作连续情况）。

示例 1：

输入：places = [0, 6, 9, 0, 7]
输出：True

示例 2：

输入：places = [7, 8, 9, 10, 11]
输出：True

提示：

places.length = 5
0 <= places[i] <= 13
*/
package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestCheckDynasty(t *testing.T) {
	s := []int{0, 6, 9, 0, 7}
	fmt.Println(checkDynasty(s) == true)
}

// 顺子=连续+无重复
func checkDynasty(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++ //记录joker数量
		} else {
			if nums[i] == nums[i+1] { //除了0之外，有相同的元素直接返回false
				return false
			}
		}
	}
	// nums[4]-nums[joker] = 最大数字和最小非零数字之间的“跨度”（不包括王）
	// 如果这五张牌能构成顺子，跨度<=4
	return nums[4]-nums[joker] < 5
}
