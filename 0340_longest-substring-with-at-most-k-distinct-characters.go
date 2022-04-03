/*
 * Medium
 *
 * 给定一个字符串 s ，找出 至多 包含 k 个不同字符的最长子串 T。
 *
 *
 *
 * 示例 1:
 *
 * 输入: s = "eceba", k = 2
 * 输出: 3
 * 解释: 则 T 为 "ece"，所以长度为 3。
 * 示例 2:
 *
 * 输入: s = "aa", k = 1
 * 输出: 2
 * 解释: 则 T 为 "aa"，所以长度为 2。
 *
 *
 * 提示：
 *
 * 1 <= s.length <= 5 * 104
 * 0 <= k <= 50
 * 通过次数15,584
 * 提交次数31,072
 */

const MAXN = 1 << 8

var count [MAXN]int

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	bs := []byte(s)
	n := len(bs)
	for i := range count {
		count[i] = 0
	}
	p, q := 0, 0
	ans := 0
	for ; q < n; q++ {
		idx := int(bs[q])
		count[idx]++
		// NOTE: 这里需要先把最长子数组算出来再更新状态.
		if getCharCount() <= k && ans < q-p+1 {
			ans = q - p + 1
		}
		for ; p < q && getCharCount() > k; p++ {
			idx := int(bs[p])
			count[idx]--
		}
	}
	return ans
}

func getCharCount() int {
	ans := 0
	for _, c := range count {
		if c > 0 {
			ans++
		}
	}
	return ans
}
