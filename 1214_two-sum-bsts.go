/*
 * Medium
 *
 * 给出两棵二叉搜索树的根节点 root1 和 root2 ，请你从两棵树中各找出一个节点，使得这两个节点的值之和等于目标值 Target。
 *
 * 如果可以找到返回 True，否则返回 False。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root1 = [2,1,4], root2 = [1,0,3], target = 5
 * 输出：true
 * 解释：2 加 3 和为 5 。
 * 示例 2：
 *
 *
 *
 * 输入：root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
 * 输出：false
 *
 *
 * 提示：
 *
 * 每棵树上节点数在 [1, 5000] 范围内。
 * -109 <= Node.val, target <= 109
 * 通过次数3,962
 * 提交次数6,127
 */

/*
 * func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
 * 	if root1 == nil || root2 == nil {
 * 		return false
 * 	}
 * 	switch v := root1.Val + root2.Val - target; {
 * 	case v == 0:
 * 		return true
 * 	case v > 0:
 * 		return twoSumBSTs(root1.Left, root2, target) || twoSumBSTs(root1, root2.Left, target)
 * 	case v < 0:
 * 		return twoSumBSTs(root1.Right, root2, target) || twoSumBSTs(root1, root2.Right, target)
 * 	default:
 * 		// unreachable
 * 	}
 * 	return false
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	arr1 := inOrder(root1)
	arr2 := inOrder(root2)
	m, n := len(arr1), len(arr2)
	p, q := 0, n-1
	for p < m && q >= 0 {
		if v := arr1[p] + arr2[q]; v == target {
			return true
		} else if v > target {
			q--
		} else {
			p++
		}
	}
	return false
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	left := inOrder(root.Left)
	right := inOrder(root.Right)
	ans = append(ans, left...)
	ans = append(ans, root.Val)
	ans = append(ans, right...)
	return ans
}

