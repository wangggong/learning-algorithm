/*
 * @lc app=leetcode.cn id=longest-well-performing-interval lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1124] 表现良好的最长时间段
 *
 * https://leetcode.cn/problems/longest-well-performing-interval/description/
 *
 * algorithms
 * Medium (33.32%)
 * Total Accepted:    16.1K
 * Total Submissions: 48.3K
 * Testcase Example:  '[9,9,6,0,6,6,9]'
 *
 * 给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
 *
 * 我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
 *
 * 所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
 *
 * 请你返回「表现良好时间段」的最大长度。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：hours = [9,9,6,0,6,6,9]
 * 输出：3
 * 解释：最长的表现良好时间段是 [9,9,6]。
 *
 * 示例 2：
 *
 *
 * 输入：hours = [6,6,6]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= hours.length <= 10^4
 * 0 <= hours[i] <= 16
 *
 *
 */

// 参考: https://leetcode.cn/problems/longest-well-performing-interval/solution/qian-zhui-he-dan-diao-zhan-python3-by-smoon1989/
func longestWPI(hours []int) int {
	n := len(hours)
	ps := make([]int, n+1)
	for i, h := range hours {
		v := -1
		if h > 8 {
			v = 1
		}
		ps[i+1] = ps[i] + v
	}
	var S []int
	for i, p := range ps {
		if len(S) == 0 || ps[S[len(S)-1]] > p {
			S = append(S, i)
		}
	}
	ans := 0
	for i := n; i > 0; i-- {
		for len(S) > 0 && ps[S[len(S)-1]] < ps[i] {
			ans = max(ans, i-S[len(S)-1])
			S = S[:len(S)-1]
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
 * // 查表:
 * func longestWPI(hours []int) int {
 * 	m := make(map[int]int)
 * 	ans := 0
 * 	cur := 0
 * 	for i, h := range hours {
 * 		if h > 8 {
 * 			cur++
 * 		} else {
 * 			cur--
 * 		}
 * 		if cur > 0 {
 * 			ans = i + 1
 * 			continue
 * 		}
 * 		if v, ok := m[cur-1]; ok {
 * 			ans = max(ans, i-v)
 * 		}
 * 		if _, ok := m[cur]; !ok {
 * 			m[cur] = i
 * 		}
 * 	}
 * 	return ans
 * }
 * 
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 */
