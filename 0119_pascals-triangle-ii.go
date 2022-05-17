/*
 * @lc app=leetcode.cn id=pascals-triangle-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (68.32%)
 * Total Accepted:    201.5K
 * Total Submissions: 294.9K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
 *
 * 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: rowIndex = 3
 * 输出: [1,3,3,1]
 *
 *
 * 示例 2:
 *
 *
 * 输入: rowIndex = 0
 * 输出: [1]
 *
 *
 * 示例 3:
 *
 *
 * 输入: rowIndex = 1
 * 输出: [1,1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 0 <= rowIndex <= 33
 *
 *
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
 *
 */
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	var rows [2][]int
	for i := range rows {
		rows[i] = make([]int, rowIndex+1)
	}
	rows[1][0] = 1
	rows[1][1] = 1
	for i := 2; i <= rowIndex; i++ {
		rows[i%2][0] = 1
		for j := 1; j < i; j++ {
			rows[i%2][j] = rows[(i-1)%2][j-1] + rows[(i-1)%2][j]
		}
		rows[i%2][i] = 1
	}
	return rows[rowIndex%2]
}
