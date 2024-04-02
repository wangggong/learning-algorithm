/*
 * @lc app=leetcode.cn id=check-if-there-is-a-valid-partition-for-the-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6137] 检查数组是否存在有效划分
 *
 * https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/description/
 *
 * algorithms
 * Medium (32.85%)
 * Total Accepted:    4.5K
 * Total Submissions: 13.6K
 * Testcase Example:  '[4,4,4,5,6]'
 *
 * 给你一个下标从 0 开始的整数数组 nums ，你必须将数组划分为一个或多个 连续 子数组。
 *
 * 如果获得的这些子数组中每个都能满足下述条件 之一 ，则可以称其为数组的一种 有效 划分：
 *
 *
 * 子数组 恰 由 2 个相等元素组成，例如，子数组 [2,2] 。
 * 子数组 恰 由 3 个相等元素组成，例如，子数组 [4,4,4] 。
 * 子数组 恰 由 3 个连续递增元素组成，并且相邻元素之间的差值为 1 。例如，子数组 [3,4,5] ，但是子数组 [1,3,5] 不符合要求。
 *
 *
 * 如果数组 至少 存在一种有效划分，返回 true ，否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,4,4,5,6]
 * 输出：true
 * 解释：数组可以划分成子数组 [4,4] 和 [4,5,6] 。
 * 这是一种有效划分，所以返回 true 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,1,1,2]
 * 输出：false
 * 解释：该数组不存在有效划分。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^6
 *
 *
 */
func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			dp[i+1] = dp[i+1] || dp[i-1]
		}
		if i > 1 && nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			dp[i+1] = dp[i+1] || dp[i-2]
		}
		if i > 1 && nums[i] == nums[i-1]+1 && nums[i] == nums[i-2]+2 {
			dp[i+1] = dp[i+1] || dp[i-2]
		}
	}
	return dp[n]
}
