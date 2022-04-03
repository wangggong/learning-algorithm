/*
 * @lc app=leetcode.cn id=remove-duplicate-letters lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [316] 去除重复字母
 *
 * https://leetcode-cn.com/problems/remove-duplicate-letters/description/
 *
 * algorithms
 * Medium (47.71%)
 * Total Accepted:    74.6K
 * Total Submissions: 156.3K
 * Testcase Example:  '"bcabc"'
 *
 * 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
 *
 * 注意：该题与 1081
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters
 * 相同
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bcabc"
 * 输出："abc"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbacdcbc"
 * 输出："acdb"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1e4
 * s 由小写英文字母组成
 *
 *
 */
// import "fmt"

/*
 * func removeDuplicateLetters(s string) string {
 * 	b := []byte(s)
 * 	n := len(b)
 *
 * 	var C [26]int
 * 	var V [26]bool
 *
 * 	for _, c := range b {
 * 		C[c-'a']++
 * 	}
 *
 * 	S := make([]byte, 0, n)
 * 	for _, c := range b {
 * 		// fmt.Printf("%s\n", S)
 * 		if !V[c-'a'] {
 * 			for len(S) > 0 && S[len(S)-1] > c {
 * 				if C[S[len(S)-1]-'a'] > 0 {
 * 					V[S[len(S)-1]-'a'] = false
 * 					S = S[:len(S)-1]
 * 				} else {
 * 					break
 * 				}
 * 			}
 * 			V[c-'a'] = true
 * 			S = append(S, c)
 * 		}
 * 		C[c-'a']--
 * 	}
 *
 * 	return string(S)
 * }
 */

// 贪心: 欠债表模型
func greedy(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	m := make([]int, 26)
	for _, c := range b {
		m[c-'a']++
	}
	index, minChar := 0, b[0]
	for i, c := range b {
		if c < minChar {
			index, minChar = i, c
		}
		m[c-'a']--
		if m[c-'a'] == 0 {
			break
		}
	}
	nb := make([]byte, 0, len(b))
	for _, c := range b[index+1:] {
		if c == minChar {
			continue
		}
		nb = append(nb, c)
	}
	ans := make([]byte, 0, len(b))
	ans = append(ans, minChar)
	ans = append(ans, greedy(nb)...)
	return ans
}

func removeDuplicateLetters(s string) string {
	b := []byte(s)
	return string(greedy(b))
}
