/*
 * @lc app=leetcode.cn id=contiguous-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [525] 连续数组
 *
 * https://leetcode-cn.com/problems/contiguous-array/description/
 *
 * algorithms
 * Medium (54.07%)
 * Total Accepted:    53.6K
 * Total Submissions: 99.1K
 * Testcase Example:  '[0,1]'
 *
 * 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [0,1]
 * 输出: 2
 * 说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
 *
 * 示例 2:
 *
 *
 * 输入: nums = [0,1,0]
 * 输出: 2
 * 说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1e5
 * nums[i] 不是 0 就是 1
 *
 *
 */
func findMaxLength(nums []int) int {
	ans := 0
	cur := 0
	m := make(map[int]int)
	m[cur] = -1
	for i, n := range nums {
		if n == 0 {
			cur++
		} else {
			cur--
		}
		if v, ok := m[cur]; ok {
			ans = max(ans, i-v)
		} else {
			m[cur] = i
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
