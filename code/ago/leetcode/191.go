// 191. 位1的个数
package main

import "fmt"

func main() {
	var n uint32
	n = 0b00000000000000000000000000001011
	fmt.Println(hammingWeight(n))
}

func hammingWeight(num uint32) int {
	res:=0
	for num>0{
		res++

		// x = x & (x-1) // 清除序列最后面(即最低位)的1
		num &= (num-1)
	}

	return res
}
