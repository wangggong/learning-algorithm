/*
 * @lc app=leetcode.cn id=reorder-list lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [143] 重排链表
 *
 * https://leetcode-cn.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (63.44%)
 * Total Accepted:    170.4K
 * Total Submissions: 268.4K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * 请将其重新排列后变为：
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[1,4,2,3]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[1,5,2,4,3]
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 5 * 10^4]
 * 1 <= node.val <= 1000
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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	h1, h2 := split(head)
	h2 = revert(h2)
	head = merge(h1, h2)
	return
}

func split(head *ListNode) (*ListNode, *ListNode) {
	p, q := head, head.Next
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	newHead := p.Next
	p.Next = nil
	return head, newHead
}

func revert(head *ListNode) *ListNode {
	prev, curr := (*ListNode)(nil), head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}

func merge(h1, h2 *ListNode) *ListNode {
	head := h1
	h1 = h1.Next
	curr := head
	mergeNewHead := true
	for h1 != nil && h2 != nil {
		if mergeNewHead {
			curr.Next = h2
			h2 = h2.Next
		} else {
			curr.Next = h1
			h1 = h1.Next
		}
		curr = curr.Next
		mergeNewHead = !mergeNewHead
	}
	if h1 != nil {
		curr.Next = h1
	} else {
		curr.Next = h2
	}
	return head
}
