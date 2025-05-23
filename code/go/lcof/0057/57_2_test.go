/*
LCR 180.简 文件组合

待传输文件被切分成多个部分，按照原排列顺序，每部分文件编号均为一个 正整数（至少含有两个文件）。传输要求为：连续文件编号总和为接收方指定数字 target 的所有文件。请返回所有符合该要求的文件传输组合列表。

注意，返回时需遵循以下规则：

每种组合按照文件编号 升序 排列；
不同组合按照第一个文件编号 升序 排列。

示例 1：

输入：target = 12
输出：[[3, 4, 5]]
解释：在上述示例中，存在一个连续正整数序列的和为 12，为 [3, 4, 5]。
示例 2：

输入：target = 18
输出：[[3,4,5,6],[5,6,7]]
解释：在上述示例中，存在两个连续正整数序列的和分别为 18，分别为 [3, 4, 5, 6] 和 [5, 6, 7]。

提示：

1 <= target <= 10^5
*/
package main

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	fmt.Println(findContinuousSequence(15))
}

func findContinuousSequence(target int) (res [][]int) {
	// 初始化滑动窗口的变量：
	// i：窗口的左边界（起始数字）。从 1 开始。
	// j：窗口的右边界（结束数字）。从 2 开始。
	// sum：当前窗口 [i, j] 内所有数字的和

	i, j, sum := 1, 2, 3

	for i <= target/2 { // i的上限是target/2
		if target > sum { // 当前和太小, 需要扩大窗口
			j++
			sum += j
		} else {
			if target == sum {
				tmp := make([]int, j-i+1)
				for k := i; k <= j; k++ {
					tmp[k-i] = k
				}
				res = append(res, tmp)
			}

			// sum >= target
			// sum> target: 需要缩小窗口, 以符合预期
			// sum = target : 需要缩小窗口, 求下一个i
			sum -= i
			i++
		}
	}
	return
}
