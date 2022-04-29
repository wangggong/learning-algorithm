/*
 * @lc app=leetcode.cn id=course-schedule-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [630] 课程表 III
 *
 * https://leetcode-cn.com/problems/course-schedule-iii/description/
 *
 * algorithms
 * Hard (46.32%)
 * Total Accepted:    24.5K
 * Total Submissions: 53K
 * Testcase Example:  '[[100,200],[200,1300],[1000,1250],[2000,3200]]'
 *
 * 这里有 n 门不同的在线课程，按从 1 到 n 编号。给你一个数组 courses ，其中 courses[i] = [durationi,
 * lastDayi] 表示第 i 门课将会 持续 上 durationi 天课，并且必须在不晚于 lastDayi 的时候完成。
 *
 * 你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。
 *
 * 返回你最多可以修读的课程数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：courses = [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
 * 输出：3
 * 解释：
 * 这里一共有 4 门课程，但是你最多可以修 3 门：
 * 首先，修第 1 门课，耗费 100 天，在第 100 天完成，在第 101 天开始下门课。
 * 第二，修第 3 门课，耗费 1000 天，在第 1100 天完成，在第 1101 天开始下门课程。
 * 第三，修第 2 门课，耗时 200 天，在第 1300 天完成。
 * 第 4 门课现在不能修，因为将会在第 3300 天完成它，这已经超出了关闭日期。
 *
 * 示例 2：
 *
 *
 * 输入：courses = [[1,2]]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：courses = [[3,2],[4,3]]
 * 输出：0
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= courses.length <= 10^4
 * 1 <= durationi, lastDayi <= 10^4
 *
 *
 */

import (
	"container/heap"
	"sort"
)

type PriorityQueue struct {
	Arr []int
}

func (pq PriorityQueue) Len() int            { return len(pq.Arr) }
func (pq PriorityQueue) Less(p, q int) bool  { return pq.Arr[p] > pq.Arr[q] }
func (pq PriorityQueue) Swap(p, q int)       { pq.Arr[p], pq.Arr[q] = pq.Arr[q], pq.Arr[p] }
func (pq *PriorityQueue) Push(a interface{}) { pq.Arr = append(pq.Arr, a.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	l := len(pq.Arr)
	v := pq.Arr[l-1]
	pq.Arr = pq.Arr[:l-1]
	return v
}

// 卡了 1h, 人麻了...
// 参考: https://leetcode-cn.com/problems/course-schedule-iii/solution/gong-shui-san-xie-jing-dian-tan-xin-yun-ghii2/
func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(p, q int) bool { return courses[p][1] < courses[q][1] })
	time := 0
	pq := &PriorityQueue{}
	for _, c := range courses {
		time += c[0]
		heap.Push(pq, c[0])
		if time > c[1] {
			time -= heap.Pop(pq).(int)
		}
	}
	return pq.Len()
}

// 二刷还是没思路, 常看常新吧快...
