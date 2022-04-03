import "math"

/*
 * @lc app=leetcode.cn id=minimum-absolute-difference-in-bst lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (62.29%)
 * Total Accepted:    100.5K
 * Total Submissions: 161K
 * Testcase Example:  '[4,2,6,1,3]'
 *
 * 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
 *
 * 差值是一个正数，其数值等于两值之差的绝对值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [4,2,6,1,3]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,0,48,null,null,12,49]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目范围是 [2, 10^4]
 * 0 <= Node.val <= 10^5
 *
 *
 *
 *
 * 注意：本题与 783
 * https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同
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
var prev *TreeNode

func getMinimumDifference(root *TreeNode) int {
	ans = math.MaxInt32
	prev = nil
	if root == nil {
		return ans
	}
	// inorder(root)
	inorder(root)
	return ans
}

/*
 * func inorder(root *TreeNode) {
 * 	// assert root != nil;
 * 	if root.Left != nil {
 * 		inorder(root.Left)
 * 		prev, curr := root, root.Left
 * 		for curr != nil {
 * 			prev, curr = curr, curr.Right
 * 		}
 * 		if ans > root.Val-prev.Val {
 * 			ans = root.Val - prev.Val
 * 		}
 * 	}
 * 	if root.Right != nil {
 * 		inorder(root.Right)
 * 		prev, curr := root, root.Right
 * 		for curr != nil {
 * 			prev, curr = curr, curr.Left
 * 		}
 * 		if ans > prev.Val-root.Val {
 * 			ans = prev.Val - root.Val
 * 		}
 * 	}
 * }
 */

func inorder(curr *TreeNode) {
	if curr.Left != nil {
		inorder(curr.Left)
	}
	if prev != nil {
		if ans > curr.Val - prev.Val {
			ans = curr.Val - prev.Val
		}
	}
	prev = curr
	if curr.Right != nil {
		inorder(curr.Right)
	}
}
