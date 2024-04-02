/*
 * @lc app=leetcode.cn id=find-right-interval lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [436] 寻找右区间
 *
 * https://leetcode-cn.com/problems/find-right-interval/description/
 *
 * algorithms
 * Medium (49.52%)
 * Total Accepted:    17.4K
 * Total Submissions: 32.8K
 * Testcase Example:  '[[1,2]]'
 *
 * 给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。
 *
 * 区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。
 *
 * 返回一个由每个区间 i 的 右侧区间 的最小起始位置组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,2]]
 * 输出：[-1]
 * 解释：集合中只有一个区间，所以输出-1。
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[3,4],[2,3],[1,2]]
 * 输出：[-1,0,1]
 * 解释：对于 [3,4] ，没有满足条件的“右侧”区间。
 * 对于 [2,3] ，区间[3,4]具有最小的“右”起点;
 * 对于 [1,2] ，区间[2,3]具有最小的“右”起点。
 *
 *
 * 示例 3：
 *
 *
 * 输入：intervals = [[1,4],[2,3],[3,4]]
 * 输出：[-1,2,-1]
 * 解释：对于区间 [1,4] 和 [3,4] ，没有满足条件的“右侧”区间。
 * 对于 [2,3] ，区间 [3,4] 有最小的“右”起点。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 2 * 10^4
 * intervals[i].length == 2
 * -10^6 <= starti <= endi <= 10^6
 * 每个间隔的起点都 不相同
 *
 *
 */
import "sort"

/*
 * // 二分, 容易想到.
 * func findRightInterval(intervals [][]int) []int {
 * 	var starts [][2]int
 * 	for i, interval := range intervals {
 * 		starts = append(starts, [2]int{interval[0], i})
 * 	}
 * 	n := len(starts)
 * 	sort.Slice(starts, func(p, q int) bool {
 * 		sp, sq := starts[p], starts[q]
 * 		return sp[0] < sq[0] || sp[0] == sq[0] && sp[1] < sq[1]
 * 	})
 * 	var ans []int
 * 	for _, interval := range intervals {
 * 		end := interval[1]
 * 		p, q := 0, n
 * 		for p < q {
 * 			mid := (p + q) >> 1
 * 			if starts[mid][0] >= end {
 * 				q = mid
 * 			} else {
 * 				p = mid + 1
 * 			}
 * 		}
 * 		if p == n {
 * 			ans = append(ans, -1)
 * 		} else {
 * 			ans = append(ans, starts[p][1])
 * 		}
 * 	}
 * 	return ans
 * }
 */

// 参考: https://leetcode.cn/problems/find-right-interval/solution/by-ac_oier-sijp/
//
// 双指针 (莫队思想), 通过调整询问的处理顺序, 可以减少目标位置的指针移动次数.
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	var starts, ends [][2]int
	for i, interval := range intervals {
		start, end := interval[0], interval[1]
		starts = append(starts, [2]int{start, i})
		ends = append(ends, [2]int{end, i})
	}
	sort.Slice(starts, func(p, q int) bool { return starts[p][0] < starts[q][0] })
	sort.Slice(ends, func(p, q int) bool { return ends[p][0] < ends[q][0] })
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for p, q := 0, 0; p < n; p++ {
		for ; q < n && starts[q][0] < ends[p][0]; q++ {
		}
		if q == n {
			break
		}
		ans[ends[p][1]] = starts[q][1]
	}
	return ans
}