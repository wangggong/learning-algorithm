/*
 * @lc app=leetcode.cn id=last-stone-weight lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1046] 最后一块石头的重量
 *
 * https://leetcode-cn.com/problems/last-stone-weight/description/
 *
 * algorithms
 * Easy (65.86%)
 * Total Accepted:    71.1K
 * Total Submissions: 108K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * 有一堆石头，每块石头的重量都是正整数。
 *
 * 每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x 。那么粉碎的可能结果如下：
 *
 *
 * 如果 x == y，那么两块石头都会被完全粉碎；
 * 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
 *
 *
 * 最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：[2,7,4,1,8,1]
 * 输出：1
 * 解释：
 * 先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
 * 再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
 * 接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
 * 最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 1000
 *
 *
 */
import (
	"container/heap"
	"sort"
)

type PQueue struct {
	sort.IntSlice
}

func (pq PQueue) Less(p, q int) bool  { return pq.IntSlice[p] > pq.IntSlice[q] }
func (pq PQueue) Swap(p, q int)       { pq.IntSlice[p], pq.IntSlice[q] = pq.IntSlice[q], pq.IntSlice[p] }
func (pq *PQueue) Push(v interface{}) { pq.IntSlice = append(pq.IntSlice, v.(int)) }
func (pq *PQueue) Pop() (v interface{}) {
	l := len(pq.IntSlice)
	v = pq.IntSlice[l-1]
	pq.IntSlice = pq.IntSlice[:l-1]
	return
}

func lastStoneWeight(stones []int) int {
	pq := &PQueue{}
	for _, s := range stones {
		heap.Push(pq, s)
	}
	for pq.Len() >= 2 {
		p := heap.Pop(pq).(int)
		q := heap.Pop(pq).(int)
		if p == q {
			continue
		}
		heap.Push(pq, abs(p-q))
	}
	if pq.Len() == 0 {
		return 0
	}
	return heap.Pop(pq).(int)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
