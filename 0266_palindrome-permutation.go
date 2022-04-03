/*
 * Easy
 *
 * 给定一个字符串，判断该字符串中是否可以通过重新排列组合，形成一个回文字符串。
 *
 * 示例 1：
 *
 * 输入: "code"
 * 输出: false
 * 示例 2：
 *
 * 输入: "aab"
 * 输出: true
 * 示例 3：
 *
 * 输入: "carerac"
 * 输出: true
 * 通过次数8,033
 * 提交次数11,700
 */

var counter [256]int

func canPermutePalindrome(s string) bool {
	bytes := []byte(s)
	for i := range counter {
		counter[i] = 0
	}
	for _, b := range bytes {
		counter[int(b)]++
	}
	single := false
	for _, c := range counter {
		if c%2 == 0 {
			continue
		}
		if single {
			return false
		}
		single = true
	}
	return true
}
