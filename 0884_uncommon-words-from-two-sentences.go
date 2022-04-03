/*
 * @lc app=leetcode.cn id=uncommon-words-from-two-sentences lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [884] 两句话中的不常见单词
 *
 * https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/description/
 *
 * algorithms
 * Easy (66.26%)
 * Total Accepted:    40.5K
 * Total Submissions: 56.2K
 * Testcase Example:  '"this apple is sweet"\n"this apple is sour"'
 *
 * 句子 是一串由空格分隔的单词。每个 单词 仅由小写字母组成。
 *
 * 如果某个单词在其中一个句子中恰好出现一次，在另一个句子中却 没有出现 ，那么这个单词就是 不常见的 。
 *
 * 给你两个 句子 s1 和 s2 ，返回所有 不常用单词 的列表。返回列表中单词可以按 任意顺序 组织。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "this apple is sweet", s2 = "this apple is sour"
 * 输出：["sweet","sour"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1 = "apple apple", s2 = "banana"
 * 输出：["banana"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 200
 * s1 和 s2 由小写英文字母和空格组成
 * s1 和 s2 都不含前导或尾随空格
 * s1 和 s2 中的所有单词间均由单个空格分隔
 *
 *
 */
import "strings"

/*
func uncommonFromSentences(s1 string, s2 string) []string {
	var ans []string
	ws1 := strings.Split(s1, " ")
	ws2 := strings.Split(s2, " ")

	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for _, w := range ws1 {
		m1[w]++
	}
	for _, w := range ws2 {
		m2[w]++
	}
	for _, w := range ws1 {
		if m1[w] == 1 && m2[w] == 0 {
			ans = append(ans, w)
		}
	}
	for _, w := range ws2 {
		if m2[w] == 1 && m1[w] == 0 {
			ans = append(ans, w)
		}
	}
	return ans
}
*/

func uncommonFromSentences(s1 string, s2 string) []string {
	var ans []string
	s := s1 + " " + s2
	ws := strings.Split(s, " ")
	m := make(map[string]int)
	for _, w := range ws {
		m[w]++
	}
	for _, w := range ws {
		if m[w] == 1 {
			ans = append(ans, w)
		}
	}
	return ans
}
