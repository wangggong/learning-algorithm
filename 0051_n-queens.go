/*
 * @lc app=leetcode.cn id=n-queens lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [51] N 皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (73.82%)
 * Total Accepted:    184K
 * Total Submissions: 249.3K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 *
 *
 *
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[["Q"]]
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

const MAXN = 15

var G []int
var ans [][]string

func solveNQueens(n int) [][]string {
	ans, G = nil, nil
	backtrack(n, 0)
	return ans
}

func backtrack(n, k int) {
	// fmt.Printf("%v %v %v\n", n, k, G)
	if k == n {
		var tmp []string
		for _, g := range G {
			tmp = append(tmp, toChessRow(n, g))
		}
		ans = append(ans, tmp)
		return
	}
	for v := 0; v < n; v++ {
		valid := true
		for j, g := range G {
			if v == g {
				valid = false
				break
			}
			if (v-g == k-j) || (v-g == j-k) {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		G = append(G, v)
		backtrack(n, k+1)
		G = G[:len(G)-1]
	}
}

func toChessRow(n, k int) string {
	var ans []byte
	for i := 0; i < n; i++ {
		ans = append(ans, '.')
	}
	ans[k] = 'Q'
	return string(ans)
}
