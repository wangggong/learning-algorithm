/*
 * @lc app=leetcode.cn id=divide-intervals-into-minimum-number-of-groups lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2406] 将区间分为最少组数
 *
 * https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups/description/
 *
 * algorithms
 * Medium (38.68%)
 * Total Accepted:    4.9K
 * Total Submissions: 12.7K
 * Testcase Example:  '[[5,10],[6,8],[1,5],[2,3],[1,10]]'
 *
 * 给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示 闭 区间 [lefti,
 * righti] 。
 *
 * 你需要将 intervals 划分为一个或者多个区间 组 ，每个区间 只 属于一个组，且同一个组中任意两个区间 不相交 。
 *
 * 请你返回 最少 需要划分成多少个组。
 *
 * 如果两个区间覆盖的范围有重叠（即至少有一个公共数字），那么我们称这两个区间是 相交 的。比方说区间 [1, 5] 和 [5, 8] 相交。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
 * 输出：3
 * 解释：我们可以将区间划分为如下的区间组：
 * - 第 1 组：[1, 5] ，[6, 8] 。
 * - 第 2 组：[2, 3] ，[5, 10] 。
 * - 第 3 组：[1, 10] 。
 * 可以证明无法将区间划分为少于 3 个组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,3],[5,6],[8,10],[11,13]]
 * 输出：1
 * 解释：所有区间互不相交，所以我们可以把它们全部放在一个组内。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * 1 <= lefti <= righti <= 10^6
 *
 *
 */

import (
	"container/heap"
	"sort"
)

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Heap) Pop() (v interface{}) {
	l := h.Len()
	v = h.IntSlice[l-1]
	h.IntSlice = h.IntSlice[:l-1]
	return
}

func minGroups(I [][]int) int {
	sort.Slice(I, func(p, q int) bool { return I[p][0] < I[q][0] })
	h := &Heap{}
	for _, i := range I {
		if h.Len() > 0 && h.IntSlice[0] < i[0] {
			heap.Pop(h)
		}
		heap.Push(h, i[1])
	}
	return h.Len()
}
