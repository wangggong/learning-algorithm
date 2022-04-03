/*
 * @lc app=leetcode.cn id=k-closest-points-to-origin lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [973] 最接近原点的 K 个点
 *
 * https://leetcode-cn.com/problems/k-closest-points-to-origin/description/
 *
 * algorithms
 * Medium (64.50%)
 * Total Accepted:    74.7K
 * Total Submissions: 115.8K
 * Testcase Example:  '[[1,3],[-2,2]]\n1'
 *
 * 给定一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点，并且是一个整数 k ，返回离原点 (0,0)
 * 最近的 k 个点。
 *
 * 这里，平面上两点之间的距离是 欧几里德距离（ √(x1 - x2)^2 + (y1 - y2)^2 ）。
 *
 * 你可以按 任何顺序 返回答案。除了点坐标的顺序之外，答案 确保 是 唯一 的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：points = [[1,3],[-2,2]], k = 1
 * 输出：[[-2,2]]
 * 解释：
 * (1, 3) 和原点之间的距离为 sqrt(10)，
 * (-2, 2) 和原点之间的距离为 sqrt(8)，
 * 由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
 * 我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
 *
 *
 * 示例 2：
 *
 *
 * 输入：points = [[3,3],[5,-1],[-2,4]], k = 2
 * 输出：[[3,3],[-2,4]]
 * （答案 [[-2,4],[3,3]] 也会被接受。）
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= points.length <= 10^4
 * -10^4 < xi, yi < 10^4
 *
 *
 */

import "container/heap"

type Heap struct {
	Arr [][]int
}

func (h Heap) Len() int { return len(h.Arr) }
func (h Heap) Less(p, q int) bool {
	x1, y1 := h.Arr[p][0], h.Arr[p][1]
	x2, y2 := h.Arr[q][0], h.Arr[q][1]
	return x1*x1+y1*y1 < x2*x2+y2*y2
}
func (h Heap) Swap(p, q int)         { h.Arr[p], h.Arr[q] = h.Arr[q], h.Arr[p] }
func (h *Heap) Push(v interface{})   { h.Arr = append(h.Arr, v.([]int)) }
func (h *Heap) Pop() (v interface{}) { l := len(h.Arr); v = h.Arr[l-1]; h.Arr = h.Arr[:l-1]; return }

func kClosest(points [][]int, k int) [][]int {
	h := &Heap{}
	for _, p := range points {
		heap.Push(h, p)
	}
	var ans [][]int
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(h).([]int))
	}
	return ans
}
