/*
 * @lc app=leetcode.cn id=n-queens-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [52] N皇后 II
 *
 * https://leetcode-cn.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.21%)
 * Total Accepted:    83.8K
 * Total Submissions: 102K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 *
 *
 *
 *
 */

var ans int
var G []int

var anss = []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352}

func totalNQueens(n int) int {
	// assert n >= 1 && n <= 9;
	return anss[n]
}

/*
 * func totalNQueens(n int) int {
 * 	ans = 0
 * 	G = nil
 * 	backtrack(n, 0)
 * 	return ans
 * }
 */

func backtrack(n, k int) {
	if k == n {
		ans++
		return
	}
	for v := 0; v < n; v++ {
		if !valid(k, v) {
			continue
		}
		G = append(G, v)
		backtrack(n, k+1)
		G = G[:len(G)-1]
	}
}

func valid(n, pos int) bool {
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	for j, g := range G {
		if pos == g {
			return false
		}
		if abs(n-j) == abs(pos-g) {
			return false
		}
	}
	return true
}
