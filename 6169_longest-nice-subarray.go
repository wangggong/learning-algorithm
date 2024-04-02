/*
 * @lc app=leetcode.cn id=longest-nice-subarray lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6169] 最长优雅子数组
 *
 * https://leetcode.cn/problems/longest-nice-subarray/description/
 *
 * algorithms
 * Medium (37.91%)
 * Total Accepted:    3.8K
 * Total Submissions: 9.7K
 * Testcase Example:  '[1,3,8,48,10]'
 *
 * 给你一个由 正 整数组成的数组 nums 。
 *
 * 如果 nums 的子数组中位于 不同 位置的每对元素按位 与（AND）运算的结果等于 0 ，则称该子数组为 优雅 子数组。
 *
 * 返回 最长 的优雅子数组的长度。
 *
 * 子数组 是数组中的一个 连续 部分。
 *
 * 注意：长度为 1 的子数组始终视作优雅子数组。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,3,8,48,10]
 * 输出：3
 * 解释：最长的优雅子数组是 [3,8,48] 。子数组满足题目条件：
 * - 3 AND 8 = 0
 * - 3 AND 48 = 0
 * - 8 AND 48 = 0
 * 可以证明不存在更长的优雅子数组，所以返回 3 。
 *
 * 示例 2：
 *
 * 输入：nums = [3,1,5,11,13]
 * 输出：1
 * 解释：最长的优雅子数组长度为 1 ，任何长度为 1 的子数组都满足题目条件。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 *
 *
 */

// 参考: https://leetcode.cn/problems/longest-nice-subarray/solution/bao-li-mei-ju-pythonjavacgo-by-endlessch-z6t6/
const digit int = 32

func longestNiceSubarray(nums []int) (ans int) {
	for i, n := 0, len(nums); i < n; i++ {
		c, j := 0, 0
		for j = i; j < n && j < i+digit; j++ {
			if c&nums[j] != 0 {
				break
			}
			c = c | nums[j]
		}
		ans = max(ans, j-i)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
