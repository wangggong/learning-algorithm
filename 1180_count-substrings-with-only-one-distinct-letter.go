/*
 * Easy
 *
 * 给你一个字符串 s，返回 只含 单一字母 的子串个数 。
 *
 * 示例 1：
 *
 * 输入： s = "aaaba"
 * 输出： 8
 * 解释： 只含单一字母的子串分别是 "aaa"， "aa"， "a"， "b"。
 * "aaa" 出现 1 次。
 * "aa" 出现 2 次。
 * "a" 出现 4 次。
 * "b" 出现 1 次。
 * 所以答案是 1 + 2 + 4 + 1 = 8。
 * 示例 2:
 *
 * 输入： s = "aaaaaaaaaa"
 * 输出： 55
 *
 *
 * 提示：
 *
 * 1 <= s.length <= 1000
 * s[i] 仅由小写英文字母组成
 * 通过次数3,591
 * 提交次数4,626
 */

func countLetters(s string) int {
	arr := []byte(s)
	n := len(arr)
	ans := 0
	for p, q := 0, 0; q < n; p = q {
		for ; q < n && arr[q] == arr[p]; q++ {
		}
		L := q - p
		ans += L * (L + 1) >> 1
	}
	return ans
}
