/*
 * @lc app=leetcode.cn id=merge-k-sorted-lists lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [23] 合并K个升序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (56.81%)
 * Total Accepted:    445.4K
 * Total Submissions: 783.9K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Heap []*ListNode

func (h Heap) Len() int              { return len(h) }
func (h Heap) Less(p, q int) bool    { return h[p].Val < h[q].Val }
func (h Heap) Swap(p, q int)         { h[p], h[q] = h[q], h[p] }
func (h *Heap) Push(v interface{})   { *h = append(*h, v.(*ListNode)) }
func (h *Heap) Pop() (v interface{}) { l := len(*h); v = (*h)[l-1]; *h = (*h)[:l-1]; return v }

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	h := &Heap{}
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}
	for curr := dummy; h.Len() > 0; curr = curr.Next {
		node := heap.Pop(h).(*ListNode)
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		curr.Next = node
	}
	return dummy.Next
}
