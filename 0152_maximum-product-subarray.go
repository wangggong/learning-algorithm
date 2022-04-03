/*
 * @lc app=leetcode.cn id=maximum-product-subarray lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (42.41%)
 * Total Accepted:    233.8K
 * Total Submissions: 550.9K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 * 测试用例的答案是一个 32-位 整数。
 *
 * 子数组 是数组的连续子序列。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
 *
 *
 */

/*
 * const MAXN int = 2e4
 *
 * var maxVal [MAXN + 5]int
 * var minVal [MAXN + 5]int
 *
 * func maxProduct(nums []int) int {
 * 	n := len(nums)
 * 	for i := 0; i < MAXN+5; i++ {
 * 		maxVal[i] = -MAXN
 * 		minVal[i] = MAXN
 * 	}
 * 	ans := nums[0]
 * 	maxVal[0], minVal[0] = nums[0], nums[0]
 * 	for i := 1; i < n; i++ {
 * 		maxVal[i] = max(nums[i], max(maxVal[i-1]*nums[i], minVal[i-1]*nums[i]))
 * 		minVal[i] = min(nums[i], min(maxVal[i-1]*nums[i], minVal[i-1]*nums[i]))
 * 		ans = max(ans, maxVal[i])
 * 	}
 * 	return ans
 * }
 */

func maxProduct(nums []int) int {
	n := len(nums)
	ans := nums[0]
	maxVal, minVal := nums[0], nums[0]
	for i := 1; i < n; i++ {
		prevMax, prevMin := maxVal, minVal
		maxVal = max(nums[i], max(prevMax*nums[i], prevMin*nums[i]))
		minVal = min(nums[i], min(prevMax*nums[i], prevMin*nums[i]))
		ans = max(ans, maxVal)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
