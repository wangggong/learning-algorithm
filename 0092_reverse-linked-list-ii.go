/*
 * @lc app=leetcode.cn id=reverse-linked-list-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (55.12%)
 * Total Accepted:    248.8K
 * Total Submissions: 451.4K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目为 n
 * 1 <= n <= 500
 * -500 <= Node.val <= 500
 * 1 <= left <= right <= n
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// assert 0 < left && left <= right && right <= n;
	dummy := &ListNode{-1, head}
	var prev, curr *ListNode
	curr = dummy
	for i := 0; i < left; i++ {
		if curr == nil {
			// unreachable
			panic("unreachable")
		}
		prev, curr = curr, curr.Next
	}
	p := prev
	l := curr
	for i := left; i <= right; i++ {
		if curr == nil {
			// unreachable
			panic("unreachable")
		}
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	p.Next = prev
	l.Next = curr
	// prev.Next = _reverse(prev, curr, right-left)
	return dummy.Next
}

