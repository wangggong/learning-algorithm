/*
 * @lc app=leetcode.cn id=combinations lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (76.98%)
 * Total Accepted:    278.7K
 * Total Submissions: 362.2K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
 *
 * 你可以按 任何顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, k = 2
 * 输出：
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 * 示例 2：
 *
 *
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 20
 * 1 <= k <= n
 *
 *
 */
const MAXN = 20 + 10

var ans [][]int
var cur []int

func combine(n int, k int) [][]int {
	ans, cur = nil, nil
	dfs(1, n+1, k)
	return ans
}

func dfs(m, n, k int) {
	if m > n {
		return
	}
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ans = append(ans, tmp)
		return
	}
	for i := m; i <= n; i++ {
		cur = append(cur, i)
		dfs(i+1, n, k-1)
		cur = cur[:len(cur)-1]
	}
}
