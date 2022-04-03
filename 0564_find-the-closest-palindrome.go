/*
 * @lc app=leetcode.cn id=find-the-closest-palindrome lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [564] 寻找最近的回文数
 *
 * https://leetcode-cn.com/problems/find-the-closest-palindrome/description/
 *
 * algorithms
 * Hard (18.50%)
 * Total Accepted:    5.4K
 * Total Submissions: 28K
 * Testcase Example:  '"123"'
 *
 * 给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。
 *
 * “最近的”定义为两个整数差的绝对值最小。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = "123"
 * 输出: "121"
 *
 *
 * 示例 2:
 *
 *
 * 输入: n = "1"
 * 输出: "0"
 * 解释: 0 和 2是最近的回文，但我们返回最小的，也就是 0。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= n.length <= 18
 * n 只由数字组成
 * n 不含前导 0
 * n 代表在 [1, 10^18 - 1] 范围内的整数
 *
 *
 */

import (
	"math"
	"strconv"
)

// 好恶心...
func nearestPalindromic(s string) string {
	bs := []byte(s)
	n := len(bs)
	candidates := []int{int(math.Pow10(n-1)) - 1, int(math.Pow10(n)) + 1}
	prefixNum, _ := strconv.Atoi(string(bs[:(n+1)/2]))
	for i := -1; i <= 1; i++ {
		v := prefixNum + i
		t := v
		if n%2 != 0 {
			t /= 10
		}
		for ; t > 0; t /= 10 {
			v = v*10 + t%10
		}
		candidates = append(candidates, v)
	}
	ans := -1
	self, _ := strconv.Atoi(s)
	for _, c := range candidates {
		if c == self {
			continue
		}
		if ans == -1 {
			ans = c
			continue
		}
		if abs(ans-self) > abs(c-self) {
			ans = c
			continue
		}
		if abs(ans-self) == abs(c-self) && c < ans {
			ans = c
			continue
		}
	}
	return strconv.Itoa(ans)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

