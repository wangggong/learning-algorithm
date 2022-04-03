/*
 * @lc app=leetcode.cn id=reverse-linked-list lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (72.53%)
 * Total Accepted:    834.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[5,4,3,2,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：[2,1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目范围是 [0, 5000]
 * -5000
 *
 *
 *
 *
 * 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
 *
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
// func reverseList(head *ListNode) *ListNode {
// 	var p, c *ListNode
// 	c = head
// 	for c != nil {
// 		t := c.Next
// 		c.Next = p
// 		p, c = c, t
// 	}
// 	return p
// }

func _reverseList(head, prev *ListNode) *ListNode {
	if head == nil {
		return prev
	}
	next := head.Next
	head.Next = prev
	return _reverseList(next, head)
}

func reverseList(head *ListNode) *ListNode {
	return _reverseList(head, nil)
}
