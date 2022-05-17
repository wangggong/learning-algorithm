/*
 * @lc app=leetcode.cn id=surrounded-regions lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [130] 被围绕的区域
 *
 * https://leetcode-cn.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (45.54%)
 * Total Accepted:    174.6K
 * Total Submissions: 383.4K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X'
 * 填充。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
 * 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
 * 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["X"]]
 * 输出：[["X"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n == board[i].length
 * 1 <= m, n <= 200
 * board[i][j] 为 'X' 或 'O'
 *
 *
 *
 *
 */
var directions = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func solve(B [][]byte) {
	n, m := len(B), len(B[0])
	var Q [][2]int
	for i := 0; i < n; i++ {
		if B[i][0] == 'O' {
			Q = append(Q, [2]int{i, 0})
			B[i][0] = 'M'
		}
		if B[i][m-1] == 'O' {
			Q = append(Q, [2]int{i, m - 1})
			B[i][m-1] = 'M'
		}
	}
	for i := 0; i < m; i++ {
		if B[0][i] == 'O' {
			Q = append(Q, [2]int{0, i})
			B[0][i] = 'M'
		}
		if B[n-1][i] == 'O' {
			Q = append(Q, [2]int{n - 1, i})
			B[n-1][i] = 'M'
		}
	}
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		for _, d := range directions {
			nx, ny := q[0]+d[0], q[1]+d[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= m || B[nx][ny] != 'O' {
				continue
			}
			Q = append(Q, [2]int{nx, ny})
			B[nx][ny] = 'M'
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if B[i][j] == 'O' {
				B[i][j] = 'X'
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if B[i][j] == 'M' {
				B[i][j] = 'O'
			}
		}
	}
	return
}
