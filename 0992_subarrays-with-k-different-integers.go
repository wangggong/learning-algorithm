/*
 * @lc app=leetcode.cn id=subarrays-with-k-different-integers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [992] K 个不同整数的子数组
 *
 * https://leetcode-cn.com/problems/subarrays-with-k-different-integers/description/
 *
 * algorithms
 * Hard (45.91%)
 * Total Accepted:    25.5K
 * Total Submissions: 55.5K
 * Testcase Example:  '[1,2,1,2,3]\n2'
 *
 * 给定一个正整数数组 nums和一个整数 k ，返回 num 中 「好子数组」 的数目。
 *
 * 如果 nums 的某个子数组中不同整数的个数恰好为 k，则称 nums 的这个连续、不一定不同的子数组为 「好子数组 」。
 *
 *
 * 例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。
 *
 *
 * 子数组 是数组的 连续 部分。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,1,2,3], k = 2
 * 输出：7
 * 解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2],
 * [1,2,1,2].
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,1,3,4], k = 3
 * 输出：3
 * 解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= nums[i], k <= nums.length
 *
 *
 */

func subarraysWithKDistinct(nums []int, k int) int {
	// assert len(nums) > 0;
	// assert k > 0;
	return subarraysWithAtLeastK(nums, k) - subarraysWithAtLeastK(nums, k-1)
}

// 返回至少 `k` 个正数构成的子数组数目.
func subarraysWithAtLeastK(arr []int, k int) int {
	n := len(arr)
	p, q := 0, 0
	ans := 0
	count := make(map[int]int)
	for ; q < n; q++ {
		count[arr[q]]++
		// NOTE: 先把不满足条件的左指针越过, 再统计.
		// 否则会漏结果.
		for ; p < n && len(count) > k; p++ {
			count[arr[p]]--
			if count[arr[p]] == 0 {
				delete(count, arr[p])
			}
		}
		ans += q - p + 1
	}
	return ans
}
