/*
 * @lc app=leetcode.cn id=non-overlapping-intervals lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [435] 无重叠区间
 *
 * https://leetcode-cn.com/problems/non-overlapping-intervals/description/
 *
 * algorithms
 * Medium (50.67%)
 * Total Accepted:    117.4K
 * Total Submissions: 231.6K
 * Testcase Example:  '[[1,2],[2,3],[3,4],[1,3]]'
 *
 * 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回
 * 需要移除区间的最小数量，使剩余区间互不重叠 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
 * 输出: 1
 * 解释: 移除 [1,3] 后，剩下的区间没有重叠。
 *
 *
 * 示例 2:
 *
 *
 * 输入: intervals = [ [1,2], [1,2], [1,2] ]
 * 输出: 2
 * 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
 *
 *
 * 示例 3:
 *
 *
 * 输入: intervals = [ [1,2], [2,3] ]
 * 输出: 0
 * 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * -5 * 10^4 <= starti < endi <= 5 * 10^4
 *
 *
 */
import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// assert len(intervals) > 0 && len(intervals) <= 1e5
	sort.Slice(intervals, func(p, q int) bool { return intervals[p][1] < intervals[q][1] })
	right := intervals[0][1]
	ans := 1
	for _, interval := range intervals[1:] {
		if interval[0] >= right {
			right = interval[1]
			ans++
		} else {
		}
	}
	return len(intervals) - ans
}

/*
 * // 奇怪的贪心.
 * func eraseOverlapIntervals(I [][]int) int {
 *     n := len(I)
 *     if n == 0 {
 *         return 0
 *     }
 *     sort.Slice(I, func(p, q int) bool {
 *         if I[p][1] != I[q][1] {
 *             return I[p][1] < I[q][1]
 *         }
 *         return I[p][0] > I[q][0]
 *     })
 *     ans := 0
 *     p, q := 0, 1
 *     for q < n {
 *         if I[p][1] <= I[q][0] {
 *             p, q = p+1, q+1
 *             continue
 *         }
 *         for q < n && I[p][1] > I[q][0] {
 *             q++
 *             ans++
 *         }
 *         p, q = q, q+1
 *     }
 *     return ans
 * }
 */

/*
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 *
 *
 * // dp, 好消息是 TLE 了, 官方解也 TLE
 * func eraseOverlapIntervals(intervals [][]int) int {
 * 	n := len(intervals)
 * 	sort.Slice(intervals, func(p, q int) bool { return intervals[p][0] < intervals[q][0] })
 * 	dp := make([]int, n)
 * 	for i := range dp {
 * 		dp[i] = 1
 * 	}
 * 	ans := 0
 * 	for i := 1; i < n; i++ {
 * 		for j := 0; j < i; j++ {
 * 			if intervals[j][1] <= intervals[i][0] {
 * 				dp[i] = max(dp[i], dp[j]+1)
 * 			}
 * 		}
 * 	}
 * 	for _, v := range dp {
 * 		ans = max(ans, v)
 * 	}
 * 	return n - ans
 * }
 */
