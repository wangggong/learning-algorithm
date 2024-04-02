/*
 * @lc app=leetcode.cn id=shortest-subarray-with-sum-at-least-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [862] 和至少为 K 的最短子数组
 *
 * https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/
 *
 * algorithms
 * Hard (22.19%)
 * Total Accepted:    38K
 * Total Submissions: 154K
 * Testcase Example:  '[1]\n1'
 *
 * 给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组
 * ，返回 -1 。
 *
 * 子数组 是数组中 连续 的一部分。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1], k = 1
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2], k = 4
 * 输出：-1
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2,-1,2], k = 3
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^5 <= nums[i] <= 10^5
 * 1 <= k <= 10^9
 *
 *
 */
func shortestSubarray(nums []int, k int) int {
	ans, cur := len(nums)+1, 0
	var Q [][2]int
	Q = append(Q, [2]int{0, 0})
	for i, n := range nums {
		cur += n
		for len(Q) > 0 && cur-Q[0][0] >= k {
			ans = min(ans, i+1-Q[0][1])
			Q = Q[1:]
		}
		for len(Q) > 0 && Q[len(Q)-1][0] >= cur {
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, [2]int{cur, i + 1})
	}
	if ans > len(nums) {
		return -1
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
