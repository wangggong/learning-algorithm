/*
 * @lc app=leetcode.cn id=group-anagrams lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (67.23%)
 * Total Accepted:    320.5K
 * Total Submissions: 476.7K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
 *
 * 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
 *
 * 示例 2:
 *
 *
 * 输入: strs = [""]
 * 输出: [[""]]
 *
 *
 * 示例 3:
 *
 *
 * 输入: strs = ["a"]
 * 输出: [["a"]]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] 仅包含小写字母
 *
 *
 */

const (
	alpha = 26
	base  = 41
	mod   = 1e9 + 7
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[int64][]string)
	for _, s := range strs {
		h := hash(s)
		m[h] = append(m[h], s)
	}
	var ans [][]string
	for _, cur := range m {
		ans = append(ans, cur)
	}
	return ans
}

func hash(s string) int64 {
	ans := int64(0)
	var count [alpha]int64
	for _, b := range []byte(s) {
		count[int(b-'a')]++
	}
	for _, c := range count {
		ans = (ans*base + c) % mod
	}
	return ans
}
