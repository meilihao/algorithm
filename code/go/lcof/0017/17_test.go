/*
17. 打印1到最大的n位数

输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则
打印出1、2、3一直到最大的3位数即999
*/
package demo

import (
	"fmt"
	"math"
	"testing"
)

// 关键是越界
func TestPrintNumbers(t *testing.T) {
	//fmt.Println(printNumbers(2))
	fmt.Println(printNumbers2(2))
}

func printNumbers(n int) []int {
	num := int(math.Pow10(n))
	ans := []int{}
	for i := 1; i < int(num); i++ {
		ans = append(ans, i)
	}
	return ans
}

func printNumbers2(n int) []string {
	res := make([]string, 0)

	num := make([]byte, n)
	for i := 0; i < n; i++ {
		num[i] = '0'
	}

	nice := 0      // 数字各位中 9 的数量为 nine, 判断何时需要移动左边界
	start := n - 1 // 字符串的左边界，以保证添加的数字字符串 num[start:] 中无高位多余的 0

	var dfs func(int)
	dfs = func(x int) {
		if x == n {
			s := string(num[start:])
			if s != "0" { // 跳过"0", 因为题目从1开始
				res = append(res, s) // 如果需要返回int, 这里可以加ParseInt()
			}
			fmt.Println(s, n, start, nice)
			// 当输出数字的所有位都是 9 时，则下个数字需要向更高位进 1 ，此时左边界 start 需要减 1 （即高位多余的 0 减少一个）
			if n-start == nice {
				start -= 1
			}
			return
		}

		for i := 0; i < 10; i++ {
			if i == 9 {
				nice += 1
			}

			num[x] = '0' + byte(i) // 固定第 x 位为 i
			dfs(x + 1)             // 开启固定第 x + 1 位
		}
		nice -= 1 // 回溯前恢复 nine=nine−1
	}

	dfs(0)

	return res
}
