/*
 * @lc app=leetcode.cn id=maximum-level-sum-of-a-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1161] 最大层内元素和
 *
 * https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (62.86%)
 * Total Accepted:    19.9K
 * Total Submissions: 30.5K
 * Testcase Example:  '[1,7,0,7,-8,null,null]'
 *
 * 给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
 *
 * 请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,7,0,7,-8,null,null]
 * 输出：2
 * 解释：
 * 第 1 层各元素之和为 1，
 * 第 2 层各元素之和为 7 + 0 = 7，
 * 第 3 层各元素之和为 7 + -8 = -1，
 * 所以我们返回第 2 层的层号，它的层内元素之和最大。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在 [1, 10^4]范围内
 * -10^5 <= Node.val <= 10^5
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

func maxLevelSum(root *TreeNode) int {
	var Q []*TreeNode
	maxVal, maxLvl := math.MinInt32, 0
	Q = append(Q, root)
	for level := 1; len(Q) > 0; level++ {
		tot := 0
		for size := len(Q); size > 0; size-- {
			q := Q[0]
			Q = Q[1:]
			tot += q.Val
			if q.Left != nil {
				Q = append(Q, q.Left)
			}
			if q.Right != nil {
				Q = append(Q, q.Right)
			}
		}
		if tot > maxVal {
			maxVal, maxLvl = tot, level
		}
	}
	return maxLvl
}
