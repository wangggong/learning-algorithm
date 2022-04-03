/*
 * @lc app=leetcode.cn id=two-sum-iv-input-is-a-bst lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [653] 两数之和 IV - 输入 BST
 *
 * https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (60.38%)
 * Total Accepted:    51.9K
 * Total Submissions: 85.9K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * 给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: root = [5,3,6,2,4,null,7], k = 9
 * 输出: true
 *
 *
 * 示例 2：
 *
 *
 * 输入: root = [5,3,6,2,4,null,7], k = 28
 * 输出: false
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是  [1, 10^4].
 * -10^4 <= Node.val <= 10^4
 * root 为二叉搜索树
 * -10^5 <= k <= 10^5
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

var m map[int]struct{}

func findTarget(root *TreeNode, k int) bool {
	m = make(map[int]struct{})
	return dfs(root, k)
}

func dfs(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = struct{}{}
	return dfs(root.Left, k) || dfs(root.Right, k)
}
