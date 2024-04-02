/*
 * @lc app=leetcode.cn id=steps-to-make-array-non-decreasing lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6080] 使数组按非递减顺序排列
 *
 * https://leetcode.cn/problems/steps-to-make-array-non-decreasing/description/
 *
 * algorithms
 * Medium (10.44%)
 * Total Accepted:    1.3K
 * Total Submissions: 12.4K
 * Testcase Example:  '[5,3,4,4,7,3,6,11,8,5,11]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。在一步操作中，移除所有满足 nums[i - 1] > nums[i] 的 nums[i] ，其中 0
 * < i < nums.length 。
 *
 * 重复执行步骤，直到 nums 变为 非递减 数组，返回所需执行的操作数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,3,4,4,7,3,6,11,8,5,11]
 * 输出：3
 * 解释：执行下述几个步骤：
 * - 步骤 1 ：[5,3,4,4,7,3,6,11,8,5,11] 变为 [5,4,4,7,6,11,11]
 * - 步骤 2 ：[5,4,4,7,6,11,11] 变为 [5,4,7,11,11]
 * - 步骤 3 ：[5,4,7,11,11] 变为 [5,7,11,11]
 * [5,7,11,11] 是一个非递减数组，因此，返回 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,7,7,13]
 * 输出：0
 * 解释：nums 已经是一个非递减数组，因此，返回 0 。
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
func totalSteps(nums []int) int {
	n := len(nums)
	ans := 0
	dp := make([]int, n)
	var S []int
	S = append(S, 0)
	for i := 1; i < n; i++ {
		cur := 0
		for len(S) > 0 && nums[S[len(S)-1]] <= nums[i] {
			cur = max(cur, dp[S[len(S)-1]])
			S = S[:len(S)-1]
		}
		if len(S) > 0 {
			dp[i] = cur + 1
			ans = max(ans, dp[i])
		}
		S = append(S, i)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
