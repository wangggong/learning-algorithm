/*
 * @lc app=leetcode.cn id=set-matrix-zeroes lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [73] 矩阵置零
 *
 * https://leetcode-cn.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (61.83%)
 * Total Accepted:    185.8K
 * Total Submissions: 300.5K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出：[[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个仅使用常量空间的解决方案吗？
 *
 *
 */
func setZeroes(M [][]int) {
	n, m := len(M), len(M[0])
	row, col := false, false

	for i := 0; i < n; i++ {
		if M[i][0] == 0 {
			col = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if M[0][i] == 0 {
			row = true
			break
		}
	}

	// 注意, 用第 `0` 行与第 `0` 列标记后无论是查询还是刷行列都要从第 `1` 行与第 `1` 列开始.
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if M[i][j] == 0 {
				M[i][0] = 0
				M[0][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		if M[i][0] != 0 {
			continue
		}
		for j := 0; j < m; j++ {
			M[i][j] = 0
		}
	}
	for j := 1; j < m; j++ {
		if M[0][j] != 0 {
			continue
		}
		for i := 0; i < n; i++ {
			M[i][j] = 0
		}
	}

	if col {
		for i := 0; i < n; i++ {
			M[i][0] = 0
		}
	}
	if row {
		for i := 0; i < m; i++ {
			M[0][i] = 0
		}
	}
	return
}
