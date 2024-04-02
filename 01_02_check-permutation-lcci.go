/*
 * Easy
 *
 * 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
 *
 * 示例 1：
 *
 * 输入: s1 = "abc", s2 = "bca"
 * 输出: true
 * 示例 2：
 *
 * 输入: s1 = "abc", s2 = "bad"
 * 输出: false
 * 说明：
 *
 * 0 <= len(s1) <= 100
 * 0 <= len(s2) <= 100
 * 通过次数92,410
 * 提交次数142,823
 */

func CheckPermutation(s1 string, s2 string) bool {
	var cnt ['z' - 'a' + 1]int
	for _, c := range s1 {
		cnt[int(c-'a')]++
	}
	for _, c := range s2 {
		cnt[int(c-'a')]--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}
