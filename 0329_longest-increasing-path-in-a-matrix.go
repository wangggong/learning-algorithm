/*
 * @lc app=leetcode.cn id=longest-increasing-path-in-a-matrix lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [329] 矩阵中的最长递增路径
 *
 * https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/description/
 *
 * algorithms
 * Hard (49.48%)
 * Total Accepted:    64.9K
 * Total Submissions: 131.1K
 * Testcase Example:  '[[9,9,4],[6,6,8],[2,1,1]]'
 *
 * 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
 *
 * 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
 * 输出：4
 * 解释：最长递增路径为 [1, 2, 6, 9]。
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
 * 输出：4
 * 解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [[1]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 200
 * 0 <= matrix[i][j] <= (1 << 31) - 1
 *
 *
 */

const MAXN = 200 + 10

var visit [MAXN][MAXN]bool
var dp [MAXN][MAXN]int
var outdegree [MAXN][MAXN]int

var directions = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var m, n int

/*
 * func longestIncreasingPath(matrix [][]int) int {
 * 	n, m = len(matrix), len(matrix[0])
 * 	ans := 0
 * 	for i := 0; i < MAXN; i++ {
 * 		for j := 0; j < MAXN; j++ {
 * 			visit[i][j], dp[i][j] = false, 0
 * 		}
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		for j := 0; j < m; j++ {
 * 			v := dfs(matrix, i, j)
 * 			if v > ans {
 * 				ans = v
 * 			}
 * 		}
 * 	}
 * 	return ans
 * }
 */

func dfs(arr [][]int, row, col int) int {
	// assert 0 <= row && row < n && 0 <= col && col < m;
	// if row < 0 || row >= n || col < 0 || col >= m {
	// 	return nil
	// }
	if visit[row][col] {
		return dp[row][col]
	}
	dp[row][col] = 1
	defer func() { visit[row][col] = true }()
	for _, d := range directions {
		nrow, ncol := row+d[0], col+d[1]
		if nrow < 0 || nrow >= n || ncol < 0 || ncol >= m {
			continue
		}
		if arr[nrow][ncol] >= arr[row][col] {
			continue
		}
		v := dfs(arr, nrow, ncol) + 1
		if v > dp[row][col] {
			dp[row][col] = v
		}
	}
	return dp[row][col]
}

func longestIncreasingPath(matrix [][]int) int {
	n, m = len(matrix), len(matrix[0])
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			outdegree[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if ni < 0 || ni >= n || nj < 0 || nj >= m {
					continue
				}
				if matrix[ni][nj] > matrix[i][j] {
					outdegree[i][j]++
				}
			}
		}
	}
	var Q [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if outdegree[i][j] == 0 {
				Q = append(Q, []int{i, j})
			}
		}
	}
	ans := 0
	for len(Q) > 0 {
		// 每次把一批数据都 pop 出来, 代表出度在同一 "层"
		ans++
		size := len(Q)
		for i := 0; i < size; i++ {
			q := Q[0]
			Q = Q[1:]
			for _, d := range directions {
				ni, nj := q[0]+d[0], q[1]+d[1]
				if ni < 0 || ni >= n || nj < 0 || nj >= m {
					continue
				}
				// 将比当前值小的值入队.
				if matrix[ni][nj] < matrix[q[0]][q[1]] {
					outdegree[ni][nj]--
					if outdegree[ni][nj] == 0 {
						Q = append(Q, []int{ni, nj})
					}
				}
			}
		}
	}
	return ans
}
