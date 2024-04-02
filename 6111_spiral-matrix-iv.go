/*
 * @lc app=leetcode.cn id=spiral-matrix-iv lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6111] 螺旋矩阵 IV
 *
 * https://leetcode.cn/problems/spiral-matrix-iv/description/
 *
 * algorithms
 * Medium (66.17%)
 * Total Accepted:    5.5K
 * Total Submissions: 8.2K
 * Testcase Example:  '3\n5\n[3,0,2,6,8,1,7,9,4,2,5,5,0]'
 *
 * 给你两个整数：m 和 n ，表示矩阵的维数。
 *
 * 另给你一个整数链表的头节点 head 。
 *
 * 请你生成一个大小为 m x n 的螺旋矩阵，矩阵包含链表中的所有整数。链表中的整数从矩阵 左上角 开始、顺时针 按 螺旋
 * 顺序填充。如果还存在剩余的空格，则用 -1 填充。
 *
 * 返回生成的矩阵。
 *
 *
 *
 * 示例 1：
 *
 * 输入：m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
 * 输出：[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
 * 解释：上图展示了链表中的整数在矩阵中是如何排布的。
 * 注意，矩阵中剩下的空格用 -1 填充。
 *
 *
 * 示例 2：
 *
 * 输入：m = 1, n = 4, head = [0,1,2]
 * 输出：[[0,1,2,-1]]
 * 解释：上图展示了链表中的整数在矩阵中是如何从左到右排布的。
 * 注意，矩阵中剩下的空格用 -1 填充。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 10^5
 * 1 <= m * n <= 10^5
 * 链表中节点数目在范围 [1, m * n] 内
 * 0 <= Node.val <= 1000
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(n int, m int, head *ListNode) [][]int {
	G := make([][]int, n)
	for i := range G {
		G[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			G[i][j] = -1
		}
	}
	curr := head
	for i := 0; i*2 < min(n, m)+1; i++ {
		for j := 0; i+j < m-i && curr != nil; j++ {
			G[i][i+j] = curr.Val
			curr = curr.Next
		}
		if curr == nil {
			break
		}
		for j := 1; i+j < n-i && curr != nil; j++ {
			G[i+j][m-1-i] = curr.Val
			curr = curr.Next
		}
		if curr == nil {
			break
		}
		for j := 1; i < n-1-i && m-1-i-j >= i && curr != nil; j++ {
			G[n-1-i][m-1-i-j] = curr.Val
			curr = curr.Next
		}
		if curr == nil {
			break
		}
		for j := 1; i < m-1-i && n-1-i-j > i && curr != nil; j++ {
			G[n-1-i-j][i] = curr.Val
			curr = curr.Next
		}
		if curr == nil {
			break
		}
	}
	return G
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
