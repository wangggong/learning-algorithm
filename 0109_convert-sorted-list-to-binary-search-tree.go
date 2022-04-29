/*
 * @lc app=leetcode.cn id=convert-sorted-list-to-binary-search-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (76.21%)
 * Total Accepted:    113.9K
 * Total Submissions: 149.5K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: head = [-10,-3,0,5,9]
 * 输出: [0,-3,9,-10,null,5]
 * 解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
 *
 *
 * 示例 2:
 *
 *
 * 输入: head = []
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * head 中的节点数在[0, 2 * 10^4] 范围内
 * -10^5 <= Node.val <= 10^5
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
 * // 快慢指针.
 * func sortedListToBST(head *ListNode) *TreeNode {
 * 	if head == nil {
 * 		return nil
 * 	}
 * 	if head.Next == nil {
 * 		return &TreeNode{Val: head.Val}
 * 	}
 * 	p, q := head, head.Next
 * 	for q.Next != nil && q.Next.Next != nil {
 * 		p, q = p.Next, q.Next.Next
 * 	}
 * 	node := p.Next
 * 	p.Next = nil
 * 	return &TreeNode{
 * 		Val:   node.Val,
 * 		Left:  sortedListToBST(head),
 * 		Right: sortedListToBST(node.Next),
 * 	}
 * }
 */

var node *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	size := 0
	for curr := head; curr != nil; curr = curr.Next {
		size++
	}
	node = head
	return inorder(0, size)
}

func inorder(l, r int) *TreeNode {
	if l >= r {
		return nil
	}
	mid := (l + r) >> 1
	left := inorder(l, mid)
	val := node.Val
	node = node.Next
	right := inorder(mid+1, r)
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}
