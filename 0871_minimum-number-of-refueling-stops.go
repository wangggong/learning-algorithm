/*
 * @lc app=leetcode.cn id=minimum-number-of-refueling-stops lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [871] 最低加油次数
 *
 * https://leetcode.cn/problems/minimum-number-of-refueling-stops/description/
 *
 * algorithms
 * Hard (35.89%)
 * Total Accepted:    17K
 * Total Submissions: 43.7K
 * Testcase Example:  '1\n1\n[]'
 *
 * 汽车从起点出发驶向目的地，该目的地位于出发位置东面 target 英里处。
 *
 * 沿途有加油站，每个 station[i] 代表一个加油站，它位于出发位置东面 station[i][0] 英里处，并且有 station[i][1]
 * 升汽油。
 *
 * 假设汽车油箱的容量是无限的，其中最初有 startFuel 升燃料。它每行驶 1 英里就会用掉 1 升汽油。
 *
 * 当汽车到达加油站时，它可能停下来加油，将所有汽油从加油站转移到汽车中。
 *
 * 为了到达目的地，汽车所必要的最低加油次数是多少？如果无法到达目的地，则返回 -1 。
 *
 * 注意：如果汽车到达加油站时剩余燃料为 0，它仍然可以在那里加油。如果汽车到达目的地时剩余燃料为 0，仍然认为它已经到达目的地。
 *
 *
 *
 * 示例 1：
 *
 * 输入：target = 1, startFuel = 1, stations = []
 * 输出：0
 * 解释：我们可以在不加油的情况下到达目的地。
 *
 *
 * 示例 2：
 *
 * 输入：target = 100, startFuel = 1, stations = [[10,100]]
 * 输出：-1
 * 解释：我们无法抵达目的地，甚至无法到达第一个加油站。
 *
 *
 * 示例 3：
 *
 * 输入：target = 100, startFuel = 10, stations =
 * [[10,60],[20,30],[30,30],[60,40]]
 * 输出：2
 * 解释：
 * 我们出发时有 10 升燃料。
 * 我们开车来到距起点 10 英里处的加油站，消耗 10 升燃料。将汽油从 0 升加到 60 升。
 * 然后，我们从 10 英里处的加油站开到 60 英里处的加油站（消耗 50 升燃料），
 * 并将汽油从 10 升加到 50 升。然后我们开车抵达目的地。
 * 我们沿途在1两个加油站停靠，所以返回 2 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= target, startFuel, stations[i][1] <= 10^9
 * 0 <= stations.length <= 500
 * 0 < stations[0][0] < stations[1][0] < ... < stations[stations.length-1][0] <
 * target
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

func (h Heap) Less(p, q int) bool  { return h.IntSlice[p] > h.IntSlice[q] }
func (h *Heap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Heap) Pop() (v interface{}) {
	l := len(h.IntSlice)
	v = h.IntSlice[l-1]
	h.IntSlice = h.IntSlice[:l-1]
	return
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	loc, remain, ans, i, n := 0, startFuel, 0, 0, len(stations)
	h := &Heap{}
	for loc < target {
		if remain == 0 {
			if ans++; h.Len() > 0 && ans >= 0 {
				remain += heap.Pop(h).(int)
			} else {
				return -1
			}
		}
		loc += remain
		remain = 0
		for ; i < n && stations[i][0] <= loc; i++ {
			heap.Push(h, stations[i][1])
		}
	}
	return ans
}
