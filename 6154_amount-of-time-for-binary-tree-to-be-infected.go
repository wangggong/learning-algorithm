/*
 * @lc app=leetcode.cn id=amount-of-time-for-binary-tree-to-be-infected lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6154] 感染二叉树需要的总时间
 *
 * https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/description/
 *
 * algorithms
 * Medium (38.92%)
 * Total Accepted:    3.2K
 * Total Submissions: 8.1K
 * Testcase Example:  '[1,5,3,null,4,10,6,9,2]\n3'
 *
 * 给你一棵二叉树的根节点 root ，二叉树中节点的值 互不相同 。另给你一个整数 start 。在第 0 分钟，感染 将会从值为 start
 * 的节点开始爆发。
 *
 * 每分钟，如果节点满足以下全部条件，就会被感染：
 *
 *
 * 节点此前还没有感染。
 * 节点与一个已感染节点相邻。
 *
 *
 * 返回感染整棵树需要的分钟数。
 *
 *
 *
 * 示例 1：
 *
 * 输入：root = [1,5,3,null,4,10,6,9,2], start = 3
 * 输出：4
 * 解释：节点按以下过程被感染：
 * - 第 0 分钟：节点 3
 * - 第 1 分钟：节点 1、10、6
 * - 第 2 分钟：节点5
 * - 第 3 分钟：节点 4
 * - 第 4 分钟：节点 9 和 2
 * 感染整棵树需要 4 分钟，所以返回 4 。
 *
 *
 * 示例 2：
 *
 * 输入：root = [1], start = 1
 * 输出：0
 * 解释：第 0 分钟，树中唯一一个节点处于感染状态，返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 10^5] 内
 * 1 <= Node.val <= 10^5
 * 每个节点的值 互不相同
 * 树中必定存在值为 start 的节点
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
func amountOfTime(root *TreeNode, start int) (ans int) {
	edges := make(map[int]map[int]struct{})
	traversal(root, edges)
	visit := make(map[int]struct{})
	var Q []int
	Q = append(Q, start)
	ans = -1
	for len(Q) > 0 && len(visit) < len(edges) {
		ans++
		for size := len(Q); size > 0; size-- {
			q := Q[0]
			Q = Q[1:]
			if _, ok := visit[q]; ok {
				continue
			}
			visit[q] = struct{}{}
			for k := range edges[q] {
				Q = append(Q, k)
			}
		}
	}
	return
}

func traversal(root *TreeNode, edges map[int]map[int]struct{}) {
	if root == nil {
		return
	}
	traversal(root.Left, edges)
	traversal(root.Right, edges)
	if edges[root.Val] == nil {
		edges[root.Val] = make(map[int]struct{})
	}
	if root.Left != nil {
		edges[root.Val][root.Left.Val] = struct{}{}
		edges[root.Left.Val][root.Val] = struct{}{}
	}
	if root.Right != nil {
		edges[root.Val][root.Right.Val] = struct{}{}
		edges[root.Right.Val][root.Val] = struct{}{}
	}
	return
}
