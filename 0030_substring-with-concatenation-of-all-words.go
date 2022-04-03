/*
 * @lc app=leetcode.cn id=substring-with-concatenation-of-all-words lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (36.53%)
 * Total Accepted:    94.8K
 * Total Submissions: 259.2K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 *
 * 注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "barfoothefoobarman", words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：
 * 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 * 输出的顺序不重要, [9,0] 也是有效答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 * 输出：[6,9,12]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1e4
 * s 由小写英文字母组成
 * 1 <= words.length <= 5000
 * 1 <= words[i].length <= 30
 * words[i] 由小写英文字母组成
 *
 *
 */

// NOTE: 脑抽了, 无脑考虑字符串匹配... 直接 HASH 即可.
// assert len(words) > 0;
// 参考: https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/solution/hardti-jian-dan-jie-xiao-bai-miao-dong-j-e9zr/
func findSubstring(s string, words []string) []int {
	bs := []byte(s)
	n := len(bs)
	m := len(words)
	l := len(words[0])
	// 查询 words 词频
	count := make(map[string]int)
	for _, w := range words {
		count[w]++
	}
	var ans []int
	// 枚举每个起点...
	for i := 0; i+m*l <= n; i++ {
		// 统计当前词频...
		c := make(map[string]int)
		j := 0
		for ; j < m; j++ {
			// 当前第 j 个词...
			cs := string(bs[i+j*l : i+(j+1)*l])
			cnt, _ := count[cs]
			/*
			 * // ... 如果不在词频表里面, 说明不符合条件 (仔细想想就是词频为 0 嘛)
			 * if !ok {
			 * 	break
			 * }
			 */
			c[cs]++
			// ... 如果词频爆了, 说明也不符合条件;
			if c[cs] > cnt {
				break
			}
		}
		// 如果中道崩殂了, 枚举下一个...
		if j < m {
			continue
		}
		ans = append(ans, i)
	}
	return ans
}
