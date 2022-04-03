/*
 * @lc app=leetcode.cn id=minimum-number-of-arrows-to-burst-balloons lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [452] 用最少数量的箭引爆气球
 *
 * https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
 *
 * algorithms
 * Medium (50.72%)
 * Total Accepted:    123.7K
 * Total Submissions: 244K
 * Testcase Example:  '[[10,16],[2,8],[1,6],[7,12]]'
 *
 *
 * 在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。
 *
 * 一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足
 * xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。
 * 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
 *
 * 给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。
 *
 *
 * 示例 1：
 *
 *
 * 输入：points = [[10,16],[2,8],[1,6],[7,12]]
 * 输出：2
 * 解释：对于该样例，x = 6 可以射爆 [2,8],[1,6] 两个气球，以及 x = 11 射爆另外两个气球
 *
 * 示例 2：
 *
 *
 * 输入：points = [[1,2],[3,4],[5,6],[7,8]]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：points = [[1,2],[2,3],[3,4],[4,5]]
 * 输出：2
 *
 *
 * 示例 4：
 *
 *
 * 输入：points = [[1,2]]
 * 输出：1
 *
 *
 * 示例 5：
 *
 *
 * 输入：points = [[2,3],[2,3]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= points.length <= 1e4
 * points[i].length == 2
 * -1 << 31 <= xstart < xend <= 1 << 31 - 1
 *
 *
 */
import "sort"

/*
 * func findMinArrowShots(P [][]int) int {
 * 	sort.Slice(P, func(p, q int) bool {
 * 		if P[p][0] != P[q][0] {
 * 			return P[p][0] < P[q][0]
 * 		}
 * 		return P[p][1] < P[q][1]
 * 	})
 * 	n := len(P)
 * 	ans := 0
 * 	var k int
 * 	for i := 0; i < n; i++ {
 * 		if k > 0 && k >= P[i][0] && k <= P[i][1] {
 * 			continue
 * 		}
 * 		// 确定上下界逐渐逼近. 然后无脑取上界. 奇怪的贪法.
 * 		lower, upper := P[i][0], P[i][1]
 * 		for ; i < n && lower < upper; i++ {
 * 			if lower > P[i][1] || upper < P[i][0] {
 * 				break
 * 			}
 * 			if P[i][0] > lower {
 * 				lower = P[i][0]
 * 			}
 * 			if P[i][1] < upper {
 * 				upper = P[i][1]
 * 			}
 * 		}
 * 		k = upper
 * 		i--
 * 		ans++
 * 	}
 * 	return ans
 * }
 */


func findMinArrowShots(P [][]int) int {
	n := len(P)
	if n == 0 {
		return 0
	}
	sort.Slice(P, func(p, q int) bool { return P[p][1] < P[q][1] })
	ans := 1
	k := P[0][1]
	for i := 0; i < n; i++ {
		if P[i][0] > k {
			k = P[i][1]
			ans++
		}
	}
	return ans
}
