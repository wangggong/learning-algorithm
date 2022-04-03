/*
 * @lc app=leetcode.cn id=number-of-islands lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (56.74%)
 * Total Accepted:    396.2K
 * Total Submissions: 697.5K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 300
 * grid[i][j] 的值为 '0' 或 '1'
 *
 *
 */

const MAXN = 300 + 10

var count int
var m, n int
var G [MAXN][MAXN]bool

var directions = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func numIslands(grid [][]byte) int {
	n, m = len(grid), len(grid[0])
	// assert 0 < n && n <= 300 && 0 < m && m <= 300;
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				continue
			}
			dfs(grid, i, j)
			ans++
		}
	}
	return ans
}

func dfs(grid [][]byte, y, x int) {
	if x < 0 || x >= m || y < 0 || y >= n || grid[y][x] == '0' {
		return
	}
	grid[y][x] = '0'
	for _, d := range directions {
		dfs(grid, y+d[1], x+d[0])
	}
}
