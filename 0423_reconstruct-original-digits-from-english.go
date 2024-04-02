/*
 * @lc app=leetcode.cn id=reconstruct-original-digits-from-english lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [423] 从英文中重建数字
 *
 * https://leetcode.cn/problems/reconstruct-original-digits-from-english/description/
 *
 * algorithms
 * Medium (61.06%)
 * Total Accepted:    35.9K
 * Total Submissions: 58.8K
 * Testcase Example:  '"owoztneoer"'
 *
 * 给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "owoztneoer"
 * 输出："012"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "fviefuro"
 * 输出："45"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"]
 * 这些字符之一
 * s 保证是一个符合题目要求的字符串
 *
 *
 */
var nums = []string{
	"zero",  // z
	"one",   // n-7-9-9
	"two",   // w
	"three", // h-8
	"four",  // u
	"five",  // v-7
	"six",   // x
	"seven", // s-6
	"eight", // g
	"nine",  // i-5-6-8
}

func originalDigits(s string) string {
	var cnt [10]int
	var cntByte [26]int
	for _, b := range s {
		cntByte[int(b-'a')]++
	}
	cnt[0] = cntByte[int('z'-'a')]
	cnt[2] = cntByte[int('w'-'a')]
	cnt[4] = cntByte[int('u'-'a')]
	cnt[6] = cntByte[int('x'-'a')]
	cnt[8] = cntByte[int('g'-'a')]
	cnt[3] = cntByte[int('h'-'a')] - cnt[8]
	cnt[7] = cntByte[int('s'-'a')] - cnt[6]
	cnt[5] = cntByte[int('v'-'a')] - cnt[7]
	cnt[9] = cntByte[int('i'-'a')] - cnt[5] - cnt[6] - cnt[8]
	cnt[1] = cntByte[int('n'-'a')] - cnt[7] - cnt[9]*2
	var ans []byte
	for i, c := range cnt {
		for j := 0; j < c; j++ {
			ans = append(ans, byte(i)+'0')
		}
	}
	return string(ans)
}
