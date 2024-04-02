/*
 * @lc app=leetcode.cn id=longest-binary-subsequence-less-than-or-equal-to-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6099] 小于等于 K 的最长二进制子序列
 *
 * https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/description/
 *
 * algorithms
 * Medium (27.22%)
 * Total Accepted:    2.5K
 * Total Submissions: 9.3K
 * Testcase Example:  '"1001010"\n5'
 *
 * 给你一个二进制字符串 s 和一个正整数 k 。
 *
 * 请你返回 s 的 最长 子序列，且该子序列对应的 二进制 数字小于等于 k 。
 *
 * 注意：
 *
 *
 * 子序列可以有 前导 0 。
 * 空字符串视为 0 。
 * 子序列 是指从一个字符串中删除零个或者多个字符后，不改变顺序得到的剩余字符序列。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "1001010", k = 5
 * 输出：5
 * 解释：s 中小于等于 5 的最长子序列是 "00010" ，对应的十进制数字是 2 。
 * 注意 "00100" 和 "00101" 也是可行的最长子序列，十进制分别对应 4 和 5 。
 * 最长子序列的长度为 5 ，所以返回 5 。
 *
 *
 * 示例 2：
 *
 * 输入：s = "00101001", k = 1
 * 输出：6
 * 解释："000001" 是 s 中小于等于 1 的最长子序列，对应的十进制数字是 1 。
 * 最长子序列的长度为 6 ，所以返回 6 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s[i] 要么是 '0' ，要么是 '1' 。
 * 1 <= k <= 10^9
 *
 *
 */

// 库函数拉满版
// 参考: https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/solution/fen-lei-tao-lun-tan-xin-by-endlesscheng-vnlx/
import (
	"math/bits"
	"strconv"
	"strings"
)

func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if n < m {
		return n
	}
	ans := m
	if v, _ := strconv.ParseInt(s[n-m:], 2, 64); int(v) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0")
}

/*
 * func longestSubsequence(s string, k int) int {
 * 	n, m := len(s), 0
 * 	for i := k; i > 0; i >>= 1 {
 * 		m++
 * 	}
 * 	// 总共还没人家长, 梭哈
 * 	if n < m {
 * 		return n
 * 	}
 * 	ans := m
 * 	// 统计 `s[n-m:]` 的数字, 如果比 k 大, 就最多只能取 `m-1` 位; 否则可以全取
 * 	v := 0
 * 	for i := n-m; i < n; i++ {
 * 		v <<= 1
 * 		if s[i] == '1' {
 * 			v |= 1
 * 		}
 * 	}
 * 	if v > k {
 * 		ans--
 * 	}
 * 	// 然后加上 `s[:n-m]` 的所有 `0` 作为前导 "0"
 * 	for i := 0; i < n-m; i++ {
 * 		if s[i] == '0' {
 * 			ans++
 * 		}
 * 	}
 * 	return ans
 * }
 */
