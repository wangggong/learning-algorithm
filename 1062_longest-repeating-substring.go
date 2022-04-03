/*
 * Medium
 * 通过次数3,989
 * 提交次数7,228
 *
 * 给定字符串 S，找出最长重复子串的长度。如果不存在重复子串就返回 0。
 *
 *
 *
 * 示例 1：
 *
 * 输入："abcd"
 * 输出：0
 * 解释：没有重复子串。
 * 示例 2：
 *
 * 输入："abbaba"
 * 输出：2
 * 解释：最长的重复子串为 "ab" 和 "ba"，每个出现 2 次。
 * 示例 3：
 *
 * 输入："aabcaabdaab"
 * 输出：3
 * 解释：最长的重复子串为 "aab"，出现 3 次。
 * 示例 4：
 *
 * 输入："aaaaa"
 * 输出：4
 * 解释：最长的重复子串为 "aaaa"，出现 2 次。
 *
 *
 * 提示：
 *
 * 字符串 S 仅包含从 'a' 到 'z' 的小写英文字母。
 * 1 <= S.length <= 1500
 *
 */
var n int

func search(bytes []byte, start int) bool {
	m := make(map[string]struct{})
	for i := start + 1; i < n; i++ {
		s := string(bytes[start:i])
		if _, ok := m[s]; ok {
			return true
		}
		m[s] = struct{}{}
	}
	return false
}

// 1 1 1 1 0 0 0
// p = -1, q = n-1
// 退出条件: p < q
// if match: p = mid
// else: q = mid-1
func longestRepeatingSubstring(s string) int {
	bytes := []byte(s)
	n = len(bytes)
	// assert n > 0 && n < 1500;
	p, q := -1, n-1
	for p < q {
		mid := (p + q + 1) >> 1
		switch v := search(bytes, mid); {
		case v:
			p = mid
		default:
			q = mid - 1
		}
	}
	return p
}
