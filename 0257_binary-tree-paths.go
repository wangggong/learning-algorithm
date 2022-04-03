/*
 * @lc app=leetcode.cn id=binary-tree-paths lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (69.00%)
 * Total Accepted:    182.8K
 * Total Submissions: 264.3K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,null,5]
 * 输出：["1->2->5","1->3"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：["1"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 */

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans []string

/*
 * func binaryTreePaths(root *TreeNode) []string {
 * 	ans = nil
 * 	if root == nil {
 * 		return ans
 * 	}
 * 	preOrder(root, "")
 * 	return ans
 * }
 */

func preOrder(root *TreeNode, cur string) {
	// assert root != nil;
	cur = cur + "->" + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		ans = append(ans, cur[2:])
		return
	}
	if root.Left != nil {
		preOrder(root.Left, cur)
	}
	if root.Right != nil {
		preOrder(root.Right, cur)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	return dfs(root)
}

func dfs(root *TreeNode) []string {
	s := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{s}
	}
	var ans []string
	prefix := s + "->"
	if root.Left != nil {
		left := dfs(root.Left)
		for _, l := range left {
			ans = append(ans, prefix+l)
		}
	}
	if root.Right != nil {
		right := dfs(root.Right)
		for _, r := range right {
			ans = append(ans, prefix+r)
		}
	}
	return ans
}
