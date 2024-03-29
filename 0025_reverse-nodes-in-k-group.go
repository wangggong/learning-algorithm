/*
 * @lc app=leetcode.cn id=reverse-nodes-in-k-group lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (66.06%)
 * Total Accepted:    290.4K
 * Total Submissions: 438.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。
 *
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * 进阶：
 *
 *
 * 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 3
 * 输出：[3,2,1,4,5]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 1
 * 输出：[1,2,3,4,5]
 *
 *
 * 示例 4：
 *
 *
 * 输入：head = [1], k = 1
 * 输出：[1]
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 列表中节点的数量在范围 sz 内
 * 1 <= sz <= 5000
 * 0 <= Node.val <= 1000
 * 1 <= k <= sz
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	if k == 1 {
		return dummy.Next
	}
	prev, curr := dummy, dummy.Next
	for curr != nil {
		first, last := curr, curr
		for i := 0; i < k-1 && last != nil; i++ {
			last = last.Next
		}
		if last == nil {
			break
		}
		next := last.Next
		var p, c *ListNode
		c = first
		for p != last {
			n := c.Next
			c.Next = p
			p, c = c, n
		}
		// last.Next = first
		prev.Next = last
		first.Next = next
		prev, curr = first, next
	}
	return dummy.Next
}

/*
 * // 大逻辑拆小逻辑可能还好写点.
 * func reverseKGroup(head *ListNode, k int) *ListNode {
 * 	dummy := &ListNode{Val: -1, Next: head}
 * 	prev := dummy
 * 	for head != nil {
 * 		tail := head
 * 		for i := 1; i < k && tail != nil; i++ {
 * 			tail = tail.Next
 * 		}
 * 		if tail == nil {
 * 			break
 * 		}
 * 		next := tail.Next
 * 		reverse(head, tail)
 * 		prev.Next, head.Next = tail, next
 * 		prev, head = head, next
 * 	}
 * 	return dummy.Next
 * }
 * 
 * func reverse(head, tail *ListNode) {
 * 	var prev, curr, next *ListNode
 * 	for curr = head; prev != tail; prev, curr = curr, next {
 * 		next = curr.Next
 * 		curr.Next = prev
 * 	}
 * }
 */
