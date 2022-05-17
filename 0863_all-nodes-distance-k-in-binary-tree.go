/*
 * @lc app=leetcode.cn id=all-nodes-distance-k-in-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [863] 二叉树中所有距离为 K 的结点
 *
 * https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/description/
 *
 * algorithms
 * Medium (60.96%)
 * Total Accepted:    42.5K
 * Total Submissions: 69.7K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n2'
 *
 * 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。
 *
 * 返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
 * 输出：[7,4,1]
 * 解释：所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1
 *
 *
 * 示例 2:
 *
 *
 * 输入: root = [1], target = 1, k = 3
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * 节点数在 [1, 500] 范围内
 * 0 <= Node.val <= 500
 * Node.val 中所有值 不同
 * 目标结点 target 是树上的结点。
 * 0 <= k <= 1000
 *
 *
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// DFS 太复杂, 直接 BFS+建图吧...
const maxn int = 500

var G [maxn + 10][]int
var visit [maxn + 10]bool

func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	for i := range G {
		G[i] = nil
	}
	for i := range visit {
		visit[i] = false
	}
	buildGraph(root)
	var Q [][2]int
	Q = append(Q, [2]int{target.Val, 0})
	visit[target.Val] = true
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		if q[1] == k {
			ans = append(ans, q[0])
			continue
		}
		for _, g := range G[q[0]] {
			if visit[g] {
				continue
			}
			Q = append(Q, [2]int{g, q[1] + 1})
			visit[g] = true
		}
	}
	return
}

func buildGraph(root *TreeNode) {
	if root == nil {
		return
	}
	val := root.Val
	if root.Left != nil {
		leftVal := root.Left.Val
		G[val] = append(G[val], leftVal)
		G[leftVal] = append(G[leftVal], val)
	}
	if root.Right != nil {
		rightVal := root.Right.Val
		G[val] = append(G[val], rightVal)
		G[rightVal] = append(G[rightVal], val)
	}
	buildGraph(root.Left)
	buildGraph(root.Right)
	return
}
