/*
 * @lc app=leetcode.cn id=pseudo-palindromic-paths-in-a-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1457] 二叉树中的伪回文路径
 *
 * https://leetcode-cn.com/problems/pseudo-palindromic-paths-in-a-binary-tree/description/
 *
 * algorithms
 * Medium (62.72%)
 * Total Accepted:    9K
 * Total Submissions: 14.4K
 * Testcase Example:  '[2,3,1,3,1,null,1]'
 *
 * 给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
 *
 * 请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [2,3,1,3,1,null,1]
 * 输出：2
 * 解释：上图为给定的二叉树。总共有 3 条从根到叶子的路径：红色路径 [2,3,3] ，绿色路径 [2,1,1] 和路径 [2,3,1] 。
 * ⁠    在这些路径中，只有红色和绿色的路径是伪回文路径，因为红色路径 [2,3,3] 存在回文排列 [3,2,3] ，绿色路径 [2,1,1]
 * 存在回文排列 [1,2,1] 。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [2,1,1,1,3,null,null,null,null,null,1]
 * 输出：1
 * 解释：上图为给定二叉树。总共有 3 条从根到叶子的路径：绿色路径 [2,1,1] ，路径 [2,1,3,1] 和路径 [2,1] 。
 * ⁠    这些路径中只有绿色路径是伪回文路径，因为 [2,1,1] 存在回文排列 [1,2,1] 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [9]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定二叉树的节点数目在范围 [1, 10^5] 内
 * 1 <= Node.val <= 9
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

var count map[int]int
var ans int

func pseudoPalindromicPaths(root *TreeNode) int {
	count = make(map[int]int)
	ans = 0
	backtrack(root)
	return ans
}

func backtrack(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		count[root.Val]++
		if checkPalindromic() {
			ans++
		}
		count[root.Val]--
		return
	}
	count[root.Val]++
	backtrack(root.Left)
	backtrack(root.Right)
	count[root.Val]--
}

func checkPalindromic() bool {
	// fmt.Println(count)
	single := false
	for val := range count {
		if count[val]%2 != 0 {
			if !single {
				single = true
			} else {
				return false
			}
		}
	}
	return true
}
