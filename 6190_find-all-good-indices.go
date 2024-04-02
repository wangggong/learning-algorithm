/*
 * @lc app=leetcode.cn id=find-all-good-indices lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6190] 找到所有好下标
 *
 * https://leetcode.cn/problems/find-all-good-indices/description/
 *
 * algorithms
 * Medium (25.39%)
 * Total Accepted:    4.2K
 * Total Submissions: 16.7K
 * Testcase Example:  '[2,1,1,1,3,4,1]\n2'
 *
 * 给你一个大小为 n 下标从 0 开始的整数数组 nums 和一个正整数 k 。
 *
 * 对于 k <= i < n - k 之间的一个下标 i ，如果它满足以下条件，我们就称它为一个 好 下标：
 *
 *
 * 下标 i 之前 的 k 个元素是 非递增的 。
 * 下标 i 之后 的 k 个元素是 非递减的 。
 *
 *
 * 按 升序 返回所有好下标。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,1,1,1,3,4,1], k = 2
 * 输出：[2,3]
 * 解释：数组中有两个好下标：
 * - 下标 2 。子数组 [2,1] 是非递增的，子数组 [1,3] 是非递减的。
 * - 下标 3 。子数组 [1,1] 是非递增的，子数组 [3,4] 是非递减的。
 * 注意，下标 4 不是好下标，因为 [4,1] 不是非递减的。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,1,1,2], k = 2
 * 输出：[]
 * 解释：数组中没有好下标。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 3 <= n <= 10^5
 * 1 <= nums[i] <= 10^6
 * 1 <= k <= n / 2
 *
 *
 */
func goodIndices(nums []int, k int) (ans []int) {
	n := len(nums)
	ls, rs := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		ls[i], rs[i] = 1, 1
	}
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			ls[i] = ls[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			rs[i] = rs[i+1] + 1
		}
	}
	for i := 1; i+1 < n; i++ {
		if ls[i-1] >= k && rs[i+1] >= k {
			ans = append(ans, i)
		}
	}
	return
}
