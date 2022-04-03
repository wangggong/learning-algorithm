/*
 * @lc app=leetcode.cn id=single-element-in-a-sorted-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [540] 有序数组中的单一元素
 *
 * https://leetcode-cn.com/problems/single-element-in-a-sorted-array/description/
 *
 * algorithms
 * Medium (58.46%)
 * Total Accepted:    62.4K
 * Total Submissions: 103.1K
 * Testcase Example:  '[1,1,2,3,3,4,4,8,8]'
 *
 * 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
 *
 * 请你找出并返回只出现一次的那个数。
 *
 * 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,2,3,3,4,4,8,8]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums =  [3,3,7,7,10,11,11]
 * 输出: 10
 *
 *
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^5
 *
 *
 */

/*
 * func singleNonDuplicate(nums []int) int {
 * 	n := len(nums)
 * 	// assert n > 0 && n < 1e5;
 * 	// assert that count of number is 2 except the number is the given one.
 * 	if n == 1 {
 * 		return nums[0]
 * 	}
 * 	if nums[0] != nums[1] {
 * 		return nums[0]
 * 	}
 * 	if nums[n-1] != nums[n-2] {
 * 		return nums[n-1]
 * 	}
 * 	p, q := 0, n
 * 	for p < q {
 * 		mid := (p + q) >> 1
 * 		hasSame := false
 * 		if mid > 0 {
 * 			hasSame = hasSame || (nums[mid] == nums[mid-1])
 * 		}
 * 		if mid < n-1 {
 * 			hasSame = hasSame || (nums[mid] == nums[mid+1])
 * 		}
 * 		if !hasSame {
 * 			return nums[mid]
 * 		}
 * 		if (mid%2 == 1 && nums[mid] == nums[mid-1]) || (mid%2 == 0 && (mid+1 == n || nums[mid] == nums[mid+1])) {
 * 			p = mid + 1
 * 		} else {
 * 			q = mid
 * 		}
 * 	}
 * 	return nums[p]
 * }
 */

/*
 * func singleNonDuplicate(nums []int) int {
 * 	ans := 0
 * 	for _, n := range nums {
 * 		ans = ans ^ n
 * 	}
 * 	return ans
 * }
 */

func singleNonDuplicate(nums []int) int {
	p, q := 0, len(nums)-1
	for p <= q {
		mid := p+(q-p) / 2
		if mid == p || mid == q {
			return nums[mid]
		}
		if (mid % 2 == 0 && nums[mid] == nums[mid+1]) || (mid % 2 == 1 && nums[mid] == nums[mid-1]) {
			p = mid+mid % 2
		} else {
			q = mid-mid % 2
		}
	}
	return -1
}
