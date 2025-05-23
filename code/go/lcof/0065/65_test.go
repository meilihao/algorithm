/*
LCR 190.简 加密运算

计算机安全专家正在开发一款高度安全的加密通信软件，需要在进行数据传输时对数据进行加密和解密操作。假定 dataA 和 dataB 分别为随机抽样的两次通信的数据量：

正数为发送量
负数为接受量
0 为数据遗失
请不使用四则运算符的情况下实现一个函数计算两次通信的数据量之和（三种情况均需被统计），以确保在数据传输过程中的高安全性和保密性。

示例 1：

输入：dataA = 5, dataB = -1
输出：4

提示：

dataA 和 dataB 均可能是负数或 0
结果不会溢出 32 位整数
*/
/*
65：不用加减乘除做加法

题目：写一个函数，求两个整数之和，要求在函数体内不得使用＋、－、×、÷四则运算符号。
*/
package main

import (
	"fmt"
	"testing"
)

func TestEncryptionCalculate(t *testing.T) {
	fmt.Println(encryptionCalculate(5, -1) == 4)
}

func encryptionCalculate(a int, b int) int {
	carry := 0
	for b != 0 { //相当于进位不等于0
		carry = (a & b) << 1 //算出进位, 1&1=1, 需要左移进位
		a ^= b               //算出不带进位的和, 1^0=1, 0^0=0(忽略),1^1=0(上面已进位)
		b = carry            //更新进位, b 在每次循环中都会被更新为上次计算出的进位
	}
	return a
}
