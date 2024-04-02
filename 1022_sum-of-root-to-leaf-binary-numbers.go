/*
 * @lc app=leetcode.cn id=sum-of-root-to-leaf-binary-numbers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1022] 从根到叶的二进制数之和
 *
 * https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/description/
 *
 * algorithms
 * Easy (71.45%)
 * Total Accepted:    29.7K
 * Total Submissions: 40.7K
 * Testcase Example:  '[1,0,1,0,1,0,1]'
 *
 * 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
 *
 *
 * 例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
 *
 *
 * 对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
 *
 * 返回这些数字之和。题目数据保证答案是一个 32 位 整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,0,1,0,1,0,1]
 * 输出：22
 * 解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [0]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在 [1, 1000] 范围内
 * Node.val 仅为 0 或 1
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

var ans int

func sumRootToLeaf(root *TreeNode) int {
	ans = 0
	preOrder(root, 0)
	return ans
}

func preOrder(root *TreeNode, val int) {
	if root == nil {
		return
	}
	val = (val << 1) | root.Val
	if root.Left == nil && root.Right == nil {
		// fmt.Println(val)
		ans += val
		return
	}
	if root.Left != nil {
		preOrder(root.Left, val)
	}
	if root.Right != nil {
		preOrder(root.Right, val)
	}
	return
}
