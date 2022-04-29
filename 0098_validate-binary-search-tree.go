import "math"

/*
 * @lc app=leetcode.cn id=validate-binary-search-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (35.52%)
 * Total Accepted:    451.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '[2,1,3]'
 *
 * 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 *
 * 有效 二叉搜索树定义如下：
 *
 *
 * 节点的左子树只包含 小于 当前节点的数。
 * 节点的右子树只包含 大于 当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,1,4,null,null,3,6]
 * 输出：false
 * 解释：根节点的值是 5 ，但是右子节点的值是 4 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在[1, 10^4] 内
 * -2^31 <= Node.val <= 2^31 - 1
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
 * func isValidBST(root *TreeNode) bool {
 * 	return inorder(root, math.MinInt64, math.MaxInt64)
 * }
 *
 * func inorder(root *TreeNode, min, max int) bool {
 * 	if root == nil {
 * 		return true
 * 	}
 * 	if root.Val <= min || root.Val >= max {
 * 		return false
 * 	}
 * 	return inorder(root.Left, min, root.Val) && inorder(root.Right, root.Val, max)
 * }
 */

/*
 * func isValidBST(root *TreeNode) bool {
 * 	return inorder(root, nil, nil)
 * }
 *
 * func inorder(root, min, max *TreeNode) bool {
 * 	if root == nil {
 * 		return true
 * 	}
 * 	if min != nil && root.Val <= min.Val {
 * 		return false
 * 	}
 * 	if max != nil && root.Val >= max.Val {
 * 		return false
 * 	}
 * 	return inorder(root.Left, min, root) && inorder(root.Right, root, max)
 * }
 */

// 发现了好俊的一发前序遍历!
var preVal int

func isValidBST(root *TreeNode) bool {
	preVal = math.MinInt64
	return preOrder(root)
}

func preOrder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !preOrder(root.Left) {
		return false
	}
	if root.Val <= preVal {
		return false
	}
	preVal = root.Val
	return preOrder(root.Right)
}
