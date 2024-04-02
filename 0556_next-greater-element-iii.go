/*
 * @lc app=leetcode.cn id=next-greater-element-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [556] 下一个更大元素 III
 *
 * https://leetcode.cn/problems/next-greater-element-iii/description/
 *
 * algorithms
 * Medium (33.55%)
 * Total Accepted:    31.5K
 * Total Submissions: 87.9K
 * Testcase Example:  '12'
 *
 * 给你一个正整数 n ，请你找出符合条件的最小整数，其由重新排列 n 中存在的每位数字组成，并且其值大于 n 。如果不存在这样的正整数，则返回 -1 。
 *
 * 注意 ，返回的整数应当是一个 32 位整数 ，如果存在满足题意的答案，但不是 32 位整数 ，同样返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 12
 * 输出：21
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 21
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 1 << 31 - 1
 *
 *
 */

import (
	"math"
	"sort"
)

func nextGreaterElement(n int) int {
	var arr []int
	for ; n > 0; n /= 10 {
		arr = append(arr, n%10)
	}
	m := len(arr)
	for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	i, j := m-1, m-1
	for ; i > 0; i-- {
		if arr[i] > arr[i-1] {
			break
		}
	}
	if i == 0 {
		return -1
	}
	i--
	for ; j >= i; j-- {
		if arr[j] > arr[i] {
			break
		}
	}
	arr[i], arr[j] = arr[j], arr[i]
	sort.Ints(arr[i+1:])
	ans := 0
	for _, a := range arr {
		ans = ans*10 + a
	}
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}
