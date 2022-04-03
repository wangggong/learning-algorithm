/*
 * @lc app=leetcode.cn id=trapping-rain-water lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (59.28%)
 * Total Accepted:    381K
 * Total Submissions: 642.7K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */
func trap(H []int) int {
	n := len(H)
	left, right := make([]int, n), make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		left[i] = H[i]
		if i > 0 && left[i-1] > left[i] {
			left[i] = left[i-1]
		}
	}
	for i := n - 1; i >= 0; i-- {
		right[i] = H[i]
		if i+1 < n && right[i+1] > right[i] {
			right[i] = right[i+1]
		}
	}
	for i := 0; i < n; i++ {
		h := left[i]
		if right[i] < h {
			h = right[i]
		}
		ans += h - H[i]
	}
	return ans
}

