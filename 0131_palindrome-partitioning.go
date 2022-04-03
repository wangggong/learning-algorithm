/*
 * @lc app=leetcode.cn id=palindrome-partitioning lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (72.43%)
 * Total Accepted:    154.2K
 * Total Submissions: 212.8K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 *
 * 回文串 是正着读和反着读都一样的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a"
 * 输出：[["a"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 16
 * s 仅由小写英文字母组成
 *
 *
 */
import "fmt"

type S struct{ p, q int }

var ans [][]string
var cur []string

func partition(s string) [][]string {
	ans, cur = nil, nil
	b := []byte(s)
	n := len(b)
	m := make(map[S]bool)
	dfs(b, 0, n, m)
	return ans
}

func dfs(b []byte, c, n int, m map[S]bool) {
	if c == n {
		v := make([]string, len(cur))
		copy(v, cur)
		ans = append(ans, v)
		// fmt.Printf("%v\n", ans)
		return
	}

	for i := c + 1; i <= n; i++ {
		if !isPalindrome(b, c, i, m) {
			continue
		}
		cur = append(cur, string(b[c:i]))
		dfs(b, i, n, m)
		cur = cur[:len(cur)-1]
	}
}

func isPalindrome(b []byte, p, q int, m map[S]bool) bool {
	if v, ok := m[S{p, q}]; ok {
		return v
	}
	v := false
	defer func() {
		m[S{p, q}] = v
	}()
	if p >= q {
		v = p == q
		return v
	}
	if q-p == 1 {
		v = true
		return v
	}
	if b[q-1] != b[p] {
		return v
	}
	v = isPalindrome(b, p+1, q-1, m)
	return v
}
