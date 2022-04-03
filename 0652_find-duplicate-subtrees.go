import "fmt"

/*
 * @lc app=leetcode.cn id=find-duplicate-subtrees lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [652] 寻找重复的子树
 *
 * https://leetcode-cn.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (57.37%)
 * Total Accepted:    44.3K
 * Total Submissions: 77K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * 给定一棵二叉树 root，返回所有重复的子树。
 *
 * 对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 *
 * 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,2,3,4,null,2,4,null,null,4]
 * 输出：[[2,4],[4]]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [2,1,1]
 * 输出：[[1]]
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入：root = [2,2,2,3,null,3,null]
 * 输出：[[2,3],[3]]
 *
 *
 *
 * 提示：
 *
 *
 * 树中的结点数在[1,10^4]范围内。
 * -200 <= Node.val <= 200
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

var ans []*TreeNode
var m map[int][]*TreeNode
var count map[string]int

/*
 * func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
 * 	ans = nil
 * 	if root == nil {
 * 		return ans
 * 	}
 * 	m = make(map[int][]*TreeNode)
 * 	traversal(root)
 * 	return ans
 * }
 */

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans = nil
	count = make(map[string]int)
	collect(root)
	return ans
}

func collect(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	s := fmt.Sprintf("[%s]%d[%s]", collect(root.Left), root.Val, collect(root.Right))
	count[s]++
	if count[s] == 2 {
		ans = append(ans, root)
	}
	return s
}

func traversal(root *TreeNode) {
	// assert root != nil;
	for _, node := range m[root.Val] {
		if !sameTree(root, node) {
			continue
		}
		contains := false
		for _, a := range ans {
			if a == node {
				contains = true
				break
			}
		}
		if !contains {
			ans = append(ans, node)
		}
		break
	}
	m[root.Val] = append(m[root.Val], root)
	if root.Left != nil {
		traversal(root.Left)
	}
	if root.Right != nil {
		traversal(root.Right)
	}
}

func sameTree(root1, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == nil && root2 == nil
	}
	if root1.Val != root2.Val {
		return false
	}
	return sameTree(root1.Left, root2.Left) && sameTree(root1.Right, root2.Right)
}
