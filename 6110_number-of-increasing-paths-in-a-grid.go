/*
 * @lc app=leetcode.cn id=number-of-increasing-paths-in-a-grid lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6110] 网格图中递增路径的数目
 *
 * https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/description/
 *
 * algorithms
 * Hard (45.27%)
 * Total Accepted:    2.8K
 * Total Submissions: 6.1K
 * Testcase Example:  '[[1,1],[3,4]]'
 *
 * 给你一个 m x n 的整数网格图 grid ，你可以从一个格子移动到 4 个方向相邻的任意一个格子。
 *
 * 请你返回在网格图中从 任意 格子出发，达到 任意 格子，且路径中的数字是 严格递增 的路径数目。由于答案可能会很大，请将结果对 10^9 + 7 取余
 * 后返回。
 *
 * 如果两条路径中访问过的格子不是完全相同的，那么它们视为两条不同的路径。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：grid = [[1,1],[3,4]]
 * 输出：8
 * 解释：严格递增路径包括：
 * - 长度为 1 的路径：[1]，[1]，[3]，[4] 。
 * - 长度为 2 的路径：[1 -> 3]，[1 -> 4]，[3 -> 4] 。
 * - 长度为 3 的路径：[1 -> 3 -> 4] 。
 * 路径数目为 4 + 3 + 1 = 8 。
 *
 *
 * 示例 2：
 *
 * 输入：grid = [[1],[2]]
 * 输出：3
 * 解释：严格递增路径包括：
 * - 长度为 1 的路径：[1]，[2] 。
 * - 长度为 2 的路径：[1 -> 2] 。
 * 路径数目为 2 + 1 = 3 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 1000
 * 1 <= m * n <= 10^5
 * 1 <= grid[i][j] <= 10^5
 *
 *
 */
const mod int = 1e9 + 7

var n, m int
var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func countPaths(G [][]int) int {
	n, m = len(G), len(G[0])
	ans := 0
	memo := make(map[[2]int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = (ans + dfs(G, i, j, memo)) % mod
		}
	}
	return ans
}

func dfs(G [][]int, x, y int, memo map[[2]int]int) int {
	// assert 0 <= x && x < n && 0 <= y && y < m
	if v, ok := memo[[2]int{x, y}]; ok {
		return v
	}
	ans := 1
	defer func() { memo[[2]int{x, y}] = ans }()
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if !(0 <= nx && nx < n && 0 <= ny && ny < m) {
			continue
		}
		if G[nx][ny] <= G[x][y] {
			continue
		}
		ans = (ans + dfs(G, nx, ny, memo)) % mod
	}
	return ans
}


