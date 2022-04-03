/*
 * @lc app=leetcode.cn id=spiral-matrix-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (77.62%)
 * Total Accepted:    169.3K
 * Total Submissions: 219.7K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：[[1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 20
 * 
 * 
 */


func generateMatrix(n int) [][]int {
	ans := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, n))
	}
	start := 0
	k := 1
	for n > 0 {
		for i := 0; i < n; i++ {
			ans[start][start+i] = k
			k++
		}
		for i := 1; i < n; i++ {
			ans[start+i][start+n-1] = k
			k++
		}
		if n > 1 {
			for i := 1; i < n; i++ {
				ans[start+n-1][start+n-1-i] = k
				k++
			}
		}
		if n > 1 {
			for i := 1; i+1 < n; i++ {
				ans[start+n-1-i][start] = k
				k++
			}
		}
		n -= 2
		start++
	}
	return ans
}
