/*
 * @lc app=leetcode.cn id=additive-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [306] 累加数
 *
 * https://leetcode-cn.com/problems/additive-number/description/
 *
 * algorithms
 * Medium (38.25%)
 * Total Accepted:    39.2K
 * Total Submissions: 102.5K
 * Testcase Example:  '"112358"'
 *
 * 累加数 是一个字符串，组成它的数字可以形成累加序列。
 *
 * 一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，序列中的每个后续数字必须是它之前两个数字之和。
 *
 * 给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。
 *
 * 说明：累加序列里的数，除数字 0 之外，不会 以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入："112358"
 * 输出：true
 * 解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
 *
 *
 * 示例 2：
 *
 *
 * 输入："199100199"
 * 输出：true
 * 解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num.length <= 35
 * num 仅由数字（0 - 9）组成
 *
 *
 *
 *
 * 进阶：你计划如何处理由过大的整数输入导致的溢出?
 *
 */
var n int

func atoi(B []byte) int {
	ans := 0
	for _, b := range B {
		ans *= 10
		ans += int(b - '0')
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func add(x, y []byte) []byte {
	var ans []byte
	n, m := len(x), len(y)
	p, q := n-1, m-1
	c := byte(0)
	for p >= 0 || q >= 0 {
		s, t := byte(0), byte(0)
		if p >= 0 {
			s = x[p] - '0'
		}
		if q >= 0 {
			t = y[q] - '0'
		}
		v := s + t + c
		ans = append(ans, (v%10)+'0')
		c = v / 10
		p, q = p-1, q-1
	}
	if c != 0 {
		ans = append(ans, (c%10)+'0')
	}
	for i, j := 0, len(ans)-1; i <= j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func check(x, y []byte) bool {
	m, n := len(x), len(y)
	if m != n {
		return false
	}
	for i := 0; i < n; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func dfs(B []byte, s, p, q int) bool {
	if q == n {
		return true
	}
	for k := q + 1; k <= n; k++ {
		// fmt.Printf("~~ %s %v %v %v %v %s %s %s %s\n", B, s, p, q, k, B[s:p], B[p:q], B[q:k], add(B[s:p], B[p:q]))
		if string(add(B[s:p], B[p:q])) != string(B[q:k]) {
			continue
		}
		v := dfs(B, p, q, k)
		// fmt.Printf("-> %s %v %v %v %s %s %s %v\n", B, p, q, k, B[p:q], B[q:k], B[k:], v)
		if v {
			return true
		}
	}
	return false
}

func isAdditiveNumber(N string) bool {
	B := []byte(N)
	n = len(B)
	for p := 1; p+1 < n; p++ {
		if B[0] == '0' && p != 1 {
			continue
		}
		for q := p + 1; q < n; q++ {
			if B[p] == '0' && q != p+1 {
				continue
			}
			if dfs(B, 0, p, q) {
				return true
			}
		}
	}
	return false
}
