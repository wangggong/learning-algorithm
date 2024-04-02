/*
 * @lc app=leetcode.cn id=height-checker lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1051] 高度检查器
 *
 * https://leetcode.cn/problems/height-checker/description/
 *
 * algorithms
 * Easy (75.97%)
 * Total Accepted:    44K
 * Total Submissions: 56.4K
 * Testcase Example:  '[1,1,4,2,1,3]'
 *
 * 学校打算为全体学生拍一张年度纪念照。根据要求，学生需要按照 非递减 的高度顺序排成一行。
 *
 * 排序后的高度情况用整数数组 expected 表示，其中 expected[i] 是预计排在这一行中第 i 位的学生的高度（下标从 0 开始）。
 *
 * 给你一个整数数组 heights ，表示 当前学生站位 的高度情况。heights[i] 是这一行中第 i 位学生的高度（下标从 0 开始）。
 *
 * 返回满足 heights[i] != expected[i] 的 下标数量 。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：heights = [1,1,4,2,1,3]
 * 输出：3
 * 解释：
 * 高度：[1,1,4,2,1,3]
 * 预期：[1,1,1,2,3,4]
 * 下标 2 、4 、5 处的学生高度不匹配。
 *
 * 示例 2：
 *
 *
 * 输入：heights = [5,1,2,3,4]
 * 输出：5
 * 解释：
 * 高度：[5,1,2,3,4]
 * 预期：[1,2,3,4,5]
 * 所有下标的对应学生高度都不匹配。
 *
 * 示例 3：
 *
 *
 * 输入：heights = [1,2,3,4,5]
 * 输出：0
 * 解释：
 * 高度：[1,2,3,4,5]
 * 预期：[1,2,3,4,5]
 * 所有下标的对应学生高度都匹配。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= heights.length <= 100
 * 1 <= heights[i] <= 100
 *
 *
 */

/*
 * import "sort"
 *
 * func heightChecker(heights []int) int {
 * 	n := len(heights)
 * 	arr := make([]int, n)
 * 	copy(arr, heights)
 * 	sort.Ints(arr)
 * 	ans := 0
 * 	for i := 0; i < n; i++ {
 * 		if heights[i] != arr[i] {
 * 			ans++
 * 		}
 * 	}
 * 	return ans
 * }
 */

// 计数排序.
const maxn int = 100

func heightChecker(heights []int) int {
	var cnt [maxn + 1]int
	for _, h := range heights {
		cnt[h]++
	}
	var hcnt []int
	for i, c := range cnt {
		for ; c > 0; c-- {
			hcnt = append(hcnt, i)
		}
	}
	ans := 0
	for i, n := 0, len(heights); i < n; i++ {
		if hcnt[i] != heights[i] {
			ans++
		}
	}
	return ans
}

/*
 * // 计数排序也可以写成这样节省空间:
 * func heightChecker(heights []int) int {
 * 	const maxn int = 100
 * 	var cnt [maxn + 1]int
 * 	for _, h := range heights {
 * 		cnt[h]++
 * 	}
 * 	ans := 0
 * 	for i, h := 0, 0; h <= maxn; h++ {
 * 		for c := cnt[h]; c > 0; c, i = c-1, i+1 {
 * 			if heights[i] != h {
 * 				ans++
 * 			}
 * 		}
 * 	}
 * 	return ans
 * }
 */
