/*
 * @lc app=leetcode.cn id=find-servers-that-handled-most-number-of-requests lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1606] 找到处理最多请求的服务器
 *
 * https://leetcode-cn.com/problems/find-servers-that-handled-most-number-of-requests/description/
 *
 * algorithms
 * Hard (35.96%)
 * Total Accepted:    6.7K
 * Total Submissions: 14.9K
 * Testcase Example:  '3\n[1,2,3,4,5]\n[5,2,3,3,3]'
 *
 * 你有 k 个服务器，编号为 0 到 k-1 ，它们可以同时处理多个请求组。每个服务器有无穷的计算能力但是 不能同时处理超过一个请求
 * 。请求分配到服务器的规则如下：
 *
 *
 * 第 i （序号从 0 开始）个请求到达。
 * 如果所有服务器都已被占据，那么该请求被舍弃（完全不处理）。
 * 如果第 (i % k) 个服务器空闲，那么对应服务器会处理该请求。
 * 否则，将请求安排给下一个空闲的服务器（服务器构成一个环，必要的话可能从第 0 个服务器开始继续找下一个空闲的服务器）。比方说，如果第 i
 * 个服务器在忙，那么会查看第 (i+1) 个服务器，第 (i+2) 个服务器等等。
 *
 *
 * 给你一个 严格递增 的正整数数组 arrival ，表示第 i 个任务的到达时间，和另一个数组 load ，其中 load[i] 表示第 i
 * 个请求的工作量（也就是服务器完成它所需要的时间）。你的任务是找到 最繁忙的服务器 。最繁忙定义为一个服务器处理的请求数是所有服务器里最多的。
 *
 * 请你返回包含所有 最繁忙服务器 序号的列表，你可以以任意顺序返回这个列表。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3]
 * 输出：[1]
 * 解释：
 * 所有服务器一开始都是空闲的。
 * 前 3 个请求分别由前 3 台服务器依次处理。
 * 请求 3 进来的时候，服务器 0 被占据，所以它呗安排到下一台空闲的服务器，也就是服务器 1 。
 * 请求 4 进来的时候，由于所有服务器都被占据，该请求被舍弃。
 * 服务器 0 和 2 分别都处理了一个请求，服务器 1 处理了两个请求。所以服务器 1 是最忙的服务器。
 *
 *
 * 示例 2：
 *
 *
 * 输入：k = 3, arrival = [1,2,3,4], load = [1,2,1,2]
 * 输出：[0]
 * 解释：
 * 前 3 个请求分别被前 3 个服务器处理。
 * 请求 3 进来，由于服务器 0 空闲，它被服务器 0 处理。
 * 服务器 0 处理了两个请求，服务器 1 和 2 分别处理了一个请求。所以服务器 0 是最忙的服务器。
 *
 *
 * 示例 3：
 *
 *
 * 输入：k = 3, arrival = [1,2,3], load = [10,12,11]
 * 输出：[0,1,2]
 * 解释：每个服务器分别处理了一个请求，所以它们都是最忙的服务器。
 *
 *
 * 示例 4：
 *
 *
 * 输入：k = 3, arrival = [1,2,3,4,8,9,10], load = [5,2,10,3,1,2,2]
 * 输出：[1]
 *
 *
 * 示例 5：
 *
 *
 * 输入：k = 1, arrival = [1], load = [1]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= 1e5
 * 1 <= arrival.length, load.length <= 1e5
 * arrival.length == load.length
 * 1 <= arrival[i], load[i] <= 1e9
 * arrival 保证 严格递增 。
 *
 *
 */

/*
 * // 太卷了, 拉倒吧. 抄答案了.
 * import (
 * 	"container/heap"
 * )
 *
 * const (
 * 	Index = iota
 * 	EndTime
 * )
 *
 * type Heap struct {
 * 	Arr [][2]int
 * }
 *
 * func (h Heap) Len() int              { return len(h.Arr) }
 * func (h Heap) Less(p, q int) bool    { return h.Arr[p][EndTime] < h.Arr[q][EndTime] }
 * func (h Heap) Swap(p, q int)         { h.Arr[p], h.Arr[q] = h.Arr[q], h.Arr[p] }
 * func (h *Heap) Push(v interface{})   { h.Arr = append(h.Arr, v.([2]int)) }
 * func (h *Heap) Pop() (v interface{}) { l := len(h.Arr); v = h.Arr[l-1]; h.Arr = h.Arr[:l-1]; return v }
 *
 * func busiestServers(k int, A []int, L []int) []int {
 * 	n := len(A)
 * 	// assert len(A) == len(L);
 * 	h := &Heap{}
 * 	cnt := make([]int, k)
 * 	for i := 0; i < k && i < n; i++ {
 * 		heap.Push(h, [2]int{i, A[i] + L[i]})
 * 		cnt[i]++
 * 	}
 * 	for i := k; i < n; i++ {
 * 		var cur [][2]int
 * 		ind := 0
 * 		for h.Len() > 0 {
 * 			q := heap.Pop(h).([2]int)
 * 			if q[EndTime] > A[i] {
 * 				heap.Push(h, q)
 * 				break
 * 			}
 * 			cur = append(cur, q)
 * 			if (cur[ind][Index]+k-(i%k))%k > (cur[len(cur)-1][Index]+k-(i%k))%k {
 * 				ind = len(cur) - 1
 * 			}
 * 		}
 * 		if len(cur) == 0 {
 * 			continue
 * 		}
 * 		q := cur[ind]
 * 		cnt[q[Index]]++
 * 		q[EndTime] = A[i] + L[i]
 * 		cur[ind] = q
 * 		for _, c := range cur {
 * 			heap.Push(h, c)
 * 		}
 * 	}
 * 	mxCnt := 0
 * 	for _, c := range cnt {
 * 		if c > mxCnt {
 * 			mxCnt = c
 * 		}
 * 	}
 * 	var ans []int
 * 	for i, c := range cnt {
 * 		if c == mxCnt {
 * 			ans = append(ans, i)
 * 		}
 * 	}
 * 	return ans
 * }
 */

func busiestServers(k int, arrival, load []int) (ans []int) {
	available := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		available.Put(i, nil)
	}
	busy := hp{}
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			available.Put(busy[0].id, nil)
			heap.Pop(&busy)
		}
		if available.Size() == 0 {
			continue
		}
		node, _ := available.Ceiling(i % k)
		if node == nil {
			node = available.Left()
		}
		id := node.Key.(int)
		requests[id]++
		if requests[id] > maxRequest {
			maxRequest = requests[id]
			ans = []int{id}
		} else if requests[id] == maxRequest {
			ans = append(ans, id)
		}
		heap.Push(&busy, pair{t + load[i], id})
		available.Remove(id)
	}
	return
}

type pair struct{ end, id int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
