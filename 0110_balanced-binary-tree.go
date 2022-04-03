/*
 * @lc app=leetcode.cn id=balanced-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (56.65%)
 * Total Accepted:    320.9K
 * Total Submissions: 565.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 *
 * 本题中，一棵高度平衡二叉树定义为：
 *
 *
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,3,3,null,null,4,4]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在范围 [0, 5000] 内
 * -1e4 <= Node.val <= 1e4
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
func isBalanced(root *TreeNode) bool {
	_, ok := balanced(root)
	return ok
}

func balanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, ok := balanced(root.Left)
	if !ok {
		return 0, false
	}
	rightDepth, ok := balanced(root.Right)
	if !ok {
		return 0, false
	}
	if leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {
		return 0, false
	}
	depth := leftDepth
	if rightDepth > depth {
		depth = rightDepth
	}
	return depth + 1, true
}
