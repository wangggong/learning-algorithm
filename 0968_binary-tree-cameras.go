/*
 * @lc app=leetcode.cn id=binary-tree-cameras lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [968] 监控二叉树
 *
 * https://leetcode-cn.com/problems/binary-tree-cameras/description/
 *
 * algorithms
 * Hard (50.90%)
 * Total Accepted:    32.3K
 * Total Submissions: 63.4K
 * Testcase Example:  '[0,0,null,0,0]'
 *
 * 给定一个二叉树，我们在树的节点上安装摄像头。
 *
 * 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
 *
 * 计算监控树的所有节点所需的最小摄像头数量。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[0,0,null,0,0]
 * 输出：1
 * 解释：如图所示，一台摄像头足以监控所有节点。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：[0,0,null,0,null,0,null,null,0]
 * 输出：2
 * 解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。
 *
 *
 *
 * 提示：
 *
 *
 * 给定树的节点数的范围是 [1, 1000]。
 * 每个节点的值都是 0。
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

// 三种状态: 无覆盖 / 覆盖 / 有摄像头.
const (
	noCover = iota
	cover
	camera
)

var result int

func minCameraCover(root *TreeNode) int {
	result = 0
	if traversal(root) == noCover {
		result++
	}
	return result
}

func traversal(root *TreeNode) int {
	// 空节点被认为是有覆盖的, 避免叶子结点再安摄像头.
	if root == nil {
		return cover
	}
	left, right := traversal(root.Left), traversal(root.Right)
	// case 1: 两边有一边没覆盖, 得安摄像头.
	if left == noCover || right == noCover {
		result++
		return camera
	}
	// case 2: 两边都覆盖了, 本节点 ---- 就覆盖不到了.
	// 别覆盖, 等父节点覆盖 or 根节点按摄像头.
	if left == cover && right == cover {
		return noCover
	}
	// case 3: 有一边安了摄像头, 就被覆盖了.
	if left == camera || right == camera {
		return cover
	}
	// unreachable
	return noCover
}

