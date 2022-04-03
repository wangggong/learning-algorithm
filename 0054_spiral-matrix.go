/*
 * @lc app=leetcode.cn id=spiral-matrix lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (48.42%)
 * Total Accepted:    231.5K
 * Total Submissions: 477.9K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */
func spiralOrder(matrix [][]int) []int {
	var ans []int
	n, m := len(matrix), len(matrix[0])
	start := 0
	for n > 0 && m > 0 {
		for i := 0; i < m; i++ {
			ans = append(ans, matrix[start][start+i])
		}
		for i := 1; i < n; i++ {
			ans = append(ans, matrix[start+i][start+m-1])
		}
		if n > 1 {
			for i := 1; i < m; i++ {
				ans = append(ans, matrix[start+n-1][start+m-1-i])
			}
		}
		if m > 1 {
			for i := 1; i < n-1; i++ {
				ans = append(ans, matrix[start+n-1-i][start])
			}
		}
		n -= 2
		m -= 2
		start++
	}
	return ans
}
