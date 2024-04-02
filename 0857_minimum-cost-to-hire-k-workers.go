/*
 * @lc app=leetcode.cn id=minimum-cost-to-hire-k-workers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [857] 雇佣 K 名工人的最低成本
 *
 * https://leetcode.cn/problems/minimum-cost-to-hire-k-workers/description/
 *
 * algorithms
 * Hard (48.56%)
 * Total Accepted:    17.4K
 * Total Submissions: 27.1K
 * Testcase Example:  '[10,20,5]\n[70,50,30]\n2'
 *
 * 有 n 名工人。 给定两个数组 quality 和 wage ，其中，quality[i] 表示第 i 名工人的工作质量，其最低期望工资为
 * wage[i] 。
 *
 * 现在我们想雇佣 k 名工人组成一个工资组。在雇佣 一组 k 名工人时，我们必须按照下述规则向他们支付工资：
 *
 *
 * 对工资组中的每名工人，应当按其工作质量与同组其他工人的工作质量的比例来支付工资。
 * 工资组中的每名工人至少应当得到他们的最低期望工资。
 *
 *
 * 给定整数 k ，返回 组成满足上述条件的付费群体所需的最小金额 。在实际答案的 10^-5 以内的答案将被接受。。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： quality = [10,20,5], wage = [70,50,30], k = 2
 * 输出： 105.00000
 * 解释： 我们向 0 号工人支付 70，向 2 号工人支付 35。
 *
 * 示例 2：
 *
 *
 * 输入： quality = [3,1,10,10,1], wage = [4,8,2,2,7], k = 3
 * 输出： 30.66667
 * 解释： 我们向 0 号工人支付 4，向 2 号和 3 号分别支付 13.33333。
 *
 *
 *
 * 提示：
 *
 *
 * n == quality.length == wage.length
 * 1 <= k <= n <= 10^4
 * 1 <= quality[i], wage[i] <= 10^4
 *
 *
 */

import (
	"container/heap"
	"math"
	"sort"
)

type Heap []int

func (h Heap) Len() int              { return len(h) }
func (h Heap) Less(p, q int) bool    { return h[p] > h[q] }
func (h Heap) Swap(p, q int)         { h[p], h[q] = h[q], h[p] }
func (h *Heap) Push(v interface{})   { *h = append(*h, v.(int)) }
func (h *Heap) Pop() (v interface{}) { l := len(*h); v = (*h)[l-1]; *h = (*h)[:l-1]; return }

func mincostToHireWorkers(Q []int, W []int, k int) (ans float64) {
	n := len(Q)
	var qw [][2]int
	for i := 0; i < n; i++ {
		qw = append(qw, [2]int{Q[i], W[i]})
	}
	sort.Slice(qw, func(p, q int) bool { return qw[p][1]*qw[q][0] < qw[q][1]*qw[p][0] })
	h := &Heap{}
	tot := 0
	for i := 0; i < k; i++ {
		heap.Push(h, qw[i][0])
		tot += qw[i][0]
	}
	ans = float64(qw[k-1][1]) / float64(qw[k-1][0]) * float64(tot)
	for i := k; i < n; i++ {
		heap.Push(h, qw[i][0])
		tot += qw[i][0] - heap.Pop(h).(int)
		ans = math.Min(ans, float64(qw[i][1])/float64(qw[i][0])*float64(tot))
	}
	return
}
