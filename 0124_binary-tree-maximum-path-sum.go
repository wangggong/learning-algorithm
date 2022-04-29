/**
 * @lc app=leetcode.cn id=binary-tree-maximum-path-sum lang=golang
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/description/
 *
 * * algorithms
 * * Hard (44.87%)
 * * Total Accepted:    214.8K
 * * Total Submissions: 478.7K
 * * Testcase Example:  '[1,2,3]'
 *
 * 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
 *
 * 路径和 是路径中各节点值的总和。
 *
 * 给你一个二叉树的根节点 root ，返回其 最大路径和 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3]
 * 输出：6
 * 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
 *
 * 示例 2：
 *
 *
 * 输入：root = [-10,9,20,null,null,15,7]
 * 输出：42
 * 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
 *
 *
 *
 *
 * 提示：
 *
 *
 * 	树中节点数目范围是 [1, 3 * 10^4]
 * 	-1000
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

// 开 memo 被 challenge 了: 这里采用后序遍历的话 memo 是用不到的. 所以就删了.
var ans int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = root.Val
	postOrder(root)
	return ans
}

func postOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := postOrder(root.Left), postOrder(root.Right)
	cur := max(max(left, right), 0) + root.Val
	ans = max(ans, max(cur, left+right+root.Val))
	return cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
 * // 另: 考虑以当前根节点为起点的最长路径和. 注意此时空节点需要返回系统最小, 避免一个节点不取.
 * var pathSum map[*TreeNode]int
 * var ans int
 *
 * func maxPathSum(root *TreeNode) int {
 * 	ans = math.MinInt32
 * 	pathSum = make(map[*TreeNode]int)
 * 	postOrder(root)
 * 	return ans
 * }
 *
 * func postOrder(root *TreeNode) int {
 * 	if root == nil {
 * 		return math.MinInt32
 * 	}
 * 	if v, ok := pathSum[root]; ok {
 * 		return v
 * 	}
 * 	v := 0
 * 	defer func() { pathSum[root] = v }()
 * 	left, right := postOrder(root.Left), postOrder(root.Right)
 * 	v = max(max(left, right), 0) + root.Val
 * 	ans = max(ans, left+right+root.Val)
 * 	ans = max(ans, v)
 * 	return v
 * }
 *
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 */
