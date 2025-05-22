/*
40. 最小的k个数

输入n个整数，找出其中最小的k个数。例如输入4、5、1、6、2、7、3、8
这8个数字，则最小的4个数字是1、2、3、4。
*/

package demo

import "testing"

/*
思路:
1.会修改原数组: quicSort, target = k-1, 这样, nums[:k]即符号条件的数组
2.不会修改原数组: 大顶堆heap(size=k), 如果cur < max(heap), 用cur取代max(heap), 再堆化
*/
func TestGetLeastNumbers(t *testing.T) {

}

func getLeastNumbers(arr []int, k int) []int {
	for i := len(arr); i > len(arr)-k; i-- {
		heap(arr[:i])
		arr[0], arr[i-1] = arr[i-1], arr[0]
	}
	return arr[len(arr)-k:]
}

func heap(arr []int) {
	len := len(arr)
	for i := len/2 - 1; i >= 0; i-- {
		if 2*i+1 < len && arr[i] > arr[2*i+1] { // 存在左节点
			arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
		}
		if 2*i+2 < len && arr[i] > arr[2*i+2] { // 存在右节点
			arr[i], arr[2*i+2] = arr[2*i+2], arr[i]
		}
	}
}
