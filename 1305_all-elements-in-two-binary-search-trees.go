/*
 * @lc app=leetcode.cn id=all-elements-in-two-binary-search-trees lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1305] 两棵二叉搜索树中的所有元素
 *
 * https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/description/
 *
 * algorithms
 * Medium (74.85%)
 * Total Accepted:    21.4K
 * Total Submissions: 28K
 * Testcase Example:  '[2,1,4]\r\n[1,0,3]\r'
 *
 * 给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root1 = [2,1,4], root2 = [1,0,3]
 * 输出：[0,1,1,2,3,4]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root1 = [1,null,8], root2 = [8,1]
 * 输出：[1,1,8,8]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每棵树的节点数在 [0, 5000] 范围内
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

var arr1, arr2, arr []int

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr = nil
	preOrder(root1)
	arr1 = copyOf(arr)
	arr = nil
	preOrder(root2)
	arr2 = copyOf(arr)
	return merge(arr1, arr2)
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	preOrder(root.Left)
	arr = append(arr, root.Val)
	preOrder(root.Right)
}

func copyOf(arr []int) []int {
	ans := make([]int, len(arr))
	copy(ans, arr)
	return ans
}

func merge(arr1, arr2 []int) []int {
	n, m := len(arr1), len(arr2)
	ans := make([]int, n+m)
	p, q, k := 0, 0, 0
	for p < n && q < m {
		if arr1[p] < arr2[q] {
			ans[k] = arr1[p]
			p++
			k++
		} else {
			ans[k] = arr2[q]
			q++
			k++
		}
	}
	for p < n {
		ans[k] = arr1[p]
		p++
		k++
	}
	for q < m {
		ans[k] = arr2[q]
		q++
		k++
	}
	return ans
}
