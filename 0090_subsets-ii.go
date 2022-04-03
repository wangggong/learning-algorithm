/*
 * @lc app=leetcode.cn id=subsets-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (63.39%)
 * Total Accepted:    177.4K
 * Total Submissions: 279.7K
 * Testcase Example:  '[1,2,2]'
 *
 * 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,2]
 * 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10
 * -10 <= nums[i] <= 10
 *
 *
 *
 *
 */

import "sort"

var ans [][]int
var cur []int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans, cur = nil, nil
	backtrack(nums, 0)
	return ans
}

func backtrack(arr []int, k int) {
	n := len(arr)
	if n == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ans = append(ans, tmp)
		return
	}
	// 先把当前结果放进结果集啊傻逼!
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	ans = append(ans, tmp)
	for i := k; i < n; i++ {
		// 然后直接去重就好了啊!
		if i > k && arr[i] == arr[i-1] {
			continue
		}
		cur = append(cur, arr[i])
		backtrack(arr, i+1)
		cur = cur[:len(cur)-1]
	}
}
