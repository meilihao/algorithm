/*
05. 替换空格

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
*/
package lcof

import (
	"fmt"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	fmt.Println(string(replaceSpace([]byte("A B"))) == `A%20B`)
	fmt.Println(string(replaceSpace([]byte(" A B"))) == `%20A%20B`)
	fmt.Println(string(replaceSpace([]byte("A B "))) == `A%20B%20`)
}

/*
1. 在字符串尾部填充任意字符，使得字符串的长度等于替换之后的长度。因为一个空格要替换成三个字符（%20），所以当遍历到一个空格时，需要在尾部填充两个任意字符。
2. 令 P1 指向字符串原来的末尾位置，P2 指向字符串现在的末尾位置。P1 和 P2 从后向前遍历，当 P1 遍历到一个空格时，就需要令 P2 指向的位置依次填充 02%（注意是逆序的），否则就填充上 P1 指向字符的值。从后向前遍是为了在改变 P2 所指向的内容时，不会影响到 P1 遍历原来字符串的内容。
3. 当 P2 遇到 P1 时（P2 <= P1），或者遍历结束（P1 < 0），退出。
*/
func replaceSpace(bs []byte) []byte {
	var i int
	var count int
	for i = 0; i < len(bs); i++ {
		if bs[i] == ' ' {
			count++
		}
	}

	if count == 0 {
		return bs
	}

	nbs := make([]byte, len(bs)+2*count)
	copy(nbs, bs)

	p1, p2 := len(bs)-1, len(nbs)-1

	var tmp byte
	for p1 >= 0 {
		tmp = bs[p1]
		if tmp == ' ' {
			nbs[p2] = '0'
			p2--
			nbs[p2] = '2'
			p2--
			nbs[p2] = '%'
			p2--
		} else {
			nbs[p2] = bs[p1]
			p2--
		}

		p1--
	}

	return nbs
}
