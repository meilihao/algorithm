/*
233.困 数字 1 的个数

给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。

示例 1：

输入：n = 13
输出：6
示例 2：

输入：n = 0
输出：0


提示：

0 <= n <= 109
*/

package demo

import (
	"testing"
)

func TestCountDigitOne(t *testing.T) {

}

// 时间复杂度：O(logn)。n 包含的数位个数与 n 呈对数关系
// 核心思想：逐位统计法即个位数1+十位数1+百位数1....
/*
考虑一个数 n，当要统计数字 1 在某一位（例如，第 k 位，从右往左数，个位是第 0 位，十位是第 1 位，以此类推）上出现的次数时，通常会把 n 分成三部分：
- 高位部分 (high)：n / (mulk * 10)，即 n 在当前位以上的数字=a
- 当前位数字 (cur)：(n / mulk) % 10，即 n 在当前位的数字=cur
- 低位部分 (low)：n % mulk，即 n 在当前位以下的数字=b
*/
func countDigitOne(n int) (ans int) {
	// mulk 表示 10^k, k=0表示在个位
	// 在下面的代码中，可以发现 k 并没有被直接使用到（都是使用 10^k）
	// 但为了让代码看起来更加直观，这里保留了 k

	/*
		cur表示当前位的数字, 分开统计的方法见[这里](https://www.bilibili.com/video/BV1v5411J77K?spm_id_from=333.788.videopod.sections&vd_source=e7aa1f79d22918cf0263abc1e8e58c26)
		b = n%base
		a = n/base
		cur = a%10
		a = a / 10
		- cur<1: ans += a*base
		- cur=1: ans += (a*base + b +1)
		- cur>1: ans += (a+1)*base
	*/
	// 这里是混合了cur>=1
	for k, mulk := 0, 1; n >= mulk; k++ {
		// n%(mulk*10) = 当前位及其以下部分组成的数字
		// n%(mulk*10)-mulk+1 = 当前位为 1 的数字在当前不完整周期中出现的次数
		// max(..., 0)：如果 n%(mulk*10) - mulk + 1 算出来是负数（说明 n 的当前位比 1 小，或者 n 的当前位就是 1 但低位不够），那么就取 0，因为 1 不可能出现负数次
		// min(..., mulk)：如果 n 的当前位是 1，且其后位数的值非常大，导致 n%(mulk*10)-mulk+1 超过了 mulk（例如 n=199, mulk=10，199%100 - 10 + 1 = 90），实际上最多也只能出现 mulk 次（即 10 次）
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
		mulk *= 10
	}
	return
}
