/*
 * @lc app=leetcode.cn id=minimum-lines-to-represent-a-line-chart lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2280] 表示一个折线图的最少线段数
 *
 * https://leetcode-cn.com/problems/minimum-lines-to-represent-a-line-chart/description/
 *
 * algorithms
 * Medium (18.79%)
 * Total Accepted:    6K
 * Total Submissions: 31.9K
 * Testcase Example:  '[[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]]'
 *
 * 给你一个二维整数数组 stockPrices ，其中 stockPrices[i] = [dayi, pricei] 表示股票在 dayi 的价格为
 * pricei 。折线图 是一个二维平面上的若干个点组成的图，横坐标表示日期，纵坐标表示价格，折线图由相邻的点连接而成。比方说下图是一个例子：
 *
 * 请你返回要表示一个折线图所需要的 最少线段数 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：stockPrices = [[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]]
 * 输出：3
 * 解释：
 * 上图为输入对应的图，横坐标表示日期，纵坐标表示价格。
 * 以下 3 个线段可以表示折线图：
 * - 线段 1 （红色）从 (1,7) 到 (4,4) ，经过 (1,7) ，(2,6) ，(3,5) 和 (4,4) 。
 * - 线段 2 （蓝色）从 (4,4) 到 (5,4) 。
 * - 线段 3 （绿色）从 (5,4) 到 (8,1) ，经过 (5,4) ，(6,3) ，(7,2) 和 (8,1) 。
 * 可以证明，无法用少于 3 条线段表示这个折线图。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：stockPrices = [[3,4],[1,2],[7,8],[2,3]]
 * 输出：1
 * 解释：
 * 如上图所示，折线图可以用一条线段表示。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= stockPrices.length <= 10^5
 * stockPrices[i].length == 2
 * 1 <= dayi, pricei <= 10^9
 * 所有 dayi 互不相同 。
 *
 *
 */
import "sort"

func minimumLines(P [][]int) int {
	n := len(P)
	if n <= 1 {
		return 0
	}
    // 傻逼! 你忘了排序了!
	sort.Slice(P, func(p, q int) bool { return P[p][0] < P[q][0] })
	ans := 1
	for i := 1; i+1 < n; i++ {
		if sameSlope(P[i-1], P[i], P[i+1]) {
			continue
		}
		ans++
	}
	return ans
}

func sameSlope(p, q, r []int) bool {
	return (q[1]-p[1])*(r[0]-q[0]) == (q[0]-p[0])*(r[1]-q[1])
}
