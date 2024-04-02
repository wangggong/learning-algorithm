/*
 * @lc app=leetcode.cn id=find-largest-value-in-each-tree-row lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (65.78%)
 * Total Accepted:    70.3K
 * Total Submissions: 106.4K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
 *
 *
 *
 * 示例1：
 *
 *
 *
 *
 * 输入: root = [1,3,2,5,3,null,9]
 * 输出: [1,3,9]
 *
 *
 * 示例2：
 *
 *
 * 输入: root = [1,2,3]
 * 输出: [1,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 二叉树的节点个数的范围是 [0,10^4]
 * -2^31 <= Node.val <= 2^31 - 1
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

import "math"

func largestValues(root *TreeNode) []int {
	var Q []*TreeNode
	var ans []int
	if root == nil {
		return ans
	}
	Q = append(Q, root)
	for len(Q) > 0 {
		mx := math.MinInt32
		for size := len(Q); size > 0; size-- {
			q := Q[0]
			Q = Q[1:]
			if q.Val > mx {
				mx = q.Val
			}
			if q.Left != nil {
				Q = append(Q, q.Left)
			}
			if q.Right != nil {
				Q = append(Q, q.Right)
			}
		}
		ans = append(ans, mx)
	}
	return ans
}

/*
 * var ans []int
 * 
 * func dfs(root *TreeNode, depth int) {
 * 	if root == nil {
 * 		return
 * 	}
 * 	if len(ans) <= depth {
 * 		ans = append(ans, math.MinInt32)
 * 	}
 * 	if root.Val > ans[depth] {
 * 		ans[depth] = root.Val
 * 	}
 * 	dfs(root.Left, depth+1)
 * 	dfs(root.Right, depth+1)
 * 	return
 * }
 * 
 * func largestValues(root *TreeNode) []int {
 * 	ans = nil
 * 	dfs(root, 0)
 * 	return ans
 * }
 */
