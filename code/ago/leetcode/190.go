// 190. 颠倒二进制位
package main

import "fmt"

func main() {
	var n uint32
	n = 0b11111111111111111111111111111101
	fmt.Printf("%d, %d", n, reverseBits(n))
}
func reverseBits(num uint32) uint32 {
	var result,temp uint32
	for i := 0; i < 32 ;i++ {
		temp = num >> i
		//1.取出要反转的每一位二进制
		temp &= 1
		//2.将该二进制位放到反转后的位置
		temp  <<= (31-i)
		//3.将该位添加到结果上
		result |= temp
	}
	return result
}

func reverseBits2(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32 ;i++ {
		result<<=1 // result = result<<1
		result+=num&1
		num>>=1 // num = num >> 1
	}
	return result
}
