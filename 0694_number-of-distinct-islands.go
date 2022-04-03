/*
 * Medium
 * 给定一个非空 01 二维数组表示的网格，一个岛屿由四连通（上、下、左、右四个方向）的 1 组成，你可以认为网格的四周被海水包围。
 *
 * 请你计算这个网格中共有多少个形状不同的岛屿。两个岛屿被认为是相同的，当且仅当一个岛屿可以通过平移变换（不可以旋转、翻转）和另一个岛屿重合。
 *
 *
 *
 * 示例 1：
 *
 * 11000
 * 11000
 * 00011
 * 00011
 * 给定上图，返回结果 1 。
 *
 * 示例 2：
 *
 * 11011
 * 10000
 * 00001
 * 11011
 * 给定上图，返回结果 3 。
 *
 * 注意：
 *
 * 11
 * 1
 * 和
 *
 *  1
 * 11
 * 是不同的岛屿，因为我们不考虑旋转、翻转操作。
 *
 *
 *
 * 提示：二维数组每维的大小都不会超过 50 。
 *
 * 通过次数5,218
 * 提交次数9,629
 *
 */
var directions = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}
var curr []byte
var n, m int

func numDistinctIslands(grid [][]int) int {
	n, m = len(grid), len(grid[0])
	islands := make(map[string]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				continue
			}
			curr = nil
			dfs(grid, i, j, 0)
			islands[string(curr)] = struct{}{}
		}
	}
	return len(islands)
}

func dfs(grid [][]int, y, x int, direction byte) {
	if y < 0 || y >= n || x < 0 || x >= m || grid[y][x] == 0 {
		return
	}
	grid[y][x] = 0
	curr = append(curr, direction)
	for i, d := range directions {
		dfs(grid, y+d[0], x+d[1], byte(i))
	}
	curr = append(curr, '-', direction)
}
