/*
 * Medium
 *
 * 有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
 *
 * 示例 1:
 *
 * 输入: k = 5
 *
 * 输出: 9
 * 通过次数36,191
 * 提交次数64,563
 */

import "math"

func getKthMagicNumber(k int) int {
	arr := []int{1}
	ps := []int{3, 5, 7}
	is := []int{0, 0, 0}
	for len(arr) < k {
		minVal := math.MaxInt32
		ms := []int{0, 0, 0}
		for i := 0; i < 3; i++ {
			ms[i] = ps[i] * arr[is[i]]
			minVal = min(minVal, ms[i])
		}
		arr = append(arr, minVal)
		for i := 0; i < 3; i++ {
			for ; is[i] < len(arr) && ps[i]*arr[is[i]] <= minVal; is[i]++ {
			}
		}
	}
	return arr[k-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
