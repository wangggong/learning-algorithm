/*
 * @lc app=leetcode.cn id=path-sum-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (62.98%)
 * Total Accepted:    226.5K
 * Total Submissions: 359.3K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：[[5,4,11,2],[5,8,4,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], targetSum = 0
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点总数在范围 [0, 5000] 内
 * -1000 <= Node.val <= 1000
 * -1000 <= targetSum <= 1000
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

var ans [][]int
var cur []int
var curSum int

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans, cur, curSum = nil, nil, 0
	if root == nil {
		return ans
	}
	preOrder(root, targetSum)
	return ans
}

func preOrder(root *TreeNode, target int) {
	if root.Left == nil && root.Right == nil {
		cur = append(cur, root.Val)
		curSum += root.Val
		if curSum == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
		}
		cur = cur[:len(cur)-1]
		curSum -= root.Val
		return
	}
	cur = append(cur, root.Val)
	curSum += root.Val
	if root.Left != nil {
		preOrder(root.Left, target)
	}
	if root.Right != nil {
		preOrder(root.Right, target)
	}
	cur = cur[:len(cur)-1]
	curSum -= root.Val
	return
}
