/*
 * @lc app=leetcode.cn id=最大矩形 lang=golang
 *
 * [85] 最大矩形
 *
 * https://leetcode-cn.com/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (52.05%)
 * Total Accepted:    109.6K
 * Total Submissions: 210.5K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：6
 * 解释：最大矩形如上图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = []
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [["0"]]
 * 输出：0
 *
 *
 * 示例 4：
 *
 *
 * 输入：matrix = [["1"]]
 * 输出：1
 *
 *
 * 示例 5：
 *
 *
 * 输入：matrix = [["0","0"]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * rows == matrix.length
 * cols == matrix[0].length
 * 1 <= row, cols <= 200
 * matrix[i][j] 为 '0' 或 '1'
 *
 *
 */
// import "fmt"

const MAXN = 200 + 10

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximalRectangle(M [][]byte) int {
	if len(M) == 0 {
		return 0
	}

	n, m := len(M), len(M[0])
	var H [MAXN][MAXN]int

	for i := 0; i < m; i++ {
		if M[0][i] == '1' {
			H[0][i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if M[i][j] == '1' {
				H[i][j] = H[i-1][j] + 1
			}
		}
	}

	// fmt.Printf("%v\n", H)
	ans := 0

	for i := 0; i < n; i++ {
		var SL, SR []int
		up, down := make([]int, m), make([]int, m)

		for j := 0; j < m; j++ {
			for len(SL) > 0 && H[i][SL[len(SL)-1]] >= H[i][j] {
				SL = SL[:len(SL)-1]
			}
			if len(SL) == 0 {
				up[j] = -1
			} else {
				up[j] = SL[len(SL)-1]
			}
			SL = append(SL, j)
		}

		for j := m - 1; j >= 0; j-- {
			for len(SR) > 0 && H[i][SR[len(SR)-1]] >= H[i][j] {
				SR = SR[:len(SR)-1]
			}
			if len(SR) == 0 {
				down[j] = m
			} else {
				down[j] = SR[len(SR)-1]
			}
			SR = append(SR, j)
		}

		// fmt.Printf("%v\n%v\n\n", up, down)

		for j := 0; j < m; j++ {
			ans = max(ans, H[i][j]*(down[j]-up[j]-1))
		}
	}

	return ans
}
