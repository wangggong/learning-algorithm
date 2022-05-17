/*
 * Hard
 *
 * 给定一个正整数、负整数和 0 组成的 N × M 矩阵，编写代码找出元素总和最大的子矩阵。
 *
 * 返回一个数组 [r1, c1, r2, c2]，其中 r1, c1 分别代表子矩阵左上角的行号和列号，r2, c2 分别代表右下角的行号和列号。若有多个满足条件的子矩阵，返回任意一个均可。
 *
 * 注意：本题相对书上原题稍作改动
 *
 * 示例：
 *
 * 输入：
 * [
 *    [-1,0],
 *    [0,-1]
 * ]
 * 输出：[0,1,0,1]
 * 解释：输入中标粗的元素即为输出所表示的矩阵
 *
 *
 * 说明：
 *
 * 1 <= matrix.length, matrix[0].length <= 200
 */

import "math"

/*
 * // 上场先来个暴力, 超时了.
 * func getMaxMatrix(M [][]int) []int {
 * 	n, m := len(M), len(M[0])
 * 	ps := make([][]int, n+1)
 * 	for i := range ps {
 * 		ps[i] = make([]int, m+1)
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		for j := 0; j < m; j++ {
 * 			ps[i+1][j+1] = ps[i+1][j] + ps[i][j+1] - ps[i][j] + M[i][j]
 * 		}
 * 	}
 * 	max, ans := math.MinInt32, []int{}
 * 	for r1 := 0; r1 < n; r1++ {
 * 		for c1 := 0; c1 < m; c1++ {
 * 			for r2 := r1 + 1; r2 <= n; r2++ {
 * 				for c2 := c1 + 1; c2 <= m; c2++ {
 * 					cur := ps[r2][c2] - ps[r2][c1] - ps[r1][c2] + ps[r1][c1]
 * 					if cur > max {
 * 						max = cur
 * 						ans = []int{r1, c1, r2 - 1, c2 - 1}
 * 					}
 * 				}
 * 			}
 * 		}
 * 	}
 * 	return ans
 * }
 */

// 优化: 类似最大子数组和的优化, 前面如果是负数就抛弃重新计数.
func getMaxMatrix(M [][]int) []int {
	n, m := len(M), len(M[0])
	ps := make([][]int, n+1)
	for i := range ps {
		ps[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ps[i+1][j+1] = ps[i+1][j] + ps[i][j+1] - ps[i][j] + M[i][j]
		}
	}
	max, ans := math.MinInt32, []int{}
	for r1 := 0; r1 < n; r1++ {
		for r2 := r1 + 1; r2 <= n; r2++ {
			c1 := 0
			for c2 := 1; c2 <= m; c2++ {
				cur := ps[r2][c2] - ps[r2][c1] - ps[r1][c2] + ps[r1][c1]
				if cur > max {
					max = cur
					ans = []int{r1, c1, r2 - 1, c2 - 1}
				}
				if cur < 0 {
					c1 = c2
				}
			}
		}
	}
	return ans
}
