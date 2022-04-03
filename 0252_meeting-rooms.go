/*
 * Easy
 *
 * 给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，请你判断一个人是否能够参加这里面的全部会议。
 *
 *
 *
 * 示例 1：
 *
 * 输入：intervals = [[0,30],[5,10],[15,20]]
 * 输出：false
 * 示例 2：
 *
 * 输入：intervals = [[7,10],[2,4]]
 * 输出：true
 *
 *
 * 提示：
 *
 * 0 <= intervals.length <= 104
 * intervals[i].length == 2
 * 0 <= starti < endi <= 106
 * 通过次数15,425
 * 提交次数27,191
 */

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	n := len(intervals)
	sort.Slice(intervals, func(p, q int) bool { return intervals[p][0] < intervals[q][0] })
	for i := 0; i+1 < n; i++ {
		// if intervals[i][1] >= intervals[i+1][0] {
		// 需要问: 区间首尾重叠可以接受吗?
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}
