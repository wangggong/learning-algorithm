/*
 * @lc app=leetcode.cn id=rectangle-area-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [850] 矩形面积 II
 *
 * https://leetcode.cn/problems/rectangle-area-ii/description/
 *
 * algorithms
 * Hard (48.54%)
 * Total Accepted:    17K
 * Total Submissions: 26.8K
 * Testcase Example:  '[[0,0,2,2],[1,0,2,3],[1,0,3,1]]'
 *
 * 我们给出了一个（轴对齐的）二维矩形列表 rectangles 。 对于 rectangle[i] = [x1, y1, x2,
 * y2]，其中（x1，y1）是矩形 i 左下角的坐标， (xi1, yi1) 是该矩形 左下角 的坐标， (xi2, yi2) 是该矩形 右上角
 * 的坐标。
 *
 * 计算平面中所有 rectangles 所覆盖的 总面积 。任何被两个或多个矩形覆盖的区域应只计算 一次 。
 *
 * 返回 总面积 。因为答案可能太大，返回 10^9 + 7 的 模 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：rectangles = [[0,0,2,2],[1,0,2,3],[1,0,3,1]]
 * 输出：6
 * 解释：如图所示，三个矩形覆盖了总面积为6的区域。
 * 从(1,1)到(2,2)，绿色矩形和红色矩形重叠。
 * 从(1,0)到(2,3)，三个矩形都重叠。
 *
 *
 * 示例 2：
 *
 *
 * 输入：rectangles = [[0,0,1000000000,1000000000]]
 * 输出：49
 * 解释：答案是 10^18 对 (10^9 + 7) 取模的结果， 即 49 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= rectangles.length <= 200
 * rectanges[i].length = 4
 * 0 <= xi1, yi1, xi2, yi2 <= 10^9
 * 矩形叠加覆盖后的总面积不会超越 2^63 - 1 ，这意味着可以用一个 64 位有符号整数来保存面积结果。
 *
 *
 */

import "sort"

const mod int = 1e9 + 7

func rectangleArea(rs [][]int) (ans int) {
	var xs []int
	for _, r := range rs {
		xs = append(xs, r[0], r[2])
	}
	sort.Ints(xs)
	for i, n := 1, len(xs); i < n; i++ {
		if xs[i] == xs[i-1] {
			continue
		}
		var as [][2]int
		for _, r := range rs {
			if r[0] <= xs[i-1] && xs[i] <= r[2] {
				as = append(as, [2]int{r[1], 1}, [2]int{r[3], -1})
			}
		}
		sort.Slice(as, func(p, q int) bool { return as[p][0] < as[q][0] })
		cur, last := 0, 0
		for j, a := range as {
			if a[1] == 1 {
				cur++
				continue
			}
			cur--
			if cur == 0 {
				ans = (ans + (xs[i]-xs[i-1])*(as[j][0]-as[last][0])) % mod
				last = j + 1
			}
		}
	}
	return
}
