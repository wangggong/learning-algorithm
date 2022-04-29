/*
 * @lc app=leetcode.cn id=rotate-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [796] 旋转字符串
 *
 * https://leetcode-cn.com/problems/rotate-string/description/
 *
 * algorithms
 * Easy (55.00%)
 * Total Accepted:    40K
 * Total Submissions: 67.9K
 * Testcase Example:  '"abcde"\n"cdeab"'
 *
 * 给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。
 *
 * s 的 旋转操作 就是将 s 最左边的字符移动到最右边。
 *
 *
 * 例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcde", goal = "cdeab"
 * 输出: true
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "abcde", goal = "abced"
 * 输出: false
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length, goal.length <= 100
 * s 和 goal 由小写英文字母组成
 *
 *
 */
/*
 * func rotateString(s string, goal string) bool {
 * 	arr := []byte(s)
 * 	n := len(arr)
 * 	if len(goal) != n {
 * 		return false
 * 	}
 * 	if s == goal {
 * 		return true
 * 	}
 * 	for i := 1; i < n; i++ {
 * 		var curr []byte
 * 		curr = append(curr, s[i:]...)
 * 		curr = append(curr, s[:i]...)
 * 		if goal == string(curr) {
 * 			return true
 * 		}
 * 	}
 * 	return false
 * }
 */

const mod int = 1e9 + 7
const base int = 53

/*
 * func rotateString(s, goal string) bool {
 * 	arr := []byte(s)
 * 	n := len(arr)
 * 	if n != len(goal) {
 * 		return false
 * 	}
 * 	if s == goal {
 * 		return true
 * 	}
 * 	prefix := make([]int, n+1)
 * 	mul := make([]int, n+1)
 * 	mul[0] = 1
 * 	for i := 0; i < n; i++ {
 * 		prefix[i+1] = (prefix[i]*base + int(arr[i]-'a')) % mod
 * 		mul[i+1] = (mul[i] * base) % mod
 * 	}
 * 	hash := 0
 * 	for i := 0; i < n; i++ {
 * 		hash = (hash*base + int(goal[i]-'a')) % mod
 * 	}
 * 	for i := 1; i < n; i++ {
 * 		h := ((prefix[n]-prefix[i]*mul[n-i]%mod+mod)*mul[i]%mod + prefix[i]) % mod
 * 		if h == hash {
 * 			return true
 * 		}
 * 	}
 * 	return false
 * }
 */

func rotateString(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return kmp(s+s, goal) >= 0
}

func kmp(haystack, needle string) int {
	hs, nd := []byte(haystack), []byte(needle)
	n, m := len(hs), len(nd)
	if m == 0 {
		return 0
	}
	dfa := make([][26]int, m)
	func(arr []byte) {
		m := len(arr)
		dfa[0][int(arr[0]-'a')] = 1
		for X, j := 0, 1; j < m; j++ {
			for b := 'a'; b <= 'z'; b++ {
				dfa[j][int(b-'a')] = dfa[X][int(b-'a')]
				dfa[j][int(arr[j]-'a')] = j + 1
				X = dfa[X][int(arr[j]-'a')]
			}
		}
	}(nd)
	for i, j := 0, 0; i < n; i++ {
		j = dfa[j][int(hs[i]-'a')]
		if j == m {
			return i
		}
	}
	return -1
}
