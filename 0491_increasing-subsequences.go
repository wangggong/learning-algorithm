/*
 * @lc app=leetcode.cn id=increasing-subsequences lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [491] 递增子序列
 *
 * https://leetcode-cn.com/problems/increasing-subsequences/description/
 *
 * algorithms
 * Medium (53.79%)
 * Total Accepted:    64.9K
 * Total Submissions: 121.6K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
 *
 * 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,6,7,7]
 * 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,4,3,2,1]
 * 输出：[[4,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 15
 * -100 <= nums[i] <= 100
 *
 *
 */

var ans [][]int
var cur []int

func findSubsequences(nums []int) [][]int {
	ans, cur = nil, nil
	backtrack(nums, 0)
	return ans
}

func backtrack(arr []int, k int) {
	n := len(arr)
	if k == n {
		if len(cur) < 2 {
			return
		}
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ans = append(ans, tmp)
		return
	}
	if len(cur) >= 2 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ans = append(ans, tmp)
	}
	// used := make(map[int]struct{})
	used := make([]bool, 200 + 10)
	for i := k; i < n; i++ {
		if len(cur) > 0 && cur[len(cur)-1] > arr[i] {
			continue
		}
		// if i > k && arr[i] == arr[i-1] {
		// if _, ok := used[arr[i]]; ok {
		if used[arr[i] + 100] {
			continue
		}
		// used[arr[i]] = struct{}{}
		used[arr[i] + 100] = true
		cur = append(cur, arr[i])
		backtrack(arr, i+1)
		cur = cur[:len(cur)-1]
	}
}
