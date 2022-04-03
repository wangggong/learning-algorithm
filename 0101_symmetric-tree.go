/*
 * @lc app=leetcode.cn id=symmetric-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (57.06%)
 * Total Accepted:    527.5K
 * Total Submissions: 920.8K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给你一个二叉树的根节点 root ， 检查它是否轴对称。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,2,3,4,4,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,null,3,null,3]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [1, 1000] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
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
 * func isSymmetric(root *TreeNode) bool {
 * 	if root == nil {
 * 		return true
 * 	}
 * 	return symmetric(root.Left, root.Right)
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var S [][2]*TreeNode
	S = append(S, [2]*TreeNode{root.Left, root.Right})
	for len(S) > 0 {
		v := S[len(S)-1]
		S = S[:len(S)-1]
		left, right := v[0], v[1]
		if (left != nil) != (right != nil) {
			return false
		}
		if left == nil {
			continue
		}
		if left.Val != right.Val {
			return false
		}
		S = append(S, [2]*TreeNode{left.Left, right.Right})
		S = append(S, [2]*TreeNode{left.Right, right.Left})
	}
	return true
}

func symmetric(left, right *TreeNode) bool {
	if (left != nil) != (right != nil) {
		return false
	}
	if left == nil {
		return true
	}
	if left.Val != right.Val {
		return false
	}
	return symmetric(left.Right, right.Left) && symmetric(left.Left, right.Right)
}
