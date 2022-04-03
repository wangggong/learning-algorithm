/*
 * @lc app=leetcode.cn id=merge-intervals lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (47.94%)
 * Total Accepted:    387.3K
 * Total Submissions: 804.7K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]
 * 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(p, q int) bool {
		// if intervals[p][0] != intervals[q][0] {
		// 	return intervals[p][0] < intervals[q][0]
		// }
		// return intervals[p][1] < intervals[q][1]
		return intervals[p][0] < intervals[q][0]
	})
	var S [][]int
	for _, intrvl := range intervals {
		if len(S) > 0 && S[len(S)-1][1] >= intrvl[0] {
			curr := S[len(S)-1]
			if curr[0] > intrvl[0] {
				curr[0] = intrvl[0]
			}
			if curr[1] < intrvl[1] {
				curr[1] = intrvl[1]
			}
		} else {
			S = append(S, intrvl)
		}
	}
	return S
}

/*
 * func merge(I [][]int) [][]int {
 * 	n := len(I)
 * 	var ans [][]int
 * 	if n == 0 {
 * 		return ans
 * 	}
 * 	sort.Slice(I, func(p, q int) bool {
 * 		return I[p][0] < I[q][0]
 * 	})
 * 	curr := []int{I[0][0], I[0][1]}
 * 	for i := 1; i < n; i++ {
 * 		if curr[1] < I[i][0] {
 * 			ans = append(ans, curr)
 * 			curr = []int{I[i][0], I[i][1]}
 * 			continue
 * 		}
 * 		// 傻逼! 这里得取两个区间终点的较大值作为合并终点!
 * 		curr[1] = max(curr[1], I[i][1])
 * 	}
 * 	ans = append(ans, curr)
 * 	return ans
 * }
 *
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 */
