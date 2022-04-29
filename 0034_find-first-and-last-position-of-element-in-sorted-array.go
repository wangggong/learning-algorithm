/*
 * @lc app=leetcode.cn id=find-first-and-last-position-of-element-in-sorted-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.22%)
 * Total Accepted:    502.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 *
 * 进阶：
 *
 *
 * 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 *
 * 示例 3：
 *
 *
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 1e5
 * -1e9 <= nums[i] <= 1e9
 * nums 是一个非递减数组
 * -1e9 <= target <= 1e9
 *
 *
 */
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || nums[0] > target || nums[n-1] < target {
		return []int{-1, -1}
	}
	var left, right int
	p, q := -1, n
	for p < q {
		mid := (p + q + 1) >> 1
		if nums[mid] < target {
			p = mid
		} else {
			q = mid - 1
		}
	}
	left = p
	if left+1 >= n || nums[left+1] != target {
		return []int{-1, -1}
	}
	p, q = 0, n
	for p < q {
		mid := (p + q) >> 1
		if nums[mid] > target {
			q = mid
		} else {
			p = mid + 1
		}
	}
	right = q
	if right-1 < 0 || nums[right-1] != target {
		return []int{-1, -1}
	}
	return []int{left + 1, right - 1}
}
