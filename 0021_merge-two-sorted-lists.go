/*
 * @lc app=leetcode.cn id=merge-two-sorted-lists lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (66.69%)
 * Total Accepted:    921.6K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：l1 = [1,2,4], l2 = [1,3,4]
 * 输出：[1,1,2,3,4,4]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：l1 = [], l2 = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：l1 = [], l2 = [0]
 * 输出：[0]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 两个链表的节点数目范围是 [0, 50]
 * -100 <= Node.val <= 100
 * l1 和 l2 均按 非递减顺序 排列
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
			curr = curr.Next
			curr.Next = nil
		} else {
			curr.Next = list2
			list2 = list2.Next
			curr = curr.Next
			curr.Next = nil
		}
	}
	if list1 == nil {
		curr.Next = list2
	} else {
		curr.Next = list1
	}
	return dummy.Next
}
