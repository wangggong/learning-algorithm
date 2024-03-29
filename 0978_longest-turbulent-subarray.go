/*
 * @lc app=leetcode.cn id=longest-turbulent-subarray lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [978] 最长湍流子数组
 *
 * https://leetcode-cn.com/problems/longest-turbulent-subarray/description/
 *
 * algorithms
 * Medium (47.39%)
 * Total Accepted:    45.3K
 * Total Submissions: 95.7K
 * Testcase Example:  '[9,4,2,10,7,8,8,1,9]'
 *
 * 给定一个整数数组 arr ，返回 arr 的 最大湍流子数组的长度 。
 *
 * 如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是 湍流子数组 。
 *
 * 更正式地来说，当 arr 的子数组 A[i], A[i+1], ..., A[j] 满足仅满足下列条件时，我们称其为湍流子数组：
 *
 *
 * 若 i <= k < j ：
 *
 *
 * 当 k 为奇数时， A[k] > A[k+1]，且
 * 当 k 为偶数时，A[k] < A[k+1]；
 *
 *
 * 或 若 i <= k < j ：
 *
 * 当 k 为偶数时，A[k] > A[k+1] ，且
 * 当 k 为奇数时， A[k] < A[k+1]。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [9,4,2,10,7,8,8,1,9]
 * 输出：5
 * 解释：arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
 *
 * 示例 2：
 *
 *
 * 输入：arr = [4,8,12,16]
 * 输出：2
 *
 *
 * 示例 3：
 *
 *
 * 输入：arr = [100]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 4 * 10^4
 * 0 <= arr[i] <= 10^9
 *
 *
 */

const MAXN int = 4e4

var dp [2]int

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	for i := 0; i < 2; i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		prev := [2]int{dp[0], dp[1]}
		dp[0], dp[1] = 1, 1
		if arr[i] > arr[i-1] {
			dp[0] = prev[1]+1
		}
		if arr[i] < arr[i-1] {
			dp[1] = prev[0]+1
		}
		ans = max(ans, max(dp[0], dp[1]))
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
