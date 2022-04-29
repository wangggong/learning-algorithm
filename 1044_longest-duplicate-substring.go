/*
 * @lc app=leetcode.cn id=longest-duplicate-substring lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1044] 最长重复子串
 *
 * https://leetcode-cn.com/problems/longest-duplicate-substring/description/
 *
 * algorithms
 * Hard (34.85%)
 * Total Accepted:    22.3K
 * Total Submissions: 63.9K
 * Testcase Example:  '"banana"'
 *
 * 给你一个字符串 s ，考虑其所有 重复子串 ：即 s 的（连续）子串，在 s 中出现 2 次或更多次。这些出现之间可能存在重叠。
 *
 * 返回 任意一个 可能具有最长长度的重复子串。如果 s 不含重复子串，那么答案为 "" 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "banana"
 * 输出："ana"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcd"
 * 输出：""
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= s.length <= 3 * 10^4
 * s 由小写英文字母组成
 *
 *
 */

// 官解就是双哈希, 还特么卡单哈希的. 逼良为娼啊...
const mod int = 1e9 + 7
const base1 int = 47
const base2 int = 541

var prefix1, prefix2 []int
var mul1, mul2 []int

func longestDupSubstring(s string) string {
	arr := []byte(s)
	n := len(arr)
	prefix1 = make([]int, n+1)
	prefix2 = make([]int, n+1)
	mul1 = make([]int, n+1)
	mul2 = make([]int, n+1)
	mul1[0], mul2[0] = 1, 1
	for i := 0; i < n; i++ {
		prefix1[i+1] = (prefix1[i]*base1 + int(arr[i]-'a')) % mod
		mul1[i+1] = (mul1[i] * base1) % mod
		prefix2[i+1] = (prefix2[i]*base2 + int(arr[i]-'a')) % mod
		mul2[i+1] = (mul2[i] * base2) % mod
	}
	ans := ""
	// 1 1 1 1 1 0 0 0
	p, q := 0, n-1
	for p < q {
		mid := (p + q + 1) >> 1
		if v := hasDupSubstringWithLength(arr, mid); v >= 0 {
			// fmt.Println(len(arr), mid, v)
			ans = string(arr[v : v+mid])
			p = mid
		} else {
			q = mid - 1
		}
	}
	return ans
}

func hasDupSubstringWithLength(arr []byte, k int) int {
	n := len(arr)
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for i := 0; i+k <= n; i++ {
		pass1, pass2 := false, false
		h := (prefix1[i+k] - prefix1[i]*mul1[k]%mod + mod) % mod
		if _, ok := m1[h]; ok {
			pass1 = true
		}
		m1[h] = struct{}{}
		h = (prefix2[i+k] - prefix2[i]*mul2[k]%mod + mod) % mod
		if _, ok := m2[h]; ok {
			pass2 = true
		}
		m2[h] = struct{}{}
		if pass1 && pass2 {
			return i
		}
	}
	return -1
}
