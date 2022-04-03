/*
 * Medium
 *
 * 给定一个二叉树，统计该二叉树数值相同的子树个数。
 *
 * 同值子树是指该子树的所有节点都拥有相同的数值。
 *
 * 示例：
 *
 * 输入: root = [5,1,5,5,5,null,5]
 *
 *               5
 *              / \
 *             1   5
 *            / \   \
 *           5   5   5
 *
 * 输出: 4
 * 通过次数5,272
 * 提交次数8,303
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var count int

func countUnivalSubtrees(root *TreeNode) int {
	count = 0
	dfs(root)
	return count
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return false
	}
	left, right := true, true
	if root.Left != nil {
		left = dfs(root.Left)
		left = left && root.Val == root.Left.Val
	}
	if root.Right != nil {
		right = dfs(root.Right)
		right = right && root.Val == root.Right.Val
	}
	ans := left && right
	if ans {
		count++
	}
	return ans
}
