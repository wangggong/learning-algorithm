/*
 * @lc app=leetcode.cn id=construct-binary-tree-from-inorder-and-postorder-traversal lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.30%)
 * Total Accepted:    172.2K
 * Total Submissions: 238.1K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder
 * 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * 输出：[3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入：inorder = [-1], postorder = [-1]
 * 输出：[-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= inorder.length <= 3000
 * postorder.length == inorder.length
 * -3000 <= inorder[i], postorder[i] <= 3000
 * inorder 和 postorder 都由 不同 的值组成
 * postorder 中每一个值都在 inorder 中
 * inorder 保证是树的中序遍历
 * postorder 保证是树的后序遍历
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// assert len(inorder) == len(postorder) && len(postorder) > 0;
	return dfs(inorder, postorder)
}

func dfs(in, post []int) *TreeNode {
	n := len(in)
	// assert n == len(post);
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: post[n-1]}
	if n == 1 {
		return root
	}
	var k int
	for k = 0; k < n; k++ {
		if in[k] == post[n-1] {
			break
		}
	}
	root.Left = dfs(in[:k], post[:k])
	root.Right = dfs(in[k+1:], post[k:n-1])
	return root
}
