/*
53.中 最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/
package demo

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//fmt.Println(maxSubArray(nums) == 6)
	fmt.Println(maxSubArray2(nums) == 6)
	fmt.Println(maxSubArray3(nums))

	nums2 := []int{-2, -1, -4}
	fmt.Println(maxSubArray2(nums2))
	fmt.Println(maxSubArray3(nums2))
}

// f(i) 代表以第 i 个数结尾的「连续子数组的最大和」, 如果i=4, 那么f(4)=max(nums[x:5]), 0<=x<=4
// f(i)=max{f(i−1)+nums[i],nums[i]}
func maxSubArray(nums []int) int {
	max := nums[0]
	dp := nums                       // nums[i] 表示以第 i 个元素结尾的最大子数组和
	for i := 1; i < len(nums); i++ { // 对于当前元素 nums[i]，判断是否将其加入子数组
		// nums[i-1]=dp[i-1]>0, 更新dp; 否则dp[i]=nums[i], 因为dp[i-1]<=0, 它对未来的连续子数组的和只会起到负面作用或无用
		if nums[i]+dp[i-1] > nums[i] { // 加入前面的子数组能获得更大和，更新 nums[i] 为这个和; 否则就取nums[i], 因为不变因此省略该else
			dp[i] = nums[i] + dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	result := nums[0]
	count := 0 // 维护当前连续子数组的和

	for _, v := range nums {
		count += v
		if count > result {
			result = count
		}
		if count < 0 { // 如果当前 count 变得小于 0，那么它对未来的连续子数组的和只会起到负面作用，所以要重置为 0，相当于从下一个元素开始重新计算一个新的连续子数组
			count = 0
		}
	}

	return result
}

// best
// 同maxSubArray
func maxSubArray3(nums []int) int {
	// pre, 维护当前连续子数组的和
	result, pre := nums[0], 0
	for _, num := range nums {
		pre = max(num, pre+num)
		result = max(result, pre)
	}

	return result
}
