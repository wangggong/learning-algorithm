/*
 * @lc app=leetcode.cn id=find-bottom-left-tree-value lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [513] 找树左下角的值
 *
 * https://leetcode-cn.com/problems/find-bottom-left-tree-value/description/
 *
 * algorithms
 * Medium (73.24%)
 * Total Accepted:    74.2K
 * Total Submissions: 101.2K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
 *
 * 假设二叉树中至少有一个节点。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: root = [2,1,3]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * ⁠
 *
 *
 * 输入: [1,2,3,4,null,5,6,null,null,7]
 * 输出: 7
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [1,104]
 * -1 << 31 <= Node.val <= 1 << 31 - 1
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
func findBottomLeftValue(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	var Q []*TreeNode
	Q = append(Q, root)
	for len(Q) > 0 {
		n := len(Q)
		for i := 0; i < n; i++ {
			node := Q[0]
			Q = Q[1:]
			if i == 0 {
				ans = node.Val
			}
			if node.Left != nil {
				Q = append(Q, node.Left)
			}
			if node.Right != nil {
				Q = append(Q, node.Right)
			}
		}
	}
	return ans
}
