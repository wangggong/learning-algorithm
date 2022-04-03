/*
 * @lc app=leetcode.cn id=palindrome-linked-list lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (50.71%)
 * Total Accepted:    381.4K
 * Total Submissions: 749.1K
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	p := leftHalf(head)
	q := reverse(p.Next)
	// 此时找到了两个 head, 偶数时前一个大 1, 奇数时大小相等.
	for p, q := head, q; p != nil && q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			return false
		}
	}
	return true
}

// 找到左半边头节点, 快慢指针.
func leftHalf(head *ListNode) *ListNode {
	p, q := head, head
	for q.Next != nil && q.Next.Next != nil {
		p, q = p.Next, q.Next.Next
	}
	return p
}

func reverse(head *ListNode) *ListNode {
	var prev, curr, next *ListNode
	for curr = head; curr != nil; prev, curr = curr, next {
		next = curr.Next
		curr.Next = prev
	}
	return prev
}
