/*
 * @lc app=leetcode.cn id=shortest-palindrome lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [214] 最短回文串
 *
 * https://leetcode-cn.com/problems/shortest-palindrome/description/
 *
 * algorithms
 * Hard (37.65%)
 * Total Accepted:    34.9K
 * Total Submissions: 92.4K
 * Testcase Example:  '"aacecaaa"'
 *
 * 给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aacecaaa"
 * 输出："aaacecaaa"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcd"
 * 输出："dcbabcd"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 1e4
 * s 仅由小写英文字母组成
 *
 *
 */

var Q int = 1e9 + 7

func shortestPalindrome(s string) string {
	arr := []byte(s)
	n := len(arr)
	rev := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		rev[i] = arr[n-1-i]
	}
	m := make(map[int64]int)
	h := hash(arr[:1])
	m[h] = 1
	for i := 2; i <= n; i++ {
		h = (h*26 + int64(arr[i-1]-'a')) % int64(Q)
		m[h] = i
	}
	k := 0
	h = int64(0)
	v := int64(1)
	for i := n - 1; i >= 0; i-- {
		h = (h + int64(rev[i]-'a')*v) % int64(Q)
		v *= 26
		if c, ok := m[h]; ok {
			k = c
		}
	}
	rev = append(rev, arr[k:]...)
	return string(rev)
}

func hash(arr []byte) int64 {
	n := len(arr)
	h := int64(0)
	for i := 0; i < n; i++ {
		h = (h*26 + int64(arr[i]-'a')) % int64(Q)
	}
	return h
}
