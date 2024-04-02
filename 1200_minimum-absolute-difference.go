/*
 * @lc app=leetcode.cn id=minimum-absolute-difference lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1200] 最小绝对差
 *
 * https://leetcode.cn/problems/minimum-absolute-difference/description/
 *
 * algorithms
 * Easy (68.33%)
 * Total Accepted:    29.2K
 * Total Submissions: 40.9K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你个整数数组 arr，其中每个元素都 不相同。
 *
 * 请你找到所有具有最小绝对差的元素对，并且按升序的顺序返回。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [4,2,1,3]
 * 输出：[[1,2],[2,3],[3,4]]
 *
 *
 * 示例 2：
 *
 * 输入：arr = [1,3,6,10,15]
 * 输出：[[1,3]]
 *
 *
 * 示例 3：
 *
 * 输入：arr = [3,8,-10,23,19,-4,-14,27]
 * 输出：[[-14,-10],[19,23],[23,27]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= arr.length <= 10^5
 * -10^6 <= arr[i] <= 10^6
 *
 *
 */

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var ans [][]int
	minAbs := math.MaxInt32
	for i, n := 0, len(arr); i+1 < n; i++ {
		if d := arr[i+1] - arr[i]; d < minAbs {
			minAbs = d
			ans = [][]int{{arr[i], arr[i+1]}}
		} else if d == minAbs {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return ans
}
