/*
 * Medium
 *
 * 给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，返回 所需会议室的最小数量 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：intervals = [[0,30],[5,10],[15,20]]
 * 输出：2
 * 示例 2：
 *
 * 输入：intervals = [[7,10],[2,4]]
 * 输出：1
 *
 *
 * 提示：
 *
 * 1 <= intervals.length <= 1e4
 * 0 <= starti < endi <= 1e6
 * 通过次数42,027
 * 提交次数82,356
 */

import "sort"

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(p, q int) bool { return intervals[p][0] < intervals[q][0] })
	ans := make([]int, 0, n)
	ans = append(ans, intervals[0][1])
	for i := 1; i < n; i++ {
		intrvl := intervals[i]
		found := false
		for j, a := range ans {
			if intrvl[0] >= a {
				ans[j] = intrvl[1]
				found = true
				// 找到合适的会议室就别往下走了傻逼!
				break
			}
		}
		if !found {
			ans = append(ans, intrvl[1])
		}
	}
	return len(ans)
}
