/*
 * @lc app=leetcode.cn id=search-in-rotated-sorted-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (43.19%)
 * Total Accepted:    457.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 整数数组 nums 按升序排列，数组中的值 互不相同 。
 *
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 ）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ...,
 * nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如，
 * [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
 *
 * 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1], target = 0
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * nums 中的每个值都 独一无二
 * 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 * -10^4 <= target <= 10^4
 *
 *
 *
 *
 * 进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？
 *
 */

// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环.
func search(nums []int, target int) int {
	n := len(nums)
	p, q := 0, n-1
	for p <= q {
		mid := (p + q) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[p] < nums[q] {
			// 整个数组有序.
			return bs(nums, p, q, target)
		}
		if nums[p] <= nums[mid] {
			// 左半边数组有序.
			if target > nums[mid] || target < nums[p] {
				p = mid + 1
			} else {
				return bs(nums, p, mid, target)
			}
		} else {
			// 右半边数组有序.
			if target > nums[q] || target < nums[mid] {
				q = mid - 1
			} else {
				return bs(nums, mid, q, target)
			}
		}
	}
	if p < len(nums) && nums[p] == target {
		return p
	}
	return -1
}

func bs(arr []int, p, q, target int) int {
	for p <= q {
		mid := (p + q) >> 1
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			q = mid - 1
			continue
		}
		p = mid + 1
	}
	if p < len(arr) && arr[p] == target {
		return p
	}
	return -1
}
