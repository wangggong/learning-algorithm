/*
 * @lc app=leetcode.cn id=largest-rectangle-in-histogram lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (43.72%)
 * Total Accepted:    218.5K
 * Total Submissions: 498.5K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入：heights = [2,1,5,6,2,3]
 * 输出：10
 * 解释：最大的矩形为图中红色区域，面积为 10
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入： heights = [2,4]
 * 输出： 4
 *
 *
 *
 * 提示：
 *
 *
 *1 <= heights.length <=1e5
 *0 <= heights[i] <= 1e4
 *
 *
 */

const MAXN int = 1e5

var S []int

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	// NOTE: dummy 技巧:
	// assert len(S) > 0;    // 下同
	S = []int{-1}
	for i, h := range heights {
		for get(heights, S[len(S)-1]) >= h {
			S = S[:len(S)-1]
		}
		left[i] = S[len(S)-1]
		S = append(S, i)
	}
	S = []int{n}
	for i := n - 1; i >= 0; i-- {
		h := heights[i]
		for get(heights, S[len(S)-1]) >= h {
			S = S[:len(S)-1]
		}
		right[i] = S[len(S)-1]
		S = append(S, i)
	}
	ans := 0
	for i, h := range heights {
		if v := h * (right[i] - left[i] - 1); v > ans {
			ans = v
		}
	}
	return ans
}

func get(arr []int, index int) int {
	n := len(arr)
	if index < 0 || index >= n {
		return -MAXN
	}
	return arr[index]
}
