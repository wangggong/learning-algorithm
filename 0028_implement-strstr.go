/*
 * @lc app=leetcode.cn id=implement-strstr lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [28] 实现 strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (40.28%)
 * Total Accepted:    560.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0
 * 开始）。如果不存在，则返回  -1 。
 *
 *
 *
 * 说明：
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf()
 * 定义相符。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：haystack = "hello", needle = "ll"
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：haystack = "aaaaa", needle = "bba"
 * 输出：-1
 *
 *
 * 示例 3：
 *
 *
 * 输入：haystack = "", needle = ""
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= haystack.length, needle.length <= 5 * 1e4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */

const MAXN = 5e4 + 10
const Q = 15485863 // 恶臭未遂
// const Q = 139 // 恶臭未遂

var dfa [26][MAXN]int
var right [26]int

/*
 * // brute force
 * func strStr(haystack string, needle string) int {
 * 	bh, bn := []byte(haystack), []byte(needle)
 * 	n, m := len(bh), len(bn)
 * 	for i := 0; i <= n-m; i++ {
 * 		j := 0
 * 		for ; j < m; j++ {
 * 			if bh[i+j] != bn[j] {
 * 				break
 * 			}
 * 		}
 * 		if j == m {
 * 			return i
 * 		}
 * 	}
 * 	return -1
 * }
 */

/*
 * // brute force #2
 * func strStr(haystack string, needle string) int {
 * 	bh, bn := []byte(haystack), []byte(needle)
 * 	n, m := len(bh), len(bn)
 * 	i, j := 0, 0
 * 	for ; i < n && j < m; i++ {
 * 		if bh[i] == bn[j] {
 * 			j++
 * 		} else {
 * 			i, j = i-j, 0
 * 		}
 * 	}
 * 	if j == m {
 * 		return i-j
 * 	}
 * 	return -1
 * }
 */

/*
 * // KMP
 *
 * func strStr(haystack string, needle string) int {
 * 	bh, bn := []byte(haystack), []byte(needle)
 * 	n, m := len(bh), len(bn)
 * 	if m == 0 {
 * 		return 0
 * 	}
 * 	for ch := 0; ch < 26; ch++ {
 * 		for j := 0; j < MAXN; j++ {
 * 			dfa[ch][j] = 0
 * 		}
 * 	}
 * 	buildDFA(bh, bn, n, m)
 * 	// fmt.Printf("%v\n", dfa)
 * 	i, j := 0, 0
 * 	for ; i < n && j < m; i++ {
 * 		j = dfa[int(bh[i]-'a')][j]
 * 	}
 * 	if j == m {
 * 		return i - j
 * 	}
 * 	return -1
 * }
 *
 */

/*
 * // BM (Boyer-Moore)
 * //
 * // 在暴力的基础上添加了反向遍历与启发式查找.
 * // **据说** 使用 buildDFA 能提速.
 * func strStr(haystack string, needle string) int {
 * 	bh, bn := []byte(haystack), []byte(needle)
 * 	n, m := len(bh), len(bn)
 * 	for i := range right {
 * 		right[i] = -1
 * 	}
 * 	for i, b := range bn {
 * 		right[int(b-'a')] = i
 * 	}
 * 	var skip int
 * 	for i := 0; i <= n-m; i += skip {
 * 		skip = 0
 * 		for j := m - 1; j >= 0; j-- {
 * 			if bh[i+j] == bn[j] {
 * 				continue
 * 			}
 * 			skip = j - right[int(bh[i+j]-'a')]
 * 			if skip < 1 {
 * 				skip = 1
 * 			}
 * 			break
 * 		}
 * 		if skip == 0 {
 * 			return i
 * 		}
 * 	}
 * 	return -1
 * }
 */

// RK (Rabin-Karp)
//
// 直接蒙特卡洛, 可能出错, 但巨快.
func strStr(haystack string, needle string) int {
	bh, bn := []byte(haystack), []byte(needle)
	n, m := len(bh), len(bn)
	RM := int64(1)
	for i := 1; i <= m-1; i++ {
		RM = (26 * RM) % Q
	}
	if n < m {
		return -1
	}
	hbn := hash(bn, m)
	hbh := hash(bh, m)
	if hbn == hbh && check(0) {
		return 0
	}
	for i := m; i < n; i++ {
		hbh = (hbh + Q - (int64(bh[i-m]-'a')*RM)%Q) % Q
		hbh = (hbh*26 + int64(bh[i]-'a')) % Q
		if hbh == hbn && check(i-m+1) {
			return i - m + 1
		}
	}
	return -1
}

func buildDFA(arr1, arr2 []byte, n, m int) {
	// assert m > 0;
	dfa[int(arr2[0]-'a')][0] = 1
	for X, j := 0, 1; j < m; j++ {
		for ch := 'a'; ch <= 'z'; ch++ {
			dfa[int(ch-'a')][j] = dfa[int(ch-'a')][X]
		}
		dfa[int(arr2[j]-'a')][j] = j + 1
		X = dfa[int(arr2[j]-'a')][X]
	}
}

// Horner 方法.
func hash(arr []byte, n int) int64 {
	h := int64(0)
	for i := 0; i < n; i++ {
		a := arr[i]
		h = (h*26 + int64(a-'a')) % Q
	}
	return h
}

func check(x int) bool { return true }
