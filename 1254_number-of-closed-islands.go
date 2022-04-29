/*
 * @lc app=leetcode.cn id=number-of-closed-islands lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1254] 统计封闭岛屿的数目
 *
 * https://leetcode-cn.com/problems/number-of-closed-islands/description/
 *
 * algorithms
 * Medium (61.50%)
 * Total Accepted:    24.2K
 * Total Submissions: 39.3K
 * Testcase Example:  '[[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]'
 *
 * 二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
 *
 * 请返回 封闭岛屿 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid =
 * [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
 * 输出：2
 * 解释：
 * 灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1,1,1,1,1,1,1],
 * [1,0,0,0,0,0,1],
 * [1,0,1,1,1,0,1],
 * [1,0,1,0,1,0,1],
 * [1,0,1,1,1,0,1],
 * [1,0,0,0,0,0,1],
 * ⁠            [1,1,1,1,1,1,1]]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= grid.length, grid[0].length <= 100
 * 0 <= grid[i][j] <=1
 *
 *
 */
var n, m int
var directions = [5]int{1, 0, -1, 0, 1}

func closedIsland(G [][]int) int {
	n, m = len(G), len(G[0])
	for i := 0; i < n; i++ {
		if G[i][0] == 0 {
			dfs(G, i, 0)
		}
		if G[i][m-1] == 0 {
			dfs(G, i, m-1)
		}
	}
	for i := 0; i < m; i++ {
		if G[0][i] == 0 {
			dfs(G, 0, i)
		}
		if G[n-1][i] == 0 {
			dfs(G, n-1, i)
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == 0 {
				dfs(G, i, j)
				ans++
			}
		}
	}
	return ans
}

func dfs(G [][]int, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || G[x][y] == 1 {
		return
	}
	G[x][y] = 1
	for i := 0; i < 4; i++ {
		dfs(G, x+directions[i], y+directions[i+1])
	}
	return
}
