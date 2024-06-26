/*
 * @lc app=leetcode.cn id=set-intersection-size-at-least-two lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [757] 设置交集大小至少为2
 *
 * https://leetcode.cn/problems/set-intersection-size-at-least-two/description/
 *
 * algorithms
 * Hard (43.67%)
 * Total Accepted:    5.7K
 * Total Submissions: 11.2K
 * Testcase Example:  '[[1,3],[1,4],[2,5],[3,5]]'
 *
 * 一个整数区间 [a, b]  ( a < b ) 代表着从 a 到 b 的所有连续整数，包括 a 和 b。
 *
 * 给你一组整数区间intervals，请找到一个最小的集合 S，使得 S 里的元素与区间intervals中的每一个整数区间都至少有2个元素相交。
 *
 * 输出这个最小集合S的大小。
 *
 * 示例 1:
 *
 * 输入: intervals = [[1, 3], [1, 4], [2, 5], [3, 5]]
 * 输出: 3
 * 解释:
 * 考虑集合 S = {2, 3, 4}. S与intervals中的四个区间都有至少2个相交的元素。
 * 且这是S最小的情况，故我们输出3。
 *
 *
 * 示例 2:
 *
 * 输入: intervals = [[1, 2], [2, 3], [2, 4], [4, 5]]
 * 输出: 5
 * 解释:
 * 最小的集合S = {1, 2, 3, 4, 5}.
 *
 *
 * 注意:
 *
 *
 * intervals 的长度范围为[1, 3000]。
 * intervals[i] 长度为 2，分别代表左、右边界。
 * intervals[i][j] 的值是 [0, 10^8]范围内的整数。
 *
 *
 */

import "sort"

// 老了, 不少东西不会贪了. 排序后贪心结果贪错了...
func intersectionSizeTwo(I [][]int) int {
	sort.Slice(I, func(p, q int) bool { return I[p][0] < I[q][0] || I[p][0] == I[q][0] && I[p][1] > I[q][1] })
	var S []int
	for i := len(I) - 1; i >= 0; i-- {
		c := 2
		for j := max(len(S)-2, 0); j < len(S); j++ {
			if I[i][1] >= S[j] {
				c--
			}
		}
		for j := I[i][0] + c - 1; c > 0; j, c = j-1, c-1 {
			S = append(S, j)
		}
	}
	return len(S)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
