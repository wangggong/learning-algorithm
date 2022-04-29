/*
 * @lc app=leetcode.cn id=first-missing-positive lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (42.42%)
 * Total Accepted:    211.6K
 * Total Submissions: 498.6K
 * Testcase Example:  '[1,2,0]'
 *
 * 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
 * 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,0]
 * 输出：3
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,4,-1,1]
 * 输出：2
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [7,8,9,11,12]
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 5 * 1e5
 * -1 << 31 <= nums[i] <= 1 << 31 - 1
 * 
 * 
 */
func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			n--
			i--
		}
	}
	nums = nums[:n]
	for i := 0; i < n; i++ {
		v := abs(nums[i])
		if v-1 < n {
			nums[v-1] = abs(nums[v-1]) * -1
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			return i+1
		}
	}
	return n+1
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
