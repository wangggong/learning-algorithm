/*
 * @lc app=leetcode.cn id=next-greater-element-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [503] 下一个更大元素 II
 *
 * https://leetcode-cn.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (64.76%)
 * Total Accepted:    128.5K
 * Total Submissions: 198.4K
 * Testcase Example:  '[1,2,1]'
 *
 * 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的
 * 下一个更大元素 。
 *
 * 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
 * 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,1]
 * 输出: [2,-1,2]
 * 解释: 第一个 1 的下一个更大的数是 2；
 * 数字 2 找不到下一个更大的数；
 * 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4,3]
 * 输出: [2,3,4,-1,4]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

/*
 * func nextGreaterElements(nums []int) []int {
 * 	n := len(nums)
 * 	ans := make([]int, n)
 * 	for i := range ans {
 * 		ans[i] = -1
 * 	}
 * 	var cur []int
 * 	cur = append(cur, nums...)
 * 	cur = append(cur, nums...)
 * 	var S []int
 * 	for i := 2*n - 1; i >= 0; i-- {
 * 		for len(S) > 0 && S[len(S)-1] <= cur[i] {
 * 			S = S[:len(S)-1]
 * 		}
 * 		if len(S) == 0 {
 * 			ans[i%n] = -1
 * 		} else {
 * 			ans[i%n] = S[len(S)-1]
 * 		}
 * 		S = append(S, cur[i])
 * 	}
 * 	return ans
 * }
 */

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	var S []int
	for i := 2*n-1; i >= 0; i-- {
		ind := i%n
		for len(S) > 0 && S[len(S)-1] <= nums[ind] {
			S = S[:len(S)-1]
		}
		if len(S) == 0 {
			ans[ind] = -1
		} else {
			ans[ind] = S[len(S)-1]
		}
		S = append(S, nums[ind])
	}
	return ans
}
