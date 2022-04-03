/*
 * Medium
 *
 * 给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
 *
 * 节点 p 的后继是值比 p.val 大的节点中键值最小的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root = [2,1,3], p = 1
 * 输出：2
 * 解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
 * 示例 2：
 *
 *
 *
 * 输入：root = [5,3,6,2,4,null,null,1], p = 6
 * 输出：null
 * 解释：因为给出的节点没有中序后继，所以答案就返回 null 了。
 *
 *
 * 提示：
 *
 * 树中节点的数目在范围 [1, 104] 内。
 * -105 <= Node.val <= 105
 * 树中各节点的值均保证唯一。
 * 通过次数10,308
 * 提交次数16,167
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	for root != nil {
		if root.Val > p.Val {
			// 如果当前节点大于目标节点, 更新答案, 当前节点变为目标节点的左节点;
			ans = root
			root = root.Left
		} else {
			// 如果当前节点小于等于目标节点, 当前节点变为目标节点的右节点;
			// 注意到如果当前节点等于目标节点, 分几种情况:
			// 1. 目标节点有右子树, 去右子树找;
			// 2. 目标节点没右子树, 当前 `ans` 仍然有效;
			// 3. 目标节点最大, 还是去右子树找, 此时 `ans` 还没被赋值, 直接返回 nil.
			root = root.Right
		}
	}
	return ans
}

/*
 * var arr []*TreeNode
 *
 * func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
 * 	if root == nil || p == nil {
 * 		return nil
 * 	}
 * 	arr = nil
 * 	inOrder(root)
 * 	ind := bs(arr, p.Val)
 * 	if ind+1 >= len(arr) {
 * 		return nil
 * 	}
 * 	return arr[ind+1]
 * }
 *
 * func inOrder(root *TreeNode) {
 * 	if root == nil {
 * 		return
 * 	}
 * 	inOrder(root.Left)
 * 	arr = append(arr, root)
 * 	inOrder(root.Right)
 * }
 *
 * func bs(arr []*TreeNode, target int) int {
 * 	p, q := 0, len(arr)-1
 * 	for p <= q {
 * 		mid := (p + q) >> 1
 * 		if arr[mid].Val == target {
 * 			return mid
 * 		} else if arr[mid].Val > target {
 * 			q = mid - 1
 * 		} else {
 * 			p = mid + 1
 * 		}
 * 	}
 * 	return -1
 * }
 */
