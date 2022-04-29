/*
 * @lc app=leetcode.cn id=binary-tree-postorder-traversal lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (75.67%)
 * Total Accepted:    419.4K
 * Total Submissions: 554.3K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[3,2,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	visit := make(map[*TreeNode]struct{})
	var S []*TreeNode
	S = append(S, root)
	for len(S) > 0 {
		node := S[len(S)-1]
		S = S[:len(S)-1]
		if _, ok := visit[node]; ok {
			ans = append(ans, node.Val)
			continue
		}
		visit[node] = struct{}{}
		S = append(S, node)
		if node.Right != nil {
			S = append(S, node.Right)
		}
		if node.Left != nil {
			S = append(S, node.Left)
		}
	}
	return ans
}
