/*
 * @lc app=leetcode.cn id=sum-of-left-leaves lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (59.40%)
 * Total Accepted:    132.1K
 * Total Submissions: 221K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定二叉树的根节点 root ，返回所有左叶子之和。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入: root = [3,9,20,null,null,15,7]
 * 输出: 24
 * 解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 *
 *
 * 示例 2:
 *
 *
 * 输入: root = [1]
 * 输出: 0
 *
 *
 *
 *
 * 提示:
 *
 *
 * 节点数在 [1, 1000] 范围内
 * -1000 <= Node.val <= 1000
 *
 *
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

const (
	_default = iota
	_left
	_right
)

var ans int

func sumOfLeftLeaves(root *TreeNode) int {
	ans = 0
	inorder(root, _default)
	return ans
}

func inorder(root *TreeNode, pos int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if pos == _left {
			ans += root.Val
		}
		return
	}
	inorder(root.Left, _left)
	inorder(root.Right, _right)
}
