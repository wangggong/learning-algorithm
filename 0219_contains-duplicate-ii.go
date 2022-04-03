/*
 * @lc app=leetcode.cn id=contains-duplicate-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [219] 存在重复元素 II
 *
 * https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (44.06%)
 * Total Accepted:    160.7K
 * Total Submissions: 361.3K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且
 * abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,1], k = 3
 * 输出：true
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,0,1,1], k = 1
 * 输出：true
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,2,3,1,2,3], k = 2
 * 输出：false
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 0 <= k <= 10^5
 *
 *
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		if m[nums[i]] > 0 {
			return true
		}
		m[nums[i]]++
		if i-k >= 0 {
			m[nums[i-k]]--
		}
	}
	return false
}
