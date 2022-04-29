/*
 * @lc app=leetcode.cn id=lowest-common-ancestor-of-a-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [236] 二叉树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (68.55%)
 * Total Accepted:    321.4K
 * Total Submissions: 468.5K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * 输出：3
 * 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * 输出：5
 * 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], p = 1, q = 2
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [2, 1e5] 内。
 * -1e9 <= Node.val <= 1e9
 * 所有 Node.val 互不相同 。
 * p != q
 * p 和 q 均存在于给定的二叉树中。
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// assert there `is` a root.
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// assert (pl || pr) && (ql || qr);
	pl, pr := inSubTree(root.Left, p), inSubTree(root.Right, p)
	ql, qr := inSubTree(root.Left, q), inSubTree(root.Right, q)
	if pl && ql {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if pr && qr {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func inSubTree(root, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val == node.Val {
		return true
	}
	return inSubTree(root.Left, node) || inSubTree(root.Right, node)
}

/*
 * // 其实也不用判定右子树的...
 * func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
 * 	if root == nil || p == nil || q == nil {
 * 		return nil
 * 	}
 * 	if p == root || q == root {
 * 		return root
 * 	}
 * 	pLeft, qLeft := inSubTree(root.Left, p), inSubTree(root.Left, q)
 * 	if pLeft != qLeft {
 * 		return root
 * 	}
 * 	switch pLeft {
 * 	case true:
 * 		return lowestCommonAncestor(root.Left, p, q)
 * 	case false:
 * 		return lowestCommonAncestor(root.Right, p, q)
 * 	default:
 * 		// unreachable
 * 	}
 * 	return nil
 * }
 * 
 * func inSubTree(root, node *TreeNode) bool {
 * 	// assert node != nil;
 * 	if root == nil {
 * 		return false
 * 	}
 * 	return root == node || inSubTree(root.Left, node) || inSubTree(root.Right, node)
 * }
 */
