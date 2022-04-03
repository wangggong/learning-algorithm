/*
 * @lc app=leetcode.cn id=matrix-block-sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1314] 矩阵区域和
 *
 * https://leetcode-cn.com/problems/matrix-block-sum/description/
 *
 * algorithms
 * Medium (75.08%)
 * Total Accepted:    15.3K
 * Total Submissions: 20.4K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]\n1'
 *
 * 给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素
 * mat[r][c] 的和：
 *
 *
 * i - k
 * j - k  且
 * (r, c) 在矩阵内。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
 * 输出：[[12,21,16],[27,45,33],[24,39,28]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
 * 输出：[[45,45,45],[45,45,45],[45,45,45]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 1 <= n, m, k <= 100
 * 1 <= mat[i][j] <= 100
 *
 *
 */

/*
 * func matrixBlockSum(mat [][]int, k int) [][]int {
 * 	n, m := len(mat), len(mat[0])
 * 	prefixSum := make([][]int, 0, n)
 * 	for _, row := range mat {
 * 		ps := make([]int, m+1)
 * 		for i, v := range row {
 * 			ps[i+1] = ps[i] + v
 * 		}
 * 		prefixSum = append(prefixSum, ps)
 * 	}
 * 	ans := make([][]int, 0, n)
 * 	for i := 0; i < n; i++ {
 * 		ans = append(ans, make([]int, m))
 * 		for j := 0; j < m; j++ {
 * 			for d := -k; d <= k; d++ {
 * 				if i+d < 0 || i+d >= n {
 * 					continue
 * 				}
 * 				left, right := j-k, j+k
 * 				if left < 0 {
 * 					left = 0
 * 				}
 * 				if right+1 > m {
 * 					right = m - 1
 * 				}
 * 				ans[i][j] += prefixSum[i+d][right+1] - prefixSum[i+d][left]
 * 			}
 * 		}
 * 	}
 * 	return ans
 * }
 */

const MAXN = 100 + 10

var prefixSum [MAXN][MAXN]int

// 二维前缀和: prefix[row2][col2] - prefix[row1][col1] + prefix[row1][col2] + prefix[row2][col1]
func matrixBlockSum(mat [][]int, k int) [][]int {
	n, m := len(mat), len(mat[0])
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			prefixSum[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			prefixSum[i+1][j+1] = prefixSum[i+1][j] + prefixSum[i][j+1] - prefixSum[i][j] + mat[i][j]
		}
	}
	ans := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, m))
		for j := 0; j < m; j++ {
			x1, x2, y1, y2 := j-k, j+k+1, i-k, i+k+1
			if x1 < 0 {
				x1 = 0
			}
			if x2 > m {
				x2 = m
			}
			if y1 < 0 {
				y1 = 0
			}
			if y2 > n {
				y2 = n
			}
			ans[i][j] = prefixSum[y2][x2] + prefixSum[y1][x1] - prefixSum[y1][x2] - prefixSum[y2][x1]
		}
	}
	return ans
}
