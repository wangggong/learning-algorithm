/*
 * @lc app=leetcode.cn id=prefix-and-suffix-search lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [745] 前缀和后缀搜索
 *
 * https://leetcode.cn/problems/prefix-and-suffix-search/description/
 *
 * algorithms
 * Hard (41.28%)
 * Total Accepted:    12.6K
 * Total Submissions: 28.9K
 * Testcase Example:  '["WordFilter","f"]\n[[["apple"]],["a","e"]]'
 *
 * 设计一个包含一些单词的特殊词典，并能够通过前缀和后缀来检索单词。
 *
 * 实现 WordFilter 类：
 *
 *
 * WordFilter(string[] words) 使用词典中的单词 words 初始化对象。
 * f(string pref, string suff) 返回词典中具有前缀 prefix 和后缀 suff
 * 的单词的下标。如果存在不止一个满足要求的下标，返回其中 最大的下标 。如果不存在这样的单词，返回 -1 。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["WordFilter", "f"]
 * [[["apple"]], ["a", "e"]]
 * 输出
 * [null, 0]
 * 解释
 * WordFilter wordFilter = new WordFilter(["apple"]);
 * wordFilter.f("a", "e"); // 返回 0 ，因为下标为 0 的单词：前缀 prefix = "a" 且 后缀 suff = "e"
 * 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= words.length <= 10^4
 * 1 <= words[i].length <= 7
 * 1 <= pref.length, suff.length <= 7
 * words[i]、pref 和 suff 仅由小写英文字母组成
 * 最多对函数 f 执行 10^4 次调用
 *
 *
 */

/*
 * // 字典树
 * type Trie struct {
 * 	Index    int
 * 	Children ['{' - 'a' + 1]*Trie
 * }
 *
 * func (t *Trie) Insert(s string, k int) {
 * 	t.Index = k
 * 	if len(s) == 0 {
 * 		return
 * 	}
 * 	ind := int(s[0] - 'a')
 * 	if t.Children[ind] == nil {
 * 		t.Children[ind] = NewTrie()
 * 	}
 * 	t.Children[ind].Insert(s[1:], k)
 * 	return
 * }
 *
 * func (t *Trie) Query(s string) int {
 * 	if len(s) == 0 {
 * 		return t.Index
 * 	}
 * 	ind := int(s[0] - 'a')
 * 	if t.Children[ind] == nil {
 * 		return -1
 * 	}
 * 	return t.Children[ind].Query(s[1:])
 * }
 *
 * func NewTrie() *Trie {
 * 	return &Trie{Index: -1}
 * }
 *
 * type WordFilter struct {
 * 	*Trie
 * }
 *
 * func Constructor(words []string) WordFilter {
 * 	t := NewTrie()
 * 	for i, w := range words {
 * 		for j, n := 0, len(w); j < n; j++ {
 * 			s := w[j:] + "{" + w
 * 			t.Insert(s, i)
 * 		}
 * 	}
 * 	return WordFilter{Trie: t}
 * }
 *
 * func (wf *WordFilter) F(pref string, suff string) int {
 * 	return wf.Query(suff + "{" + pref)
 * }
 */

// 哈希瞎搞
type WordFilter struct {
	Indexes map[string]int
}

func Constructor(words []string) WordFilter {
	m := make(map[string]int)
	for i, w := range words {
		for j, n := 0, len(w); j < n; j++ {
			for k := 0; k < n; k++ {
				s := w[k:] + "{" + w[:j+1]
				m[s] = i
			}
		}
	}
	return WordFilter{Indexes: m}
}

func (wf *WordFilter) F(pref string, suff string) int {
	if i, ok := wf.Indexes[suff+"{"+pref]; ok {
		return i
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
