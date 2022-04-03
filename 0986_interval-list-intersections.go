/*
 * @lc app=leetcode.cn id=interval-list-intersections lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [986] 区间列表的交集
 *
 * https://leetcode-cn.com/problems/interval-list-intersections/description/
 *
 * algorithms
 * Medium (68.40%)
 * Total Accepted:    38.4K
 * Total Submissions: 56.2K
 * Testcase Example:  '[[0,2],[5,10],[13,23],[24,25]]\n[[1,5],[8,12],[15,24],[25,26]]'
 *
 * 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而
 * secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
 *
 * 返回这 两个区间列表的交集 。
 *
 * 形式上，闭区间 [a, b]（其中 a ）表示实数 x 的集合，而 a  。
 *
 * 两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList =
 * [[1,5],[8,12],[15,24],[25,26]]
 * 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：firstList = [[1,3],[5,9]], secondList = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：firstList = [], secondList = [[4,8],[10,12]]
 * 输出：[]
 *
 *
 * 示例 4：
 *
 *
 * 输入：firstList = [[1,7]], secondList = [[3,10]]
 * 输出：[[3,7]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * firstList.length + secondList.length >= 1
 * 0 i < endi
 * endi < starti+1
 * 0 j < endj
 * endj < startj+1
 *
 *
 */
func intervalIntersection(fl [][]int, sl [][]int) [][]int {
	n, m := len(fl), len(sl)
	p, q := 0, 0
	var ans [][]int
	// NOTE: 区间合并, 取区间起点的较大值和终点的较小值. 然后取两个区间中终点较小的前进.
	for p < n && q < m {
		low := max(fl[p][0], sl[q][0])
		high := min(fl[p][1], sl[q][1])
		if low <= high {
			ans = append(ans, []int{low, high})
		}
		if fl[p][1] < sl[q][1] {
			p++
		} else {
			q++
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
