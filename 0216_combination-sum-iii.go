/*
 * @lc app=leetcode.cn id=combination-sum-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [216] 组合总和 III
 *
 * https://leetcode-cn.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (73.41%)
 * Total Accepted:    122K
 * Total Submissions: 166.2K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
 *
 * 说明：
 *
 *
 * 所有数字都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 *
 *
 * 示例 2:
 *
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */
var ans [][]int
var cur []int

const endValue = 10

func combinationSum3(k int, n int) [][]int {
	ans, cur = nil, nil
	// assert k > 0;
	dfs(k, 0, n, 1)
	return ans
}

func dfs(k, m, n, begin int) {
	if k == 0 {
		if m != n {
			return
		}
		v := make([]int, len(cur))
		copy(v, cur)
		ans = append(ans, v)
		return
	}
	if m+begin > n {
		return
	}
	for i := begin; i < endValue; i++ {
		cur = append(cur, i)
		dfs(k-1, m+i, n, i+1)
		cur = cur[:len(cur)-1]
	}
}
