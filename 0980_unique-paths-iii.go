/*
 * @lc app=leetcode.cn id=unique-paths-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [980] 不同路径 III
 *
 * https://leetcode-cn.com/problems/unique-paths-iii/description/
 *
 * algorithms
 * Hard (73.65%)
 * Total Accepted:    15.6K
 * Total Submissions: 21.1K
 * Testcase Example:  '[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]'
 *
 * 在二维网格 grid 上，有 4 种类型的方格：
 *
 *
 * 1 表示起始方格。且只有一个起始方格。
 * 2 表示结束方格，且只有一个结束方格。
 * 0 表示我们可以走过的空方格。
 * -1 表示我们无法跨越的障碍。
 *
 *
 * 返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。
 *
 * 每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
 * 输出：2
 * 解释：我们有以下两条路径：
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
 * 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
 *
 * 示例 2：
 *
 * 输入：[[1,0,0,0],[0,0,0,0],[0,0,0,2]]
 * 输出：4
 * 解释：我们有以下四条路径：
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
 * 2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
 * 3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
 * 4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
 *
 * 示例 3：
 *
 * 输入：[[0,1],[2,0]]
 * 输出：0
 * 解释：
 * 没有一条路能完全穿过每一个空的方格一次。
 * 请注意，起始和结束方格可以位于网格中的任意位置。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= grid.length * grid[0].length <= 20
 *
 *
 */
const MAXN = 20 + 5

var n, m int
var visit [MAXN][MAXN]bool
var directions = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func dfs(grid [][]int, cnt, srcX, srcY, dstX, dstY int, ans *int) {
	if cnt < 0 {
		if srcX == dstX && srcY == dstY {
			*ans++
		}
		// fmt.Printf("%v %v %v %v %v %v\n", srcX, srcY, dstX, dstY, cnt, *ans)
		return
	}
	for _, d := range directions {
		nx, ny := srcX+d[1], srcY+d[0]
		if nx < 0 || nx >= m || ny < 0 || ny >= n || visit[ny][nx] || grid[ny][nx] == -1 {
			continue
		}
		visit[ny][nx] = true
		dfs(grid, cnt-1, nx, ny, dstX, dstY, ans)
		visit[ny][nx] = false
	}
}

func uniquePathsIII(grid [][]int) int {
	n, m = len(grid), len(grid[0])
	// assert n > 0 && n <= 20;
	// assert m > 0 && m <= 20;
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			visit[i][j] = false
		}
	}
	ans, cnt := 0, 0
	srcX, srcY, dstX, dstY := -1, -1, -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch grid[i][j] {
			case 0:
				cnt++
			case 1:
				srcX, srcY = j, i
			case 2:
				dstX, dstY = j, i
			}
		}
	}
	visit[srcY][srcX] = true
	dfs(grid, cnt, srcX, srcY, dstX, dstY, &ans)
	return ans
}
