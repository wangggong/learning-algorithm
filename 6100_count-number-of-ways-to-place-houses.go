/*
 * @lc app=leetcode.cn id=count-number-of-ways-to-place-houses lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6100] 统计放置房子的方式数
 *
 * https://leetcode.cn/problems/count-number-of-ways-to-place-houses/description/
 *
 * algorithms
 * Medium (36.13%)
 * Total Accepted:    5.1K
 * Total Submissions: 13.8K
 * Testcase Example:  '1'
 *
 * 一条街道上共有 n * 2 个 地块 ，街道的两侧各有 n 个地块。每一边的地块都按从 1 到 n 编号。每个地块上都可以放置一所房子。
 *
 * 现要求街道同一侧不能存在两所房子相邻的情况，请你计算并返回放置房屋的方式数目。由于答案可能很大，需要对 10^9 + 7 取余后再返回。
 *
 * 注意，如果一所房子放置在这条街某一侧上的第 i 个地块，不影响在另一侧的第 i 个地块放置房子。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 1
 * 输出：4
 * 解释：
 * 可能的放置方式：
 * 1. 所有地块都不放置房子。
 * 2. 一所房子放在街道的某一侧。
 * 3. 一所房子放在街道的另一侧。
 * 4. 放置两所房子，街道两侧各放置一所。
 *
 *
 * 示例 2：
 *
 * 输入：n = 2
 * 输出：9
 * 解释：如上图所示，共有 9 种可能的放置方式。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */
const mod int = 1e9 + 7

func countHousePlacements(n int) int {
	dp := make([][4]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			for j := 0; j < 4; j++ {
				dp[i][j] = 1
			}
			continue
		}
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if j&k != 0 {
					continue
				}
				dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
			}
		}
	}
	ans := 0
	for i := 0; i < 4; i++ {
		ans = (ans + dp[n-1][i]) % mod
	}
	return ans
}
