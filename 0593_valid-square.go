/*
 * @lc app=leetcode.cn id=valid-square lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [593] 有效的正方形
 *
 * https://leetcode.cn/problems/valid-square/description/
 *
 * algorithms
 * Medium (44.10%)
 * Total Accepted:    27.8K
 * Total Submissions: 59.9K
 * Testcase Example:  '[0,0]\n[1,1]\n[1,0]\n[0,1]'
 *
 * 给定2D空间中四个点的坐标 p1, p2, p3 和 p4，如果这四个点构成一个正方形，则返回 true 。
 *
 * 点的坐标 pi 表示为 [xi, yi] 。输入 不是 按任何顺序给出的。
 *
 * 一个 有效的正方形 有四条等边和四个等角(90度角)。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
 * 输出: True
 *
 *
 * 示例 2:
 *
 *
 * 输入：p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
 * 输出：false
 *
 *
 * 示例 3:
 *
 *
 * 输入：p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
 * 输出：true
 *
 *
 *
 *
 * 提示:
 *
 *
 * p1.length == p2.length == p3.length == p4.length == 2
 * -10^4 <= xi, yi <= 10^4
 *
 *
 */
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	d12, d13, d14 := D(p1, p2), D(p1, p3), D(p1, p4)
	if d12 != d13 && d12 != d14 && d13 != d14 {
		return false
	}
	if d12 == d13 {
		p3, p4 = p4, p3
	}
	if d13 == d14 {
		p2, p3 = p3, p2
	}
	a, b, c, d := D(p1, p2), D(p2, p3), D(p3, p4), D(p4, p1)
	return a != 0 && a == b && b == c && c == d && dot([]int{p2[0] - p1[0], p2[1] - p1[1]}, []int{p4[0] - p1[0], p4[1] - p1[1]}) == 0
}

func D(p, q []int) int   { return (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1]) }
func dot(p, q []int) int { return p[0]*q[0] + p[1]*q[1] }
