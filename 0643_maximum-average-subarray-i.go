/*
 * @lc app=leetcode.cn id=maximum-average-subarray-i lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [643] 子数组最大平均数 I
 *
 * https://leetcode-cn.com/problems/maximum-average-subarray-i/description/
 *
 * algorithms
 * Easy (44.45%)
 * Total Accepted:    78.3K
 * Total Submissions: 176.1K
 * Testcase Example:  '[1,12,-5,-6,50,3]\n4'
 *
 * 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
 *
 * 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
 *
 * 任何误差小于 10^-5 的答案都将被视为正确答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,12,-5,-6,50,3], k = 4
 * 输出：12.75
 * 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5], k = 1
 * 输出：5.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= k <= n <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */
func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	ps := make([]int, n+1)
	for i, v := range nums {
		ps[i+1] = ps[i] + v
	}
	ans := float64(-1e4 - 10)
	for i := k; i <= n; i++ {
		ans = max(ans, float64(ps[i]-ps[i-k])/float64(k))
	}
	return ans
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
