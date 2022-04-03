/*
 * @lc app=leetcode.cn id=search-a-2d-matrix lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (46.84%)
 * Total Accepted:    211K
 * Total Submissions: 449.2K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -1e4 <= matrix[i][j], target <= 1e4
 *
 *
 */
func searchMatrix(M [][]int, target int) bool {
	n, m := len(M), len(M[0])
	if target < M[0][0] || target > M[n-1][m-1] {
		return false
	}
	row := -1
	p, q := 0, n-1
	for p <= q {
		mid := (p + q) >> 1
		if M[mid][0] <= target && M[mid][m-1] >= target {
			row = mid
			break
		} else if M[mid][0] > target {
			q = mid - 1
		} else {
			p = mid + 1
		}
	}
	if row < 0 {
		return false
	}
	p, q = 0, m-1
	for p <= q {
		mid := (p + q) >> 1
		if M[row][mid] == target {
			return true
		} else if M[row][mid] > target {
			q = mid - 1
		} else {
			p = mid + 1
		}
	}
	return false
}
