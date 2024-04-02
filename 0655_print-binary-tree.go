/*
 * @lc app=leetcode.cn id=print-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [655] 输出二叉树
 *
 * https://leetcode.cn/problems/print-binary-tree/description/
 *
 * algorithms
 * Medium (60.83%)
 * Total Accepted:    26.9K
 * Total Submissions: 39.2K
 * Testcase Example:  '[1,2]'
 *
 * 给你一棵二叉树的根节点 root ，请你构造一个下标从 0 开始、大小为 m x n 的字符串矩阵 res ，用以表示树的 格式化布局
 * 。构造此格式化布局矩阵需要遵循以下规则：
 *
 *
 * 树的 高度 为 height ，矩阵的行数 m 应该等于 height + 1 。
 * 矩阵的列数 n 应该等于 2^height+1 - 1 。
 * 根节点 需要放置在 顶行 的 正中间 ，对应位置为 res[0][(n-1)/2] 。
 * 对于放置在矩阵中的每个节点，设对应位置为 res[r][c] ，将其左子节点放置在 res[r+1][c-2^height-r-1] ，右子节点放置在
 * res[r+1][c+2^height-r-1] 。
 * 继续这一过程，直到树中的所有节点都妥善放置。
 * 任意空单元格都应该包含空字符串 "" 。
 *
 *
 * 返回构造得到的矩阵 res 。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2]
 * 输出：
 * [["","1",""],
 * ["2","",""]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3,null,4]
 * 输出：
 * [["","","","1","","",""],
 * ["","2","","","","3",""],
 * ["","","4","","","",""]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数在范围 [1, 2^10] 内
 * -99 <= Node.val <= 99
 * 树的深度在范围 [1, 10] 内
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
func printTree(root *TreeNode) (ans [][]string) {
	depth := getDepth(root)
	cols := 1<<depth - 1
	ans = make([][]string, depth)
	for i := range ans {
		ans[i] = make([]string, cols)
	}
	traversal(root, ans, 0, cols>>1)
	return
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}

func traversal(root *TreeNode, M [][]string, r, c int) {
	if root == nil {
		return
	}
	M[r][c] = itoa(root.Val)
	diff := (1 << (len(M) - r - 1)) >> 1
	traversal(root.Left, M, r+1, c-diff)
	traversal(root.Right, M, r+1, c+diff)
	return
}

func itoa(v int) string {
	neg := false
	if v < 0 {
		neg, v = true, -v
	}
	var vals []byte
	for ; v != 0; v /= 10 {
		vals = append(vals, byte(v%10)+'0')
	}
	if len(vals) == 0 {
		vals = append(vals, '0')
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
	if neg {
		return "-" + string(vals)
	}
	return string(vals)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
