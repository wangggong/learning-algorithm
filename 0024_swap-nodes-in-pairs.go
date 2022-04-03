/*
 * @lc app=leetcode.cn id=swap-nodes-in-pairs lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (70.63%)
 * Total Accepted:    397.6K
 * Total Submissions: 562.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：head = [1]
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目在范围 [0, 100] 内
 * 0 <= Node.val <= 100
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
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	prev, curr := dummy, dummy.Next
	for curr != nil {
		first, last := curr, curr
		for i := 0; i < 1 && last != nil; i++ {
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
