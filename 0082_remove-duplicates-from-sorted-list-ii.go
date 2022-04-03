/*
 * @lc app=leetcode.cn id=remove-duplicates-from-sorted-list-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (53.24%)
 * Total Accepted:    223.8K
 * Total Submissions: 420K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,3,4,4,5]
 * 输出：[1,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,1,1,2,3]
 * 输出：[2,3]
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

/*
 * func deleteDuplicates(head *ListNode) *ListNode {
 * 	dummy := &ListNode{Val: -1, Next: head}
 * 	m := make(map[int]int)
 * 	for curr := head; curr != nil; curr = curr.Next {
 * 		m[curr.Val]++
 * 	}
 * 	prev, curr := dummy, head
 * 	for curr != nil {
 * 		if m[curr.Val] > 1 {
 * 			prev.Next = curr.Next
 * 			curr = curr.Next
 * 			continue
 * 		}
 * 		prev, curr = curr, curr.Next
 * 	}
 * 	return dummy.Next
 * }
 */

// 1. 舍得用变量;
// 2. dummy 技巧;
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	prev, curr := dummy, dummy.Next
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			v := curr.Val
			for curr != nil && curr.Val == v {
				prev.Next = curr.Next
				curr = curr.Next
			}
			continue
		}
		prev, curr = curr, curr.Next
	}
	return dummy.Next
}
