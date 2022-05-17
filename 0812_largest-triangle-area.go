/*
 * @lc app=leetcode.cn id=largest-triangle-area lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [812] 最大三角形面积
 *
 * https://leetcode.cn/problems/largest-triangle-area/description/
 *
 * algorithms
 * Easy (63.13%)
 * Total Accepted:    17K
 * Total Submissions: 26K
 * Testcase Example:  '[[0,0],[0,1],[1,0],[0,2],[2,0]]'
 *
 * 给定包含多个点的集合，从其中取三个点组成三角形，返回能组成的最大三角形的面积。
 *
 *
 * 示例:
 * 输入: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
 * 输出: 2
 * 解释:
 * 这五个点如下图所示。组成的橙色三角形是最大的，面积为2。
 *
 *
 *
 *
 * 注意:
 *
 *
 * 3 <= points.length <= 50.
 * 不存在重复的点。
 * -50 <= points[i][j] <= 50.
 * 结果误差值在 10^-6 以内都认为是正确答案。
 *
 *
 */

import "math"

const eps = 1e-6

func largestTriangleArea(P [][]int) float64 {
	n := len(P)
	ans := float64(0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				ans = math.Max(ans, Area(P[i], P[j], P[k]))
			}
		}
	}
	return ans
}

func Area(A, B, C []int) float64 {
	a, b, c := dist(B, C), dist(A, C), dist(A, B)
	p := (a + b + c) / 2.0
	if math.Abs(p-a) < eps || math.Abs(p-b) < eps || math.Abs(p-c) < eps {
		return 0.0
	}
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func dist(A, B []int) float64 {
	return math.Sqrt(float64((A[0]-B[0])*(A[0]-B[0]) + (A[1]-B[1])*(A[1]-B[1])))
}

