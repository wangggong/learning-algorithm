/*
 * @lc app=leetcode.cn id=insert-into-a-binary-search-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [701] 二叉搜索树中的插入操作
 *
 * https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (72.13%)
 * Total Accepted:    106.2K
 * Total Submissions: 147.4K
 * Testcase Example:  '[4,2,7,1,3]\n5'
 *
 * 给定二叉搜索树（BST）的根节点 root 和要插入树中的值 value ，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证
 * ，新值和原始二叉搜索树中的任意节点值都不同。
 *
 * 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [4,2,7,1,3], val = 5
 * 输出：[4,2,7,1,3,5]
 * 解释：另一个满足题目要求可以通过的树是：
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [40,20,60,10,30,50,70], val = 25
 * 输出：[40,20,60,10,30,50,70,null,null,25]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
 * 输出：[4,2,7,1,3,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数将在 [0, 10^4]的范围内。
 * -10^8 <= Node.val <= 10^8
 * 所有值 Node.val 是 独一无二 的。
 * -10^8 <= val <= 10^8
 * 保证 val 在原始BST中不存在。
 *
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
 * func insertIntoBST(root *TreeNode, val int) *TreeNode {
 * 	node := &TreeNode{Val: val}
 * 	if root == nil {
 * 		return node
 * 	}
 * 	// [5,null,14,10,77,null,null,null,95,null,null]\n4
 * 	prev, curr := root, root
 * 	for curr != nil && (curr.Left != nil || curr.Right != nil) {
 * 		if curr.Val > val {
 * 			prev, curr = curr, curr.Left
 * 		} else {
 * 			prev, curr = curr, curr.Right
 * 		}
 * 	}
 * 	if curr == nil {
 * 		if prev.Val > val {
 * 			prev.Left = node
 * 		} else {
 * 			prev.Right = node
 * 		}
 * 		return root
 * 	}
 * 	if curr.Val > val {
 * 		curr.Left = node
 * 	} else {
 * 		curr.Right = node
 * 	}
 * 	return root
 * }
 */

/*
 * func insertIntoBST(root *TreeNode, val int) *TreeNode {
 * 	if root == nil {
 * 		return &TreeNode{Val: val}
 * 	}
 * 	// assert root.Val != val;
 * 	if root.Val > val {
 * 		root.Left = insertIntoBST(root.Left, val)
 * 	} else {
 * 		root.Right = insertIntoBST(root.Right, val)
 * 	}
 * 	return root
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if root == nil {
		return newNode
	}
	for curr := root; curr != nil; {
		// assert curr.Val != val;
		if curr.Val > val {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = newNode
				break
			}
		} else {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = newNode
				break
			}
		}
	}
	return root
}
