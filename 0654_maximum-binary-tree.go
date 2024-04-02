/*
 * @lc app=leetcode.cn id=maximum-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [654] 最大二叉树
 *
 * https://leetcode-cn.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (81.05%)
 * Total Accepted:    86.1K
 * Total Submissions: 106.3K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
 *
 *
 * 创建一个根节点，其值为 nums 中的最大值。
 * 递归地在最大值 左边 的 子数组前缀上 构建左子树。
 * 递归地在最大值 右边 的 子数组后缀上 构建右子树。
 *
 *
 * 返回 nums 构建的 最大二叉树 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,2,1,6,0,5]
 * 输出：[6,3,5,null,2,0,null,null,1]
 * 解释：递归调用如下所示：
 * - [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
 * ⁠   - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
 * ⁠       - 空数组，无子节点。
 * ⁠       - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
 * ⁠           - 空数组，无子节点。
 * ⁠           - 只有一个元素，所以子节点是一个值为 1 的节点。
 * ⁠   - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
 * ⁠       - 只有一个元素，所以子节点是一个值为 0 的节点。
 * ⁠       - 空数组，无子节点。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[3,null,2,null,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * 0 <= nums[i] <= 1000
 * nums 中的所有整数 互不相同
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return preorder(nums, 0, len(nums))
}

func preorder(arr []int, l, r int) *TreeNode {
	if l >= r {
		return nil
	}
	maxInd := l
	for i := l; i < r; i++ {
		if arr[i] > arr[maxInd] {
			maxInd = i
		}
	}
	return &TreeNode{
		Val:   arr[maxInd],
		Left:  preorder(arr, l, maxInd),
		Right: preorder(arr, maxInd+1, r),
	}
}

/*
 * // 单调栈解
 * func constructMaximumBinaryTree(nums []int) *TreeNode {
 * 	var nodes []*TreeNode
 * 	for _, n := range nums {
 * 		nodes = append(nodes, &TreeNode{Val: n})
 * 	}
 * 	var S []int
 * 	for i, n := range nums {
 * 		for len(S) > 0 && nums[S[len(S)-1]] < n {
 * 			nodes[i].Left = nodes[S[len(S)-1]]
 * 			S = S[:len(S)-1]
 * 		}
 * 		if len(S) > 0 {
 * 			nodes[S[len(S)-1]].Right = nodes[i]
 * 		}
 * 		S = append(S, i)
 * 	}
 * 	return nodes[S[0]]
 * }
 */
