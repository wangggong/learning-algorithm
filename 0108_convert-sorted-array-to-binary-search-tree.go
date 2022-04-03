/*
 * @lc app=leetcode.cn id=convert-sorted-array-to-binary-search-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (76.39%)
 * Total Accepted:    237.3K
 * Total Submissions: 310.1K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
 *
 * 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-10,-3,0,5,9]
 * 输出：[0,-3,9,-10,null,5]
 * 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,3]
 * 输出：[3,1]
 * 解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 按 严格递增 顺序排列
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

/*
 * func sortedArrayToBST(nums []int) *TreeNode {
 * 	return preorder(nums, 0, len(nums))
 * }
 *
 * func preorder(arr []int, p, q int) *TreeNode {
 * 	if p >= q {
 * 		return nil
 * 	}
 * 	mid := p + (q-p)>>1
 * 	return &TreeNode{
 * 		Val:   arr[mid],
 * 		Left:  preorder(arr, p, mid),
 * 		Right: preorder(arr, mid+1, q),
 * 	}
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	var S [][2]int
	m := make(map[[2]int]*TreeNode)
	S = append(S, [2]int{0, n})
	for len(S) > 0 {
		k := S[len(S)-1]
		S = S[:len(S)-1]
		p, q := k[0], k[1]
		mid := p + (q-p)>>1
		left, right := [2]int{p, mid}, [2]int{mid + 1, q}
		node, ok := m[k]
		if !ok {
			m[k] = &TreeNode{Val: nums[mid]}
			S = append(S, k)
			if mid+1 < q {
				S = append(S, right)
			}
			if p < mid {
				S = append(S, left)
			}
		} else {
			if mid+1 < q {
				node.Right = m[right]
			}
			if p < mid {
				node.Left = m[left]
			}
		}
	}
	return m[[2]int{0, n}]
}
