/*
 * @lc app=leetcode.cn id=number-of-rectangles-that-can-form-the-largest-square lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1725] 可以形成最大正方形的矩形数目
 *
 * https://leetcode-cn.com/problems/number-of-rectangles-that-can-form-the-largest-square/description/
 *
 * algorithms
 * Easy (78.92%)
 * Total Accepted:    22.1K
 * Total Submissions: 26.8K
 * Testcase Example:  '[[5,8],[3,9],[5,12],[16,5]]'
 *
 * 给你一个数组 rectangles ，其中 rectangles[i] = [li, wi] 表示第 i 个矩形的长度为 li 、宽度为 wi 。
 *
 * 如果存在 k 同时满足 k <= li 和 k <= wi ，就可以将第 i 个矩形切成边长为 k 的正方形。例如，矩形 [4,6] 可以切成边长最大为 4
 * 的正方形。
 *
 * 设 maxLen 为可以从矩形数组 rectangles 切分得到的 最大正方形 的边长。
 *
 * 请你统计有多少个矩形能够切出边长为 maxLen 的正方形，并返回矩形 数目 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：rectangles = [[5,8],[3,9],[5,12],[16,5]]
 * 输出：3
 * 解释：能从每个矩形中切出的最大正方形边长分别是 [5,3,5,5] 。
 * 最大正方形的边长为 5 ，可以由 3 个矩形切分得到。
 *
 *
 * 示例 2：
 *
 *
 * 输入：rectangles = [[2,3],[3,7],[4,3],[3,7]]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= rectangles.length <= 1000
 * rectangles[i].length == 2
 * 1 <= l_i, w_i <= 1e9
 * l_i != w_i
 *
 *
 */
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countGoodRectangles(M [][]int) int {
	L := 0
	ans := 0
	for _, m := range M {
		c := min(m[0], m[1])
		if L < c {
			L = c
			ans = 1
			continue
		}
		if L == c {
			ans++
		}
	}
	return ans
}
