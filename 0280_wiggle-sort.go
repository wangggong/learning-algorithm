/*
 * Medium
 *
 * 给你一个的整数数组 nums, 将该数组重新排序后使 nums[0] <= nums[1] >= nums[2] <= nums[3]...
 *
 * 输入数组总是有一个有效的答案。
 *
 *
 *
 * 示例 1:
 *
 * 输入：nums = [3,5,2,1,6,4]
 * 输出：[3,5,1,6,2,4]
 * 解释：[1,6,2,5,3,4]也是有效的答案
 * 示例 2:
 *
 * 输入：nums = [6,6,5,6,3,8]
 * 输出：[6,6,5,6,3,8]
 *
 *
 * 提示：
 *
 * 1 <= nums.length <= 5 * 104
 * 0 <= nums[i] <= 104
 * 输入的 nums 保证至少有一个答案。
 *
 *
 *
 * 进阶：你能在 O(n) 时间复杂度下解决这个问题吗？
 *
 * 通过次数6,677
 * 提交次数9,730
 */

func wiggleSort(nums []int) {
	n := len(nums)
	less := true
	for i := 0; i+1 < n; i++ {
		if less {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		} else {
			if nums[i] < nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		less = !less
	}
	return
}
