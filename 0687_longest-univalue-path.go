/*
 * @lc app=leetcode.cn id=longest-univalue-path lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [687] 最长同值路径
 *
 * https://leetcode-cn.com/problems/longest-univalue-path/description/
 *
 * algorithms
 * Medium (44.56%)
 * Total Accepted:    42.9K
 * Total Submissions: 96K
 * Testcase Example:  '[5,4,5,1,1,5]'
 *
 * 给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
 *
 * 两个节点之间的路径长度 由它们之间的边数表示。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入：root = [5,4,5,1,1,5]
 * 输出：2
 *
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入：root = [1,4,5,4,4,5]
 * 输出：2
 *
 *
 *
 *
 * 提示:
 *
 *
 * 树的节点数的范围是 [0, 10^4]
 * -1000 <= Node.val <= 1000
 * 树的深度将不超过 1000
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

var m map[*TreeNode]int
var ans int

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m = make(map[*TreeNode]int)
	ans = 0
	inorder(root)
	return ans - 1
}

func inorder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := m[root]; ok {
		return v
	}
	v := 1
	defer func() {
		ans = max(ans, v)
		m[root] = v
	}()
	left, right := root.Left, root.Right
	leftVal, rightVal := inorder(left), inorder(right)
	if leftVal != 0 && rightVal != 0 && root.Val == right.Val && root.Val == left.Val {
		v = max(leftVal, rightVal) + 1
		ans = max(ans, leftVal+rightVal+1)
		return v
	}
	if leftVal != 0 && root.Val == left.Val {
		v += leftVal
	}
	if rightVal != 0 && root.Val == right.Val {
		v += rightVal
	}
	return v
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
