/*
 * @lc app=leetcode.cn id=permutations lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (78.48%)
 * Total Accepted:    517.9K
 * Total Submissions: 659.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */

var ans [][]int
var n int

func permute(nums []int) [][]int {
	ans = nil
	n = len(nums)
	backtrack(nums, 0)
	return ans
}

func backtrack(arr []int, k int) {
	// NOTE: 每进一次 backtrack 要决定的是第 `k` 个元素是什么, 而不是第 `i` 个元素是什么.
	if k == n {
		tmp := make([]int, n)
		copy(tmp, arr)
		ans = append(ans, tmp)
		return
	}
	for i := k; i < n; i++ {
		arr[i], arr[k] = arr[k], arr[i]
		backtrack(arr, k+1)
		arr[i], arr[k] = arr[k], arr[i]
	}
}
