import "strings"

/*
 * @lc app=leetcode.cn id=repeated-substring-pattern lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [459] 重复的子字符串
 *
 * https://leetcode-cn.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (50.95%)
 * Total Accepted:    94.2K
 * Total Submissions: 184.9K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abab"
 * 输出: true
 * 解释: 可由子串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "aba"
 * 输出: false
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "abcabcabcabc"
 * 输出: true
 * 解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 10^4
 * s 由小写英文字母组成
 *
 *
 */
func repeatedSubstringPattern(s string) bool {
	// bytes := []byte(s)
	// n := len(bytes)
	// if n == 0 {
	// 	return true
	// }
	// var indexes []int
	// b := bytes[0]
	// for i := 1; i * 2 <= n; i++ {
	// 	if bytes[i] == b && n % i == 0 {
	// 		indexes = append(indexes, i)
	// 	}
	// }
	// for _, index := range indexes {
	// 	allTheSame := true
	// 	origin := string(bytes[:index])
	// 	for i := index; i < n; i += index {
	// 		curr := string(bytes[i:i+index])
	// 		if origin != curr {
	// 			allTheSame = false
	// 			break
	// 		}
	// 	}
	// 	if allTheSame {
	// 		return true
	// 	}
	// }
	// return false

	tmp := []byte(s + s)
	ns := string(tmp[1:])
	return strings.Index(ns, s) < len(s)-1

	// 当然也可以手撕 KMP
}

/*
 * func repeatedSubstringPattern(s string) bool {
 *     return kmp(s + s, s)
 * }
 *
 * func kmp(query, pattern string) bool {
 *     n, m := len(query), len(pattern)
 *     fail := make([]int, m)
 *     for i := 0; i < m; i++ {
 *         fail[i] = -1
 *     }
 *     for i := 1; i < m; i++ {
 *         j := fail[i - 1]
 *         for j != -1 && pattern[j + 1] != pattern[i] {
 *             j = fail[j]
 *         }
 *         if pattern[j + 1] == pattern[i] {
 *             fail[i] = j + 1
 *         }
 *     }
 *     match := -1
 *     for i := 1; i < n - 1; i++ {
 *         for match != -1 && pattern[match + 1] != query[i] {
 *             match = fail[match]
 *         }
 *         if pattern[match + 1] == query[i] {
 *             match++
 *             if match == m - 1 {
 *                 return true
 *             }
 *         }
 *     }
 *     return false
 * }
 */
