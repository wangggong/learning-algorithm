/*
 * @lc app=leetcode.cn id=word-break-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [140] 单词拆分 II
 *
 * https://leetcode-cn.com/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (51.02%)
 * Total Accepted:    60.4K
 * Total Submissions: 118.4K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序
 * 返回所有这些可能的句子。
 *
 * 注意：词典中的同一个单词可能在分段中被重复使用多次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
 * 输出:["cats and dog","cat sand dog"]
 *
 *
 * 示例 2：
 *
 *
 * 输入:s = "pineapplepenapple", wordDict =
 * ["apple","pen","applepen","pine","pineapple"]
 * 输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
 * 解释: 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 *
 * 输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
 * 输出:[]
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 20
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 10
 * s 和 wordDict[i] 仅有小写英文字母组成
 * wordDict 中所有字符串都 不同
 *
 *
 */
import "strings"

var ans []string
var cur []string
var wordMap map[string]struct{}

func dfs(bytes []byte, k, n int) {
	if k >= n {
		if k == n {
			ans = append(ans, strings.Join(cur, " "))
		}
		return
	}
	for i := 1; i <= 10 && k+i <= n; i++ {
		s := string(bytes[k : k+i])
		if _, ok := wordMap[s]; !ok {
			continue
		}
		cur = append(cur, s)
		dfs(bytes, k+i, n)
		cur = cur[:len(cur)-1]
	}
}

func wordBreak(s string, wordDict []string) []string {
	bytes := []byte(s)
	n := len(bytes)
	// assert n > 0 && n <= 20;
	// assert len(wordDict[i]) <= 10;
	ans, cur, wordMap = nil, nil, make(map[string]struct{})
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
	}
	dfs(bytes, 0, n)
	return ans
}
