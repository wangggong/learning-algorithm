/*
 * @lc app=leetcode.cn id=minimum-path-sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (69.09%)
 * Total Accepted:    333.3K
 * Total Submissions: 482.3K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 100
 *
 *
 */

const MAXN int = 200

var dp [MAXN + 10][MAXN + 10]int

func minPathSum(G [][]int) int {
	n, m := len(G), len(G[0])
	dp[0][0] = G[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + G[i][0]
	}
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + G[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + G[i][j]
		}
	}
	return dp[n-1][m-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
