/*
 * @lc app=leetcode.cn id=subarray-sums-divisible-by-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [974] 和可被 K 整除的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/description/
 *
 * algorithms
 * Medium (47.06%)
 * Total Accepted:    44.4K
 * Total Submissions: 94.4K
 * Testcase Example:  '[4,5,0,-2,-3,1]\n5'
 *
 * 给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
 *
 * 子数组 是数组的 连续 部分。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,5,0,-2,-3,1], k = 5
 * 输出：7
 * 解释：
 * 有 7 个子数组满足其元素之和可被 k = 5 整除：
 * [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2,
 * -3]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [5], k = 9
 * 输出: 0
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -10^4 <= nums[i] <= 10^4
 * 2 <= k <= 10^4
 *
 *
 */
func subarraysDivByK(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i, n := range nums {
		preSum[i+1] = preSum[i] + ((n%k)+k)%k
	}
	ans := 0
	cnt := make([]int, k)
	cnt[0] = 1
	for i := 1; i <= n; i++ {
		ans += cnt[preSum[i]%k]
		cnt[preSum[i]%k]++
	}
	return ans
}
