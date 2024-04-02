/*
 * @lc app=leetcode.cn id=smallest-number-in-infinite-set lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2336] 无限集中的最小数字
 *
 * https://leetcode.cn/problems/smallest-number-in-infinite-set/description/
 *
 * algorithms
 * Medium (69.56%)
 * Total Accepted:    7.5K
 * Total Submissions: 10.7K
 * Testcase Example:  '["SmallestInfiniteSet","addBack","popSmallest","popSmallest","popSmallest","addBack","popSmallest","popSmallest","popSmallest"]\n' +
  '[[],[2],[],[],[],[1],[],[],[]]'
 *
 * 现有一个包含所有正整数的集合 [1, 2, 3, 4, 5, ...] 。
 *
 * 实现 SmallestInfiniteSet 类：
 *
 *
 * SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数。
 * int popSmallest() 移除 并返回该无限集中的最小整数。
 * void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集中。
 *
 *
 *
 *
 * 示例：
 *
 * 输入
 * ["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest",
 * "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
 * [[], [2], [], [], [], [1], [], [], []]
 * 输出
 * [null, null, 1, 2, 3, null, 1, 4, 5]
 *
 * 解释
 * SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
 * smallestInfiniteSet.addBack(2);    // 2 已经在集合中，所以不做任何变更。
 * smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 是最小的整数，并将其从集合中移除。
 * smallestInfiniteSet.popSmallest(); // 返回 2 ，并将其从集合中移除。
 * smallestInfiniteSet.popSmallest(); // 返回 3 ，并将其从集合中移除。
 * smallestInfiniteSet.addBack(1);    // 将 1 添加到该集合中。
 * smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 在上一步中被添加到集合中，
 * ⁠                                  // 且 1 是最小的整数，并将其从集合中移除。
 * smallestInfiniteSet.popSmallest(); // 返回 4 ，并将其从集合中移除。
 * smallestInfiniteSet.popSmallest(); // 返回 5 ，并将其从集合中移除。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 1000
 * 最多调用 popSmallest 和 addBack 方法 共计 1000 次
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
func (h *Heap) Pop() interface{} {
	l := len(h.IntSlice)
	v := h.IntSlice[l-1]
	h.IntSlice = h.IntSlice[:l-1]
	return v
}

type SmallestInfiniteSet struct {
	k      int
	heap   *Heap
	record map[int]struct{}
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		k:      1,
		heap:   &Heap{},
		record: make(map[int]struct{}),
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.heap.Len() == 0 {
		ans := s.k
		s.k++
		return ans
	}
	ans := heap.Pop(s.heap).(int)
	delete(s.record, ans)
	return ans
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := s.record[num]; s.k > num && !ok {
		heap.Push(s.heap, num)
		s.record[num] = struct{}{}
	}
	return
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
