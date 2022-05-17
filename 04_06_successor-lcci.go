/*
 * 设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
 *
 * 如果指定节点没有对应的“下一个”节点，则返回null。
 *
 * 示例 1:
 *
 * 输入: root = [2,1,3], p = 1
 *
 *   2
 *  / \
 * 1   3
 *
 * 输出: 2
 * 示例 2:
 *
 * 输入: root = [5,3,6,2,4,null,null,1], p = 6
 *
 *       5
 *      / \
 *     3   6
 *    / \
 *   2   4
 *  /
 * 1
 *
 * 输出: null
 * 通过次数30,273
 * 提交次数50,477
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var prev, ans *TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	prev, ans = nil, nil
	inorder(root, p)
	return ans
}

func inorder(root, p *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, p)
	if prev == p {
		ans = root
	}
	prev = root
	inorder(root.Right, p)
	return
}