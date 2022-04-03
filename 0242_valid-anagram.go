/*
 * @lc app=leetcode.cn id=valid-anagram lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (64.96%)
 * Total Accepted:    374.8K
 * Total Submissions: 576K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length, t.length <= 5 * 1e4
 * s 和 t 仅包含小写字母
 *
 *
 *
 *
 * 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */

func isAnagram(s string, t string) bool {
	// assert islower(s) && islower(t);
	count := make([]int, 26)
	bs, bt := []byte(s), []byte(t)
	for _, b := range bs {
		count[int(b-'a')]++
	}
	for _, b := range bt {
		count[int(b-'a')]--
	}
	for b := 'a'; b <= 'z'; b++ {
		if count[int(b-'a')] != 0 {
			return false
		}
	}
	return true
}
