/*
 * @lc app=leetcode.cn id=goat-latin lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [824] 山羊拉丁文
 *
 * https://leetcode-cn.com/problems/goat-latin/description/
 *
 * algorithms
 * Easy (61.97%)
 * Total Accepted:    26.1K
 * Total Submissions: 40.7K
 * Testcase Example:  '"I speak Goat Latin"'
 *
 * 给你一个由若干单词组成的句子 sentence ，单词间由空格分隔。每个单词仅由大写和小写英文字母组成。
 *
 * 请你将句子转换为 “山羊拉丁文（Goat Latin）”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。山羊拉丁文的规则如下：
 *
 *
 * 如果单词以元音开头（'a', 'e', 'i', 'o', 'u'），在单词后添加"ma"。
 *
 *
 * 例如，单词 "apple" 变为 "applema" 。
 *
 *
 * 如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
 *
 * 例如，单词 "goat" 变为 "oatgma" 。
 *
 *
 * 根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从 1 开始。
 *
 * 例如，在第一个单词后添加 "a" ，在第二个单词后添加 "aa" ，以此类推。
 *
 *
 *
 *
 * 返回将 sentence 转换为山羊拉丁文后的句子。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：sentence = "I speak Goat Latin"
 * 输出："Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
 *
 *
 * 示例 2：
 *
 *
 * 输入：sentence = "The quick brown fox jumped over the lazy dog"
 * 输出："heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa
 * hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= sentence.length <= 150
 * sentence 由英文字母和空格组成
 * sentence 不含前导或尾随空格
 * sentence 中的所有单词由单个空格分隔
 *
 *
 */

const (
	ma     = "ma"
	vowels = "aeiouAEIOU"
)

func toGoatLatin(sentence string) string {
	words := Split(sentence, " ")
	var ans []string
	for i, w := range words {
		goat := goatLatin(w) + Repeat("a", i+1)
		ans = append(ans, goat)
	}
	return Join(ans, " ")
}

func goatLatin(w string) string {
	if len(w) == 0 {
		return w
	}
	if isVowel(w[0]) {
		return w + ma
	}
	return w[1:] + w[0:1] + ma
}

func isVowel(b byte) bool {
	for _, v := range vowels {
		if byte(v) == b {
			return true
		}
	}
	return false
}

func Split(s, split string) []string {
	n, m := len(s), len(split)
	var ans []string
	for i, j := 0, 0; i < n; {
		for j = i; j < n; j++ {
			if s[j:j+m] == split {
				break
			}
		}
		ans = append(ans, s[i:j])
		i = j + m
	}
	return ans
}

func Repeat(s string, n int) string {
	var ans string
	for ; n > 0; n-- {
		ans = ans + s
	}
	return ans
}

func Join(s []string, split string) string {
	var ans string
	n := len(s)
	for i := 0; i < n; i++ {
		ans = ans + s[i]
		if i+1 < n {
			ans = ans + split
		}
	}
	return ans
}