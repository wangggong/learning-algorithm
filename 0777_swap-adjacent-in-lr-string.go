/*
 * @lc app=leetcode.cn id=swap-adjacent-in-lr-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [777] 在LR字符串中交换相邻字符
 *
 * https://leetcode.cn/problems/swap-adjacent-in-lr-string/description/
 *
 * algorithms
 * Medium (33.04%)
 * Total Accepted:    18.5K
 * Total Submissions: 50.7K
 * Testcase Example:  '"RXXLRXRXL"\n"XRLXXRRLX"'
 *
 * 在一个由 'L' , 'R' 和 'X'
 * 三个字符组成的字符串（例如"RXXLRXRXL"）中进行移动操作。一次移动操作指用一个"LX"替换一个"XL"，或者用一个"XR"替换一个"RX"。现给定起始字符串start和结束字符串end，请编写代码，当且仅当存在一系列移动操作使得start可以转换成end时，
 * 返回True。
 *
 *
 *
 * 示例 :
 *
 * 输入: start = "RXXLRXRXL", end = "XRLXXRRLX"
 * 输出: True
 * 解释:
 * 我们可以通过以下几步将start转换成end:
 * RXXLRXRXL ->
 * XRXLRXRXL ->
 * XRLXRXRXL ->
 * XRLXXRRXL ->
 * XRLXXRRLX
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= len(start) = len(end) <= 10000。
 * start和end中的字符串仅限于'L', 'R'和'X'。
 *
 *
 */
func canTransform(s, e string) bool {
	for p, q, n := 0, 0, len(s); p < n || q < n; p, q = p+1, q+1 {
		for ; p < n && s[p] == 'X'; p++ {
		}
		for ; q < n && e[q] == 'X'; q++ {
		}
		if p == n || q == n {
			return p == n && q == n
		}
		if s[p] != e[q] {
			return false
		}
		if s[p] == 'L' && p < q {
			return false
		}
		if s[p] == 'R' && p > q {
			return false
		}
	}
	return true
}

/*
 * // 自己写的愚蠢的大模拟
 * func canTransform(s, e string) bool {
 * 	var irs, ire, ils, ile []int
 * 	bs := []byte(s)
 * 	for i, b := range s {
 * 		if b == 'R' {
 * 			irs = append(irs, i)
 * 		}
 * 	}
 * 	for i, b := range e {
 * 		if b == 'R' {
 * 			ire = append(ire, i)
 * 		}
 * 	}
 * 	if len(irs) != len(ire) {
 * 		return false
 * 	}
 * 	for i := len(irs) - 1; i >= 0; i-- {
 * 		if irs[i] > ire[i] {
 * 			return false
 * 		}
 * 		for j := irs[i] + 1; j <= ire[i]; j++ {
 * 			if bs[j] != 'X' {
 * 				return false
 * 			}
 * 		}
 * 		bs[irs[i]], bs[ire[i]] = bs[ire[i]], bs[irs[i]]
 * 	}
 * 	s = string(bs)
 * 	for i, b := range s {
 * 		if b == 'L' {
 * 			ils = append(ils, i)
 * 		}
 * 	}
 * 	for i, b := range e {
 * 		if b == 'L' {
 * 			ile = append(ile, i)
 * 		}
 * 	}
 * 	if len(ils) != len(ile) {
 * 		return false
 * 	}
 * 	for i, m := 0, len(ils); i < m; i++ {
 * 		if ils[i] < ile[i] {
 * 			return false
 * 		}
 * 		for j := ils[i] - 1; j >= ile[i]; j-- {
 * 			if bs[j] != 'X' {
 * 				return false
 * 			}
 * 		}
 * 		bs[ils[i]], bs[ile[i]] = bs[ile[i]], bs[ils[i]]
 * 	}
 * 	return string(bs) == e
 * }
 */
