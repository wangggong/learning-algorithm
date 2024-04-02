/*
 * @lc app=leetcode.cn id=number-of-distinct-roll-sequences lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6107] 不同骰子序列的数目
 *
 * https://leetcode.cn/problems/number-of-distinct-roll-sequences/description/
 *
 * algorithms
 * Hard (52.55%)
 * Total Accepted:    1.6K
 * Total Submissions: 3K
 * Testcase Example:  '4'
 *
 * 给你一个整数 n 。你需要掷一个 6 面的骰子 n 次。请你在满足以下要求的前提下，求出 不同 骰子序列的数目：
 *
 *
 * 序列中任意 相邻 数字的 最大公约数 为 1 。
 * 序列中 相等 的值之间，至少有 2 个其他值的数字。正式地，如果第 i 次掷骰子的值 等于 第 j 次的值，那么 abs(i - j) > 2 。
 *
 *
 * 请你返回不同序列的 总数目 。由于答案可能很大，请你将答案对 10^9 + 7 取余 后返回。
 *
 * 如果两个序列中至少有一个元素不同，那么它们被视为不同的序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：184
 * 解释：一些可行的序列为 (1, 2, 3, 4) ，(6, 1, 2, 3) ，(1, 2, 3, 1) 等等。
 * 一些不可行的序列为 (1, 2, 1, 3) ，(1, 2, 3, 6) 。
 * (1, 2, 1, 3) 是不可行的，因为第一个和第三个骰子值相等且 abs(1 - 3) = 2 （下标从 1 开始表示）。
 * (1, 2, 3, 6) i是不可行的，因为 3 和 6 的最大公约数是 3 。
 * 总共有 184 个不同的可行序列，所以我们返回 184 。
 *
 * 示例 2：
 *
 *
 * 输入：n = 2
 * 输出：22
 * 解释：一些可行的序列为 (1, 2) ，(2, 1) ，(3, 2) 。
 * 一些不可行的序列为 (3, 6) ，(2, 4) ，因为最大公约数不为 1 。
 * 总共有 22 个不同的可行序列，所以我们返回 22 。
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

// 参考: https://leetcode.cn/problems/number-of-distinct-roll-sequences/solution/by-endlesscheng-tgkn/
//
// 其实没那么难的 DP, 但如果用二维的状态转移确实不好处理.
const mod int = 1e9 + 7

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func distinctSequences(n int) int {
	if n == 1 {
		return 6
	}
	dp := make([][7][7]int, n)
	for i := 1; i <= 6; i++ {
		dp[0][i][0] = 1
	}
	for i := 1; i < n; i++ {
		if i == 1 {
			for j := 1; j <= 6; j++ {
				for k := 1; k <= 6; k++ {
					if j != k && gcd(j, k) == 1 {
						dp[i][j][k] = 1
					}
				}
			}
			continue
		}
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				for l := 1; l <= 6; l++ {
					if j != k && j != l && gcd(j, k) == 1 && k != l && gcd(k, l) == 1 {
						dp[i][j][k] = (dp[i][j][k] + dp[i-1][k][l]) % mod
					}
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			ans = (ans + dp[n-1][i][j]) % mod
		}
	}
	return ans
}



// 参考: https://leetcode.cn/problems/number-of-distinct-roll-sequences/solution/pythonjs-by-981377660lmt-8fud/
//
// 问题的关键是对于诸如 `1 3 2 1` 的场景的处理. 你会发现这种东西算是 **减多了的**; 在计算 `dp[i][j]` 的时候就不如把 `dp[i-1][j]` **回滚**回没有减掉 `dp[i-2][j]` 的状态.
func distinctSequences(n int) int {
	dp := make([][7]int, n)
	for i := 1; i <= 6; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		if i > 2 {
			for j := 1; j <= 6; j++ {
				dp[i-1][j] = (dp[i-1][j] + dp[i-3][j]) % mod
			}
		}
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				if gcd(j, k) == 1 && j != k {
					dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
					if i > 1 {
						dp[i][j] = ((dp[i][j]-dp[i-2][j])%mod + mod) % mod
					}
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= 6; i++ {
		ans = (ans + dp[n-1][i]) % mod
	}
	return ans
}

