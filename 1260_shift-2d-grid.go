/*
 * @lc app=leetcode.cn id=shift-2d-grid lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1260] 二维网格迁移
 *
 * https://leetcode.cn/problems/shift-2d-grid/description/
 *
 * algorithms
 * Easy (60.44%)
 * Total Accepted:    16K
 * Total Submissions: 26.5K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]\n1'
 *
 * 给你一个 m 行 n 列的二维网格 grid 和一个整数 k。你需要将 grid 迁移 k 次。
 *
 * 每次「迁移」操作将会引发下述活动：
 *
 *
 * 位于 grid[i][j] 的元素将会移动到 grid[i][j + 1]。
 * 位于 grid[i][n - 1] 的元素将会移动到 grid[i + 1][0]。
 * 位于 grid[m - 1][n - 1] 的元素将会移动到 grid[0][0]。
 *
 *
 * 请你返回 k 次迁移操作后最终得到的 二维网格。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
 * 输出：[[9,1,2],[3,4,5],[6,7,8]]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
 * 输出：[[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
 * 输出：[[1,2,3],[4,5,6],[7,8,9]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m <= 50
 * 1 <= n <= 50
 * -1000 <= grid[i][j] <= 1000
 * 0 <= k <= 100
 *
 *
 */
func shiftGrid(G [][]int, k int) [][]int {
	n, m := len(G), len(G[0])
	next := make([][]int, n)
	for i := range next {
		next[i] = make([]int, m)
	}
	for ; k > 0; k-- {
		for i := 0; i < n; i++ {
			for j := 0; j+1 < m; j++ {
				next[i][j+1] = G[i][j]
			}
		}
		for i := 0; i+1 < n; i++ {
			next[i+1][0] = G[i][m-1]
		}
		next[0][0] = G[n-1][m-1]
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				G[i][j] = next[i][j]
			}
		}
	}
	return G
}
