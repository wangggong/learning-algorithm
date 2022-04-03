/*
 * @lc app=leetcode.cn id=remove-duplicates-from-sorted-list lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (53.67%)
 * Total Accepted:    394.4K
 * Total Submissions: 734.9K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,1,2]
 * 输出：[1,2]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [1,1,2,3,3]
 * 输出：[1,2,3]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点数目在范围 [0, 300] 内
 * -100 <= Node.val <= 100
 * 题目数据保证链表已经按升序 排列
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
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	prev, curr := dummy, head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			prev = curr
			for curr != nil && curr.Val == prev.Val {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev, curr = curr, curr.Next
		}
	}
	return dummy.Next
}
