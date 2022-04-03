/*
 * @lc app=leetcode.cn id=binary-tree-level-order-traversal lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (64.41%)
 * Total Accepted:    503.8K
 * Total Submissions: 781.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[9,20],[15,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
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

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var Q []*TreeNode
	Q = append(Q, root)
	lvl := 0
	for len(Q) > 0 {
		for size := len(Q); size > 0; size-- {
			q := Q[0]
			Q = Q[1:]
			if lvl >= len(ans) {
				ans = append(ans, nil)
			}
			ans[lvl] = append(ans[lvl], q.Val)
			if q.Left != nil {
				Q = append(Q, q.Left)
			}
			if q.Right != nil {
				Q = append(Q, q.Right)
			}
		}
		lvl++
	}
	return ans
}
