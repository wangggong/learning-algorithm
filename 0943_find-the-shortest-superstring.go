/*
 * @lc app=leetcode.cn id=find-the-shortest-superstring lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [943] 最短超级串
 *
 * https://leetcode-cn.com/problems/find-the-shortest-superstring/description/
 *
 * algorithms
 * Hard (45.50%)
 * Total Accepted:    2.2K
 * Total Submissions: 4.9K
 * Testcase Example:  '["alex","loves","leetcode"]'
 *
 * 给定一个字符串数组 words，找到以 words 中每个字符串作为子字符串的最短字符串。如果有多个有效最短字符串满足题目条件，返回其中 任意一个
 * 即可。
 *
 * 我们可以假设 words 中没有字符串是 words 中另一个字符串的子字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["alex","loves","leetcode"]
 * 输出："alexlovesleetcode"
 * 解释："alex"，"loves"，"leetcode" 的所有排列都会被接受。
 *
 * 示例 2：
 *
 *
 * 输入：words = ["catg","ctaagt","gcta","ttca","atgcatc"]
 * 输出："gctaagttcatgcatc"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= words.length <= 12
 * 1 <= words[i].length <= 20
 * words[i] 由小写英文字母组成
 * words 中的所有字符串 互不相同
 *
 *
 */

const MAXN int = 12

var overlap [MAXN + 5][MAXN + 5]int
var dp [(1 << MAXN) + 5][MAXN + 5]string

// 状态压缩+DP. 压根没想过这个方向.
func shortestSuperstring(words []string) string {
	n := len(words)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			overlap[i][j] = getOverlap(words[i], words[j])
		}
	}
	var ans string
	for _, w := range words {
		ans = ans + w
	}
	for i := 0; i < (1<<MAXN)+5; i++ {
		for j := 0; j < MAXN+5; j++ {
			dp[i][j] = ans
		}
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = words[i]
	}
	states := 1 << n
	// 永远在最外层枚举每个状态, 傻逼!
	for s := 0; s < states; s++ {
		for i := 0; i < n; i++ {
			if (s>>i)|1 == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				// a | 1 != 0 恒成立, 傻逼
				if (s>>j)&1 != 0 {
					continue
				}
				str := dp[s][i] + words[j][overlap[i][j]:]
				if len(str) < len(dp[s|(1<<j)][j]) {
					dp[s|(1<<j)][j] = str
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if len(ans) > len(dp[states-1][i]) {
			ans = dp[states-1][i]
		}
	}
	return ans
}

func getOverlap(s, t string) int {
	bs, bt := []byte(s), []byte(t)
	n, m := len(bs), len(bt)
	ans := 0
	for i := 0; i <= m && i <= n; i++ {
		if string(s[n-i:]) == string(t[:i]) {
			ans = i
		}
	}
	return ans
}
