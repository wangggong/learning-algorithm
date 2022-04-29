/*
 * @lc app=leetcode.cn id=maximum-width-of-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [662] 二叉树最大宽度
 *
 * https://leetcode-cn.com/problems/maximum-width-of-binary-tree/description/
 *
 * algorithms
 * Medium (40.87%)
 * Total Accepted:    42.3K
 * Total Submissions: 103.2K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary
 * tree）结构相同，但一些节点为空。
 *
 * 每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
 *
 * 示例 1:
 *
 *
 * 输入:
 *
 * ⁠          1
 * ⁠        /   \
 * ⁠       3     2
 * ⁠      / \     \
 * ⁠     5   3     9
 *
 * 输出: 4
 * 解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        /
 * ⁠       3
 * ⁠      / \
 * ⁠     5   3
 *
 * 输出: 2
 * 解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
 *
 *
 * 示例 3:
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      /
 * ⁠     5
 *
 * 输出: 2
 * 解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
 *
 *
 * 示例 4:
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      /     \
 * ⁠     5       9
 * ⁠    /         \
 * ⁠   6           7
 * 输出: 8
 * 解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。
 *
 *
 * 注意: 答案在32位有符号整数的表示范围内。
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

type nodeInfo struct {
	node  *TreeNode
	depth int
	pos   int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	curDepth, left := 0, 0
	var Q []nodeInfo
	Q = append(Q, nodeInfo{root, 0, 0})
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		node, depth, pos := q.node, q.depth, q.pos
		if node.Left != nil {
			Q = append(Q, nodeInfo{node.Left, depth + 1, pos * 2})
		}
		if node.Right != nil {
			Q = append(Q, nodeInfo{node.Right, depth + 1, pos*2 + 1})
		}
		if depth != curDepth {
			curDepth = depth
			left = q.pos
		}
		ans = max(ans, pos-left+1)
	}
	return int(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
