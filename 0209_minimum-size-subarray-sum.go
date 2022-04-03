/*
 * @lc app=leetcode.cn id=minimum-size-subarray-sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (48.12%)
 * Total Accepted:    264.4K
 * Total Submissions: 548.5K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 target 。
 *
 * 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr]
 * ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：target = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：target = 4, nums = [1,4,4]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= target <= 1e9
 * 1 <= nums.length <= 1e5
 * 1 <= nums[i] <= 1e5
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
 *
 *
 */

/*
 * func minSubArrayLen(target int, nums []int) int {
 * 	n := len(nums)
 * 	prefixSum := make([]int, n+1)
 * 	for i := 0; i < n; i++ {
 * 		prefixSum[i+1] = prefixSum[i] + nums[i]
 * 	}
 * 	ans := n + 1
 * 	for i := 0; i < n; i++ {
 * 		// 二分
 * 		// [6]	2	3	1	2	4	3
 * 		// 0	2	5	6	8	12	15
 * 		// 0	0	0	0	1	1	1
 * 		p, q := i+1, n+1
 * 		for p < q {
 * 			mid := (p + q) >> 1
 * 			if prefixSum[mid]-prefixSum[i] >= target {
 * 				q = mid
 * 			} else {
 * 				p = mid + 1
 * 			}
 * 		}
 * 		if p > n {
 * 			continue
 * 		}
 * 		if ans > p-i {
 * 			ans = p - i
 * 		}
 * 	}
 * 	if ans == n+1 {
 * 		return 0
 * 	}
 * 	return ans
 * }
 */

/*
 * func minSubArrayLen(target int, nums []int) int {
 * 	n := len(nums)
 * 	prefixSum := make([]int, n+1)
 * 	for i := 0; i < n; i++ {
 * 		prefixSum[i+1] = prefixSum[i] + nums[i]
 * 	}
 * 	ans := n + 1
 * 	for p, q := 0, 0; p < n; p++ {
 * 		for q < n && prefixSum[q]-prefixSum[p] < target {
 * 			q++
 * 		}
 * 		if prefixSum[q]-prefixSum[p] >= target && ans > q-p {
 * 			ans = q - p
 * 		}
 * 	}
 * 	if ans == n+1 {
 * 		return 0
 * 	}
 * 	return ans
 * }
 */

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	p, q := 0, 0
	cur := 0
	ans := n+1
	for q < n {
		cur += nums[q]
		for cur >= target {
			ans = min(ans, q-p+1)
			cur -= nums[p]
			p++
		}
		q++
	}
	if ans > n {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
