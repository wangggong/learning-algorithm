/*
 * @lc app=leetcode.cn id=permutations-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (64.28%)
 * Total Accepted:    265.1K
 * Total Submissions: 412.4K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */

/*
 * import "sort"
 * const MAXN = 10+5
 *
 * var ans [][]int
 * var cur []int
 * var visit [MAXN]bool
 * var n int
 *
 * func permuteUnique(nums []int) [][]int {
 * 	sort.Ints(nums)
 * 	ans, cur = nil, nil
 * 	n = len(nums)
 * 	for i := range visit {
 * 		visit[i] = false
 * 	}
 * 	backtrack(nums, 0)
 * 	return ans
 * }
 *
 * func backtrack(arr []int, k int) {
 * 	if k == n {
 * 		tmp := make([]int, n)
 * 		copy(tmp, cur)
 * 		ans = append(ans, tmp)
 * 		return
 * 	}
 * 	// NOTE: 不重复的关键是需要跳过 "第一个 `1` 还没用先用上第二个 `1` 了" 这种情况.
 * 	// 在 4sum 中已经用到了这个技巧.
 * 	for i := 0; i < n; i++ {
 * 		if visit[i] {
 * 			continue
 * 		}
 * 		if i > 0 && arr[i] == arr[i-1] && !visit[i-1] {
 * 			continue
 * 		}
 * 		visit[i] = true
 * 		cur = append(cur, arr[i])
 * 		backtrack(arr, k+1)
 * 		visit[i] = false
 * 		cur = cur[:len(cur)-1]
 * 	}
 * }
 */

var ans [][]int

func permuteUnique(nums []int) [][]int {
	ans = nil
	backtrack(nums, 0, len(nums))
	return ans
}

func backtrack(arr []int, k, n int) {
	if k == n {
		ans = append(ans, copyOf(arr))
		return
	}
	// 傻逼! 这个地方不能用排序去重的技巧了! 数组在此处可能是无序的!
	visit := make(map[int]struct{})
	for i := k; i < n; i++ {
		if _, ok := visit[arr[i]]; ok {
			continue
		}
		visit[arr[i]] = struct{}{}
		arr[i], arr[k] = arr[k], arr[i]
		backtrack(arr, k+1, n)
		arr[i], arr[k] = arr[k], arr[i]
	}
	return
}

func copyOf(arr []int) []int {
	ans := make([]int, len(arr))
	copy(ans, arr)
	return ans
}