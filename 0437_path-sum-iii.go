/*
 * @lc app=leetcode.cn id=path-sum-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [437] 路径总和 III
 *
 * https://leetcode-cn.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Medium (56.97%)
 * Total Accepted:    162.4K
 * Total Submissions: 285.3K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
 *
 * 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
 * 输出：3
 * 解释：和等于 8 的路径有 3 条，如图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：3
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [0,1000]
 * -1e9 <= Node.val <= 1e9
 * -1000 <= targetSum <= 1000
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

func pathSum(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	return traversal(root, target, 0) + pathSum(root.Left, target) + pathSum(root.Right, target)
}

func traversal(root *TreeNode, target, curr int) int {
	ans := 0
	if root == nil {
		return ans
	}
	curr += root.Val
	if target == curr {
		ans++
	}
	ans += traversal(root.Left, target, curr)
	ans += traversal(root.Right, target, curr)
	return ans
}
