/*
 * @lc app=leetcode.cn id=subsets lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (80.35%)
 * Total Accepted:    397.9K
 * Total Submissions: 495K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 * 
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
 * nums 中的所有元素 互不相同
 * 
 * 
 */

func subsets(nums []int) [][]int {
	n := len(nums)
	lim := 1 << n
	var ans [][]int
	for i := 0; i < lim; i++ {
		ans = append(ans, getSubsetFromVal(nums, i))
	}
	return ans
}

func getSubsetFromVal(arr []int, k int) []int {
	var ans []int
	n := len(arr)
	for i := 0; i < n; i++ {
		if k & (1 << i) != 0 {
			ans = append(ans, arr[i])
		}
	}
	return ans
}
