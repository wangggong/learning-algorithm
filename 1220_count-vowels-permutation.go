/*
 * @lc app=leetcode.cn id=count-vowels-permutation lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1220] 统计元音字母序列的数目
 *
 * https://leetcode-cn.com/problems/count-vowels-permutation/description/
 *
 * algorithms
 * Hard (60.83%)
 * Total Accepted:    24.6K
 * Total Submissions: 40.4K
 * Testcase Example:  '1'
 *
 * 给你一个整数 n，请你帮忙统计一下我们可以按下述规则形成多少个长度为 n 的字符串：
 *
 *
 * 字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
 * 每个元音 'a' 后面都只能跟着 'e'
 * 每个元音 'e' 后面只能跟着 'a' 或者是 'i'
 * 每个元音 'i' 后面 不能 再跟着另一个 'i'
 * 每个元音 'o' 后面只能跟着 'i' 或者是 'u'
 * 每个元音 'u' 后面只能跟着 'a'
 *
 *
 * 由于答案可能会很大，所以请你返回 模 10^9 + 7 之后的结果。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 1
 * 输出：5
 * 解释：所有可能的字符串分别是："a", "e", "i" , "o" 和 "u"。
 *
 *
 * 示例 2：
 *
 * 输入：n = 2
 * 输出：10
 * 解释：所有可能的字符串分别是："ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" 和
 * "ua"。
 *
 *
 * 示例 3：
 *
 * 输入：n = 5
 * 输出：68
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 2 * 10^4
 *
 *
 */
const VOWELS = 5
const MAXN = 2e4 + 5
const MOD = 1e9 + 7

/*
 * var dp [VOWELS][MAXN]int
 *
 * func countVowelPermutation(n int) int {
 * 	for i := 0; i < VOWELS; i++ {
 * 		dp[i][0] = 1
 * 		for j := 1; j < n; j++ {
 * 			dp[i][j] = 0
 * 		}
 * 	}
 * 	for i := 1; i < n; i++ {
 * 		dp[0][i] = (dp[1][i-1] + dp[2][i-1] + dp[4][i-1]) % MOD // 'a'
 * 		dp[1][i] = (dp[0][i-1] + dp[2][i-1]) % MOD              // 'e'
 * 		dp[2][i] = (dp[1][i-1] + dp[3][i-1]) % MOD              // 'i'
 * 		dp[3][i] = (dp[2][i-1]) % MOD                           // 'o'
 * 		dp[4][i] = (dp[2][i-1] + dp[3][i-1]) % MOD              // 'u'
 * 	}
 * 	ans := 0
 * 	for i := 0; i < VOWELS; i++ {
 * 		ans = (ans + dp[i][n-1]) % MOD
 * 	}
 * 	return ans
 * }
 */

type matrix [VOWELS][VOWELS]int

func (a matrix) mul(b matrix) matrix {
	c := matrix{}
	for i := 0; i < VOWELS; i++ {
		for j := 0; j < VOWELS; j++ {
			for k := 0; k < VOWELS; k++ {
				c[i][j] = (c[i][j] + a[i][k]*b[k][j]) % MOD
			}
		}
	}
	return c
}

func (a matrix) pow(n int) matrix {
	res := matrix{}
	for i := 0; i < VOWELS; i++ {
		res[i][i] = 1
	}
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

// 矩阵快速幂 -- 模版
func countVowelPermutation(n int) int {
	m := matrix{
		{0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0},
	}
	ans := 0
	res := m.pow(n - 1)
	for i := 0; i < VOWELS; i++ {
		for j := 0; j < VOWELS; j++ {
			ans = (ans + res[i][j]) % MOD
		}
	}
	return ans
}
