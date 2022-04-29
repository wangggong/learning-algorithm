/*
 * @lc app=leetcode.cn id=max-sum-of-rectangle-no-larger-than-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [363] 矩形区域不超过 K 的最大数值和
 *
 * https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/description/
 *
 * algorithms
 * Hard (48.64%)
 * Total Accepted:    35.9K
 * Total Submissions: 73.8K
 * Testcase Example:  '[[1,0,1],[0,-2,3]]\n2'
 *
 * 给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
 *
 * 题目数据保证总会存在一个数值和不超过 k 的矩形区域。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,0,1],[0,-2,3]], k = 2
 * 输出：2
 * 解释：蓝色边框圈出来的矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[2,2,-1]], k = 3
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -100
 * -10^5
 *
 *
 *
 *
 * 进阶：如果行数远大于列数，该如何设计解决方案？
 *
 */

/*
 * import "math"
 *
 * // 现场是直接暴力了一波, 所幸数据集没有整大活, 反倒很顺利.
 * func maxSumSubmatrix(M [][]int, k int) int {
 * 	n, m := len(M), len(M[0])
 * 	ans := math.MinInt32
 * 	ps := make([][]int, n+1)
 * 	for i := range ps {
 * 		ps[i] = make([]int, m+1)
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		for j := 0; j < m; j++ {
 * 			ps[i+1][j+1] = ps[i+1][j] + ps[i][j+1] + M[i][j] - ps[i][j]
 * 		}
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		for j := i + 1; j <= n; j++ {
 * 			for p := 0; p < m; p++ {
 * 				for q := p + 1; q <= m; q++ {
 * 					cur := ps[j][q] - ps[j][p] - ps[i][q] + ps[i][p]
 * 					if cur <= k && cur > ans {
 * 						ans = cur
 * 					}
 * 				}
 * 			}
 * 		}
 * 	}
 * 	return ans
 * }
 */

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func maxSumSubmatrix(M [][]int, K int) int {
	n, m := len(M), len(M[0])
	preSum := make([][]int, n+1)
	for i := range preSum {
		preSum[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + M[i][j]
		}
	}
	ans := math.MinInt32
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			treap := redblacktree.NewWithIntComparator()
			treap.Put(0, 0)
			for k := 1; k <= m; k++ {
				right := preSum[j][k] - preSum[i][k]
				node, _ := treap.Ceiling(right - K)
				if node != nil {
					left := node.Key.(int)
					if right-left <= K && right-left > ans {
						ans = right - left
					}
				}
				treap.Put(right, k)
			}
		}
	}
	return ans
}
