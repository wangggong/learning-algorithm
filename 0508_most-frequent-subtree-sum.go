/*
 * @lc app=leetcode.cn id=most-frequent-subtree-sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [508] 出现次数最多的子树元素和
 *
 * https://leetcode-cn.com/problems/most-frequent-subtree-sum/description/
 *
 * algorithms
 * Medium (68.21%)
 * Total Accepted:    16.7K
 * Total Submissions: 24.5K
 * Testcase Example:  '[5,2,-3]'
 *
 * 给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
 *
 * 一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入: root = [5,2,-3]
 * 输出: [2,-3,4]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入: root = [5,2,-5]
 * 输出: [2]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 节点数在 [1, 10^4] 范围内
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

// 一开始以为问题很恶心, 结果没看到 **所有** 节点总和.
var counter map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	counter = make(map[int]int)
	dfs(root)
	maxVal := 0
	for v := range counter {
		if counter[v] > maxVal {
			maxVal = counter[v]
		}
	}
	var ans []int
	for v := range counter {
		if counter[v] == maxVal {
			ans = append(ans, v)
		}
	}
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := dfs(root.Left), dfs(root.Right)
	v := left + right + root.Val
	counter[v]++
	return v
}
