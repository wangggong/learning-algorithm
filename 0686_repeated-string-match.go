/*
 * @lc app=leetcode.cn id=repeated-string-match lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [686] 重复叠加字符串匹配
 *
 * https://leetcode-cn.com/problems/repeated-string-match/description/
 *
 * algorithms
 * Medium (40.16%)
 * Total Accepted:    43.3K
 * Total Submissions: 107.9K
 * Testcase Example:  '"abcd"\n"cdabcdab"'
 *
 * 给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。
 *
 * 注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。
 *
 *
 *
 * 示例 1：
 *
 * 输入：a = "abcd", b = "cdabcdab"
 * 输出：3
 * 解释：a 重复叠加三遍后为 "abcdabcdabcd", 此时 b 是其子串。
 *
 *
 * 示例 2：
 *
 * 输入：a = "a", b = "aa"
 * 输出：2
 *
 *
 * 示例 3：
 *
 * 输入：a = "a", b = "a"
 * 输出：1
 *
 *
 * 示例 4：
 *
 * 输入：a = "abc", b = "wxyz"
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= a.length <= 10^4
 * 1 <= b.length <= 10^4
 * a 和 b 由小写英文字母组成
 *
 *
 */

const MAXN = 5e4 + 10

// const MAXN = 8

var dfa [26][MAXN]int
var n, m int

func repeatedStringMatch(a string, b string) int {
	ba, bb := []byte(a), []byte(b)
	n, m = len(ba), len(bb)
	if m == 0 {
		return 1
	}
	var cur []byte
	ans := 0
	for len(cur) < m {
		cur = append(cur, ba...)
		ans++
	}
	cur = append(cur, ba...)
	index := strStr(cur, bb)
	if index == -1 {
		return -1
	}
	if index+m <= ans * n{
		return ans
	}
	return ans + 1
	/*
		if m <= n-index {
			return 1
		}
		// NOTE: ceil(p / q) = (p - 1) / q + 1
		return ((m-(n-index)-1)/n + 1) + 1
	*/
}

/*
 * // KMP
 * func strStr(text, pattern []byte) int {
 * 	for ch := 'a'; ch <= 'z'; ch++ {
 * 		for j := 0; j < MAXN; j++ {
 * 			dfa[int(ch-'a')][j] = 0
 * 		}
 * 	}
 * 	buildDFA(pattern)
 * 	i, j := 0, 0
 * 	for ; i < n+j && j < m; i++ {
 * 		j = dfa[int(text[i%n]-'a')][j]
 * 	}
 * 	if j == m {
 * 		return i - j
 * 	}
 * 	return -1
 * }
 */

// BM
func strStr(text, pattern []byte) int {
	n := len(text) // evil enough
	var right [26]int
	for i := range right {
		right[i] = -1
	}
	for i, p := range pattern {
		right[int(p-'a')] = i
	}
	var i, j, skip int
	// NOTE: 注意终止条件: i + j < n => i <= n-m
	for i = 0; i <= n-m; i += skip {
		skip = 0
		for j = m - 1; j >= 0; j-- {
			ch := text[i+j]
			if ch == pattern[j] {
				continue
			}
			skip = max(j-right[int(ch-'a')], 1)
			break
		}
		if skip == 0 {
			return i
		}
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func buildDFA(pattern []byte) {
	// assert len(pattern) > 0;
	dfa[int(pattern[0]-'a')][0] = 1
	for X, j := 0, 1; j < m; j++ {
		for ch := 'a'; ch <= 'z'; ch++ {
			dfa[int(ch-'a')][j] = dfa[int(ch-'a')][X]
		}
		dfa[int(pattern[j]-'a')][j] = j + 1
		X = dfa[int(pattern[j]-'a')][X]
	}
}
