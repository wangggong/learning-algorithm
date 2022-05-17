/*
 * @lc app=leetcode.cn id=move-zeroes lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (63.96%)
 * Total Accepted:    745.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [0]
 * 输出: [0]
 *
 *
 *
 * 提示:
 *
 *
 *
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你能尽量减少完成的操作次数吗？
 *
 */
func moveZeroes(nums []int) {
	cnt := 0
	for p, q, n := 0, 0, len(nums); q < n; q++ {
		if nums[q] == 0 {
			cnt++
			continue
		}
		nums[p] = nums[q]
		p++
	}
	for i := len(nums) - 1; cnt > 0; i, cnt = i-1, cnt-1 {
		nums[i] = 0
	}
	return
}
