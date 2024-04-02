/*
 * @lc app=leetcode.cn id=matchsticks-to-square lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [473] 火柴拼正方形
 *
 * https://leetcode.cn/problems/matchsticks-to-square/description/
 *
 * algorithms
 * Medium (42.25%)
 * Total Accepted:    34.8K
 * Total Submissions: 80.2K
 * Testcase Example:  '[1,1,2,2,2]'
 *
 * 你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。你要用 所有的火柴棍 拼成一个正方形。你
 * 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
 *
 * 如果你能使这个正方形，则返回 true ，否则返回 false 。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: matchsticks = [1,1,2,2,2]
 * 输出: true
 * 解释: 能拼成一个边长为2的正方形，每边两根火柴。
 *
 *
 * 示例 2:
 *
 *
 * 输入: matchsticks = [3,3,3,3,4]
 * 输出: false
 * 解释: 不能用所有火柴拼成一个正方形。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= matchsticks.length <= 15
 * 1 <= matchsticks[i] <= 10^8
 *
 *
 */
func makesquare(M []int) bool {
	tot := 0
	for _, m := range M {
		tot += m
	}
	if tot%4 != 0 {
		return false
	}
	return backtrack(M, make(map[[3]int]bool), len(M), 0, 4, tot/4, 0)
}

func backtrack(M []int, memo map[[3]int]bool, n, visit, k, target, curr int) (ans bool) {
	s := [3]int{visit, k, curr}
	if ans, ok := memo[s]; ok {
		return ans
	}
	defer func() { memo[s] = ans }()
	if visit == (1<<n)-1 {
		ans = k == 0
		return
	}
	for i := 0; i < n; i++ {
		if visit&(1<<i) != 0 {
			continue
		}
		if curr+M[i] > target {
			continue
		}
		if curr+M[i] == target {
			if backtrack(M, memo, n, visit|(1<<i), k-1, target, 0) {
				ans = true
				return
			}
		} else {
			if backtrack(M, memo, n, visit|(1<<i), k, target, curr+M[i]) {
				ans = true
				return
			}
		}
	}
	return false
}
