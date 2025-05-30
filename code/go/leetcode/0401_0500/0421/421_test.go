/*
421.中 数组中两个数的最大异或值

给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。

示例 1：

输入：nums = [3,10,5,25,2,8]
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
示例 2：

输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
输出：127

提示：

1 <= nums.length <= 2 * 105
0 <= nums[i] <= 231 - 1
*/
package leetcode

import (
	"fmt"
	"math/bits"
	"slices"
	"testing"
)

func TestFindMaximumXOR(t *testing.T) {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR3(nums))
}

const highBit = 30 // 整数最大值是 2^31 −1

type trie struct {
	left, right *trie // left 存储值为 0 的子节点，right 存储值为 1 的子节点, 存储方向: 最高->最低位
}

// 将所有数字做成tire
func (t *trie) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- { // 从最高位到最低位遍历数字的每一个二进制位
		bit := num >> i & 1 // 获取 num 的第 i 位
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}

func (t *trie) check(num int) (x int) { // 返回异或值
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		// 贪心策略：为了让异或结果的当前位为 1，在字典树中找到与 num 当前位相反的路径
		if bit == 0 {
			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1 // 异或结果的当前位是 1
			} else { //如果 1 的子节点不存在，只能走 0 这条路
				cur = cur.left
				x = x * 2
			}
		} else {
			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}

// 最大化 a⊕b，需要尽可能多地让 a 和 b 的对应二进制位不同
func findMaximumXOR(nums []int) (x int) {
	root := &trie{}
	for i := 1; i < len(nums); i++ {
		// 将 nums[i-1] 放入字典树，此时 nums[0 .. i-1] 都在字典树中
		root.add(nums[i-1])
		// 将 nums[i] 看作 ai，找出最大的 x 更新答案
		x = max(x, root.check(nums[i]))
	}
	return
}

func findMaximumXOR2(nums []int) (ans int) {
	highBit := bits.Len(uint(slices.Max(nums))) - 1 // slices.Max(nums): 找到 nums 数组中的最大值. bits.Len(value): 返回 value 的二进制表示中有效比特的位数比如bits.Len(5)=bits.Len(0b101)=3
	seen := map[int]bool{}
	mask := 0
	for i := highBit; i >= 0; i-- { // 从最高位开始枚举
		clear(seen)
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		for _, x := range nums {
			x &= mask // 低于 i 的比特位置为 0
			if seen[newAns^x] {
				ans = newAns // 这个比特位可以是 1
				break
			}
			seen[x] = true
		}
	}
	return
}

type trie2 struct {
	child [2]*trie2
}

func (t *trie2) add(num int) {
	cur := t

	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if cur.child[bit] == nil {
			cur.child[bit] = &trie2{}
		}

		cur = cur.child[bit]
	}
}

// best
func findMaximumXOR3(nums []int) (ans int) {
	root := &trie2{}
	for _, num := range nums {
		root.add(num)
	}

	res := 0

	for _, num := range nums {
		node := root
		xor := 0
		bit := 0

		for i := highBit; i >= 0; i-- {
			bit = num >> i & 1

			if node.child[1-bit] != nil {
				xor = (xor << 1) + 1
				node = node.child[1-bit]
			} else {
				xor <<= 1
				node = node.child[bit]
			}
		}

		res = max(res, xor)
	}

	return res
}
