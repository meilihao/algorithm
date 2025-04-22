package leetcode

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	s := []int{1, 2, 3}

	//fmt.Println(subsets(s))
	fmt.Println(subsets3(s))
}

func TestSubsetsOne(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(subsetsOne(1, s))
	fmt.Println(subsetsOne(3, s))
}

func subsetsOne(mask int, nums []int) []int {
	set := []int{}
	n := len(nums) - 1
	for i, v := range nums {
		fmt.Println(i, v, (mask>>i)&1 > 0)

		if (mask>>(n-i))&1 > 0 {
			set = append(set, v)
		}
	}

	return set
}

// best
// 迭代法实现子集枚举, 按位标记是否存在该数
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	max := 1 << n
	for mask := 0; mask < max; mask++ {
		set := []int{}
		for i, v := range nums {
			if (mask>>i)&1 > 0 { // 这里实际应该用`(mask>>((len(nums)-1)-i))&1 > 0`, 反转bits后才是`(mask>>i)&1 > 0`, 不过不影响实际输出结果, 仅影响输出顺序
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

func subsets2(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) { // 当遍历到决策树的叶子节点时，就终止了. 即当前路径搜索到末尾时，递归终止
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		fmt.Println("a:", cur, set)
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
		fmt.Println("b:", cur, set)
	}
	dfs(0)
	return
}

func subsets3(nums []int) (ans [][]int) {
	set := []int{} // 存放当前符合条件的结果

	var backtracking func([]int, int)
	backtracking = func(nums []int, index int) { // 正在考虑可选元素列表中第 index 个元素
		ans = append(ans, append([]int(nil), set...))
		if index >= len(nums) {
			return
		}
		for i := index; i < len(nums); i++ { //
			set = append(set, nums[i]) // 选择元素
			backtracking(nums, i+1)    // 递归搜索
			set = set[:len(set)-1]     // 撤销选择
		}
	}
	backtracking(nums, 0)
	return
}

func TestDfsDemo(t *testing.T) {
	//fmt.Println(dfsDemo(0, 3))
	fmt.Println(dfsDemo2([]int{1, 2, 3}))
}

// 找到一个长度为 n 的序列 a 的所有子序列，代码框架是这样的
// 时间复杂度是 O(2^n)
// `cur, n int` => a=[cur, n-1] & len(a)=n
func dfsDemo(cur, n int) (ans [][]int) {
	set := []int{} // 存放已经被选出的数字

	//dfs(cur,n) 参数表示当前位置是 cur，原序列总长度为 n
	var dfs func(int, int)
	dfs = func(cur, n int) {
		if cur == n {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, cur) // 考虑选择当前位置
		fmt.Println("a:", cur, set)
		dfs(cur+1, n)
		fmt.Println("b:", cur, set)
		set = set[:len(set)-1] // 考虑不选择当前位置
		dfs(cur+1, n)
		fmt.Println("c:", cur, set)
	}
	dfs(cur, n)
	return
}

// 找到一个长度为 n 的序列 a 的所有子序列，代码框架是这样的
func dfsDemo2(nums []int) (ans [][]int) {
	set := []int{} // 存放已经被选出的数字

	//dfs(cur,n) 参数表示当前位置是 cur，原序列总长度为 n
	var dfs func(int, int)
	dfs = func(cur, n int) {
		if cur == n {
			fmt.Println("r:", cur, n, set)
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur]) // 考虑选择当前位置
		fmt.Println("a:", cur, set)
		dfs(cur+1, n)
		fmt.Println("b:", cur, set)
		set = set[:len(set)-1] // 考虑不选择当前位置
		fmt.Println("c:", cur, set)
		dfs(cur+1, n)
		fmt.Println("d:", cur, set)
	}
	dfs(0, len(nums))
	return
}
