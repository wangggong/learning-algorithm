/*
 * @lc app=leetcode.cn id=maximal-square lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [221] 最大正方形
 *
 * https://leetcode-cn.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (48.13%)
 * Total Accepted:    162.2K
 * Total Submissions: 336.5K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [["0","1"],["1","0"]]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [["0"]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 300
 * matrix[i][j] 为 '0' 或 '1'
 *
 *
 */
// import "fmt"

const MAXN = 300 + 10

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
 * func maximalSquare(M [][]byte) int {
 * 	n, m := len(M), len(M[0])
 * 	var dp [MAXN][MAXN]int
 * 	ans := 0
 * 	for i := 0; i < n; i++ {
 * 		if M[i][0] == '1' {
 * 			dp[i][0] = 1
 * 			ans = max(ans, dp[i][0])
 * 		}
 * 	}
 * 	for i := 0; i < m; i++ {
 * 		if M[0][i] == '1' {
 * 			dp[0][i] = 1
 * 			ans = max(ans, dp[0][i])
 * 		}
 * 	}
 * 	for i := 1; i < n; i++ {
 * 		for j := 1; j < m; j++ {
 * 			if M[i][j] != '1' {
 * 				continue
 * 			}
 * 			dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
 * 			ans = max(ans, dp[i][j])
 * 		}
 * 	}
 * 	return ans * ans
 * }
 */

func maximalSquare(M [][]byte) int {
	n, m := len(M), len(M[0])
	var H [MAXN][MAXN]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				if M[i][j] == '1' {
					H[i][j]++
				}
				continue
			}
			if M[i][j] == '1' {
				H[i][j] = H[i-1][j] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		var t []int
		S := make([]int, 0, m)
		t = append(t, 0)
		for j := 0; j < m; j++ {
			t = append(t, H[i][j])
		}
		t = append(t, 0)
		for j := 0; j < len(t); j++ {
			for len(S) > 0 && t[S[len(S)-1]] > t[j] {
				h := t[S[len(S)-1]]
				S = S[:len(S)-1]
				w := j - S[len(S)-1] - 1
				m := min(h, w)
				// fmt.Printf("%v %v %v %v %v\n", t, S, j, h, w)
				ans = max(ans, m*m)
			}
			S = append(S, j)
		}
	}
	return ans
}
