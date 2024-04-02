/*
 * @lc app=leetcode.cn id=maximum-length-of-pair-chain lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [646] 最长数对链
 *
 * https://leetcode.cn/problems/maximum-length-of-pair-chain/description/
 *
 * algorithms
 * Medium (58.58%)
 * Total Accepted:    45.9K
 * Total Submissions: 75.9K
 * Testcase Example:  '[[1,2],[2,3],[3,4]]'
 *
 * 给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
 *
 * 现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
 *
 * 给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：[[1,2], [2,3], [3,4]]
 * 输出：2
 * 解释：最长的数对链是 [1,2] -> [3,4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给出数对的个数在 [1, 1000] 范围内。
 *
 *
 */
import (
	"math"
	"sort"
)

/*
 * func findLongestChain(pairs [][]int) (ans int) {
 * 	n := len(pairs)
 * 	dp := make([]int, n)
 * 	sort.Slice(pairs, func(p, q int) bool { return pairs[p][1] < pairs[q][1] })
 * 	for i := 0; i < n; i++ {
 * 		dp[i] = 1
 * 		for j := 0; j < i && pairs[j][1] < pairs[i][0]; j++ {
 * 			dp[i] = max(dp[i], dp[j]+1)
 * 		}
 * 		ans = max(ans, dp[i])
 * 	}
 * 	return
 * }
 *
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 */

func findLongestChain(pairs [][]int) (ans int) {
	sort.Slice(pairs, func(p, q int) bool { return pairs[p][1] < pairs[q][1] })
	cur := math.MinInt32
	for _, p := range pairs {
		if cur < p[0] {
			cur, ans = p[1], ans+1
		}
	}
	return
}
