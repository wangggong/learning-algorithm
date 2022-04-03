/*
 * @lc app=leetcode.cn id=delete-node-in-a-bst lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [450] 删除二叉搜索树中的节点
 *
 * https://leetcode-cn.com/problems/delete-node-in-a-bst/description/
 *
 * algorithms
 * Medium (49.68%)
 * Total Accepted:    90.9K
 * Total Submissions: 181.5K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n3'
 *
 * 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key
 * 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
 *
 * 一般来说，删除节点可分为两个步骤：
 *
 *
 * 首先找到需要删除的节点；
 * 如果找到了，删除它。
 *
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入：root = [5,3,6,2,4,null,7], key = 3
 * 输出：[5,4,6,2,null,null,7]
 * 解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
 * 一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
 * 另一个正确答案是 [5,2,6,null,4,null,7]。
 *
 *
 *
 *
 * 示例 2:
 *
 *
 * 输入: root = [5,3,6,2,4,null,7], key = 0
 * 输出: [5,3,6,2,4,null,7]
 * 解释: 二叉树不包含值为 0 的节点
 *
 *
 * 示例 3:
 *
 *
 * 输入: root = [], key = 0
 * 输出: []
 *
 *
 *
 * 提示:
 *
 *
 * 节点数的范围 [0, 10^4].
 * -10^5 <= Node.val <= 10^5
 * 节点值唯一
 * root 是合法的二叉搜索树
 * -10^5 <= key <= 10^5
 *
 *
 *
 *
 * 进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。
 *
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
 * func deleteNode(root *TreeNode, key int) *TreeNode {
 * 	if root == nil {
 * 		return root
 * 	}
 * 	if root.Val == key {
 * 		if root.Left == nil {
 * 			return root.Right
 * 		}
 * 		if root.Right == nil {
 * 			return root.Left
 * 		}
 * 		var prev, curr *TreeNode
 * 		curr = root.Right
 * 		for curr.Left != nil {
 * 			prev, curr = curr, curr.Left
 * 		}
 * 		if prev != nil {
 * 			prev.Left = curr.Right
 * 			curr.Right = root.Right
 * 		}
 * 		curr.Left = root.Left
 * 		return curr
 * 	}
 * 	if root.Val < key {
 * 		root.Right = deleteNode(root.Right, key)
 * 		return root
 * 	}
 * 	if root.Val > key {
 * 		root.Left = deleteNode(root.Left, key)
 * 		return root
 * 	}
 * 	// unreachable
 * 	return nil
 * }
 */

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	left, right := root.Left, root.Right
	root.Left, root.Right = nil, nil
	var prev, curr *TreeNode
	for curr = right; curr != nil; prev, curr = curr, curr.Left {
	}
	prev.Left = left
	return right
}
