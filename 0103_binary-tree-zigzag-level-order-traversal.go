/*
 * @lc app=leetcode.cn id=binary-tree-zigzag-level-order-traversal lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [103] 二叉树的锯齿形层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (57.22%)
 * Total Accepted:    211.2K
 * Total Submissions: 369.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[20,9],[15,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 2000] 内
 * -100 <= Node.val <= 100
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

type NodeWithLvl struct {
	Node *TreeNode
	Lvl  int
}

var ans [][]int

// func zigzagLevelOrder(root *TreeNode) [][]int {
// 	ans = nil
// 	if root == nil {
// 		return ans
// 	}
// 	var Q []*NodeWithLvl
// 	Q = append(Q, &NodeWithLvl{root, 0})
// 	for len(Q) > 0 {
// 		q := Q[0]
// 		Q = Q[1:]
// 		node, lvl := q.Node, q.Lvl
// 		if lvl >= len(ans) {
// 			ans = append(ans, nil)
// 		}
// 		if lvl % 2 == 0 {
// 			ans[lvl] = append(ans[lvl], node.Val)
// 		} else {
// 			ans[lvl] = append([]int{node.Val}, ans[lvl]...)
// 		}
// 		if node.Left != nil {
// 			Q = append(Q, &NodeWithLvl{node.Left, lvl + 1})
// 		}
// 		if node.Right != nil {
// 			Q = append(Q, &NodeWithLvl{node.Right, lvl + 1})
// 		}
// 	}
// 	// n := len(ans)
// 	// for i := 1; i < n; i += 2 {
// 	// 	reverse(i)
// 	// }
// 	return ans
// }

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans = nil
	if root == nil {
		return ans
	}
	var Q []*TreeNode
	Q = append(Q, root)
	lvl := 0
	for len(Q) > 0 {
		size := len(Q)
		for i := 0; i < size; i++ {
			q := Q[0]
			Q = Q[1:]
			if lvl >= len(ans) {
				ans = append(ans, nil)
			}
			if lvl%2 == 0 {
				ans[lvl] = append(ans[lvl], q.Val)
			} else {
				ans[lvl] = append([]int{q.Val}, ans[lvl]...)
			}
			if q.Left != nil {
				Q = append(Q, q.Left)
			}
			if q.Right != nil {
				Q = append(Q, q.Right)
			}
		}
		lvl++
	}
	return ans
}

func reverse(k int) {
	for i, j := 0, len(ans[k])-1; i < j; i, j = i+1, j-1 {
		ans[k][i], ans[k][j] = ans[k][j], ans[k][i]
	}
}
