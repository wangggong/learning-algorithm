/*
 * @lc app=leetcode.cn id=maximum-sum-of-an-hourglass lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6193] 沙漏的最大总和
 *
 * https://leetcode.cn/problems/maximum-sum-of-an-hourglass/description/
 *
 * algorithms
 * Medium (75.25%)
 * Total Accepted:    4.8K
 * Total Submissions: 6.4K
 * Testcase Example:  '[[6,2,1,3],[4,2,1,5],[9,2,8,7],[4,1,2,9]]'
 *
 * 给你一个大小为 m x n 的整数矩阵 grid 。
 *
 * 按以下形式将矩阵的一部分定义为一个 沙漏 ：
 *
 * 返回沙漏中元素的 最大 总和。
 *
 * 注意：沙漏无法旋转且必须整个包含在矩阵中。
 *
 *
 *
 * 示例 1：
 *
 * 输入：grid = [[6,2,1,3],[4,2,1,5],[9,2,8,7],[4,1,2,9]]
 * 输出：30
 * 解释：上图中的单元格表示元素总和最大的沙漏：6 + 2 + 1 + 2 + 9 + 2 + 8 = 30 。
 *
 *
 * 示例 2：
 *
 * 输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：35
 * 解释：上图中的单元格表示元素总和最大的沙漏：1 + 2 + 3 + 5 + 7 + 8 + 9 = 35 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 3 <= m, n <= 150
 * 0 <= grid[i][j] <= 10^6
 *
 *
 */
func maxSum(G [][]int) (ans int) {
	n, m := len(G), len(G[0])
	for i := 1; i+1 < n; i++ {
		for j := 1; j+1 < m; j++ {
			ans = max(ans, get(G, i, j))
		}
	}
	return
}

func get(G [][]int, i, j int) int {
	return G[i-1][j-1] + G[i-1][j] + G[i-1][j+1] + G[i][j] + G[i+1][j-1] + G[i+1][j] + G[i+1][j+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
