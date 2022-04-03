/*
 * @lc app=leetcode.cn id=number-of-enclaves lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1020] 飞地的数量
 *
 * https://leetcode-cn.com/problems/number-of-enclaves/description/
 *
 * algorithms
 * Medium (56.40%)
 * Total Accepted:    19.9K
 * Total Submissions: 33.4K
 * Testcase Example:  '[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]'
 *
 * 给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
 *
 * 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
 *
 * 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
 * 输出：3
 * 解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
 * 输出：0
 * 解释：所有 1 都在边界上或可以到达边界。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 500
 * grid[i][j] 的值为 0 或 1
 *
 *
 */

type Pair struct{ x, y int }

var directions = [4]Pair{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func numEnclaves(grid [][]int) int {
	// assert len(grid) > 1 && len(grid) <= 500;
	// assert len(grid[0]) > 1 && len(grid[0]) <= 500;
	n, m := len(grid), len(grid[0])
	var Q []Pair
	for i := 0; i < n; i++ {
		if grid[i][0] == 1 {
			Q = append(Q, Pair{i, 0})
			grid[i][0] = 0
		}
		if grid[i][m-1] == 1 {
			Q = append(Q, Pair{i, m - 1})
			grid[i][m-1] = 0
		}
	}
	for i := 0; i < m; i++ {
		if grid[0][i] == 1 {
			Q = append(Q, Pair{0, i})
			grid[0][i] = 0
		}
		if grid[n-1][i] == 1 {
			Q = append(Q, Pair{n - 1, i})
			grid[n-1][i] = 0
		}
	}
	for len(Q) > 0 {
		q := Q[len(Q)-1]
		Q = Q[:len(Q)-1]
		for _, d := range directions {
			nx, ny := q.x+d.x, q.y+d.y
			if nx < 0 || nx >= n || ny < 0 || ny >= m || grid[nx][ny] == 0 {
				continue
			}
			Q = append(Q, Pair{nx, ny})
			grid[nx][ny] = 0
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}
