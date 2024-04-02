/*
 * @lc app=leetcode.cn id=diagonal-traverse lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [498] 对角线遍历
 *
 * https://leetcode.cn/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (50.28%)
 * Total Accepted:    68.3K
 * Total Submissions: 131.4K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,4,7,5,3,6,8,9]
 *
 *
 * 示例 2：
 *
 *
 * 输入：mat = [[1,2],[3,4]]
 * 输出：[1,2,3,4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 1 <= m, n <= 10^4
 * 1 <= m * n <= 10^4
 * -10^5 <= mat[i][j] <= 10^5
 *
 *
 */
func findDiagonalOrder(M [][]int) []int {
	n, m := len(M), len(M[0])
	var ans []int
	r, c := 0, 0
	for i := 0; i < m+n-1; i++ {
		if i%2 == 0 {
			for ; r < 0 || r >= n || c < 0 || c >= m; r, c = r-1, c+1 {
			}
			for ; 0 <= r && r < n && 0 <= c && c < m; r, c = r-1, c+1 {
				ans = append(ans, M[r][c])
			}
			c++
		} else {
			for ; r < 0 || r >= n || c < 0 || c >= m; r, c = r+1, c-1 {
			}
			for ; 0 <= r && r < n && 0 <= c && c < m; r, c = r+1, c-1 {
				ans = append(ans, M[r][c])
			}
			r++
		}
	}
	return ans
}
