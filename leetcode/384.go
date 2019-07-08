package main

import (
	"math/rand"
	"time"
	"fmt"
)

type Solution struct {
	Origin []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	
	return Solution{
		Origin: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.Origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	tmp := make([]int, len(this.Origin))
	copy(tmp, this.Origin)

	// Fisher–Yates shuffle 洗牌算法
	for i := len(tmp) - 1; i > 0; i-- {
		r := rand.Intn(i + 1)
		tmp[i], tmp[r] = tmp[r], tmp[i]
	}

	return tmp
}

func main() {
	nums := []int{1,2,3,4,5,6}
	solution := Constructor(nums)

	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
}
