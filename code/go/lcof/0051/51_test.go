/*
51：数组中的逆序对

在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组
成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
*/

package demo

import (
	"fmt"
	"testing"
)

func TestNInversePairs(t *testing.T) {
	s := []int{7, 5, 6, 4}
	fmt.Println(InversePairs2(s))
	fmt.Println(InversePairs(s))
}

func InversePairs(nums []int) int {
	return merge(nums, 0, len(nums)-1)
}

// 逆序对(x,y)=逆序对全在num[l,mid+1] + 逆序对全在num[mid+1,r] + 逆序对每个数在一部分(num[l,mid+1], num[mid+1,r])
// 空间复杂度: O(N*logN)
func merge(nums []int, l, r int) int {
	//递归的退出条件
	if l >= r {
		return 0
	}
	mid := (l + r) >> 1
	// 递归处理左边和右边
	res := merge(nums, l, mid) + merge(nums, mid+1, r)

	//单层递归的逻辑
	var tmp []int
	i, j := l, mid+1 // 合并两部分的起点
	for i <= mid && j <= r {
		if nums[i] <= nums[j] { //如果左边的比右边的小，正常放入合并后的tmp数组
			tmp = append(tmp, nums[i])
			i++
		} else { // nums[i] > nums[j]
			// 否则右边比左边小, 又因为两部分数组都是有序的,
			// 因此num[i,..,mid]均可是逆序对的一员x，逆序对的数量是 mid -i +1

			// 循环中x=num[i,..,mid]和下一次x=num[i+1,..,mid]是重复的, 不需要去重的原因: nums[j]也变化即`j++`
			res += mid - i + 1
			tmp = append(tmp, nums[j]) //将右边的放入tmp
			j++
		}
	}
	// 将剩余的未处理的数，直接加入到tmp中
	for i <= mid {
		tmp = append(tmp, nums[i])
		i++
	}
	for j <= r {
		tmp = append(tmp, nums[j])
		j++
	}
	//最后这里，将合并后的数组，覆盖原数组
	for i, v := range tmp {
		nums[l+i] = v
	}
	return res
}

// 冒泡排序: target=需要交换的次数
func InversePairs2(data []int) int {
	// write code here
	var length int = len(data)
	sum := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if data[i] > data[j] {
				sum++
			}
		}
	}
	answer := sum % 1000000007
	return answer
}
