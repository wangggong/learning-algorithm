/*
 * Medium
 *
 * 给定两个 稀疏矩阵 ：大小为 m x k 的稀疏矩阵 mat1 和大小为 k x n 的稀疏矩阵 mat2 ，返回 mat1 x mat2 的结果。你可以假设乘法总是可能的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：mat1 = [[1,0,0],[-1,0,3]], mat2 = [[7,0,0],[0,0,0],[0,0,1]]
 * 输出：[[7,0,0],[-7,0,3]]
 *  示例 2:
 *
 * 输入：mat1 = [[0]], mat2 = [[0]]
 * 输出：[[0]]
 *
 *
 * 提示:
 *
 * m == mat1.length
 * k == mat1[i].length == mat2.length
 * n == mat2[i].length
 * 1 <= m, n, k <= 100
 * -100 <= mat1[i][j], mat2[i][j] <= 100
 * 通过次数4,997
 * 提交次数6,544
 */

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	// assert len(mat1[0]) == len(mat2)
	n, m, p := len(mat1), len(mat1[0]), len(mat2[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, p)
	}
	for i := 0; i < n; i++ {
		for k := 0; k < m; k++ {
            // 提示稀疏了, 所以遇到 `0` 可以先考虑跳过.
			if mat1[i][k] == 0 {
				continue
			}
			for j := 0; j < p; j++ {
				ans[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}
	return ans
}
