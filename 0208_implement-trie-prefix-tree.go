/*
 * @lc app=leetcode.cn id=implement-trie-prefix-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (71.88%)
 * Total Accepted:    166.6K
 * Total Submissions: 231.8K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * Trie（发音类似 "try"）或者说 前缀树
 * 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
 *
 * 请你实现 Trie 类：
 *
 *
 * Trie() 初始化前缀树对象。
 * void insert(String word) 向前缀树中插入字符串 word 。
 * boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回
 * false 。
 * boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true
 * ；否则，返回 false 。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
 * [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
 * 输出
 * [null, null, true, false, true, null, true]
 *
 * 解释
 * Trie trie = new Trie();
 * trie.insert("apple");
 * trie.search("apple");   // 返回 True
 * trie.search("app");     // 返回 False
 * trie.startsWith("app"); // 返回 True
 * trie.insert("app");
 * trie.search("app");     // 返回 True
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * word 和 prefix 仅由小写英文字母组成
 * insert、search 和 startsWith 调用次数 总计 不超过 3 * 10^4 次
 *
 *
*/
const ALPHA_SIZE = 26

type Trie struct {
	count    int
	children [ALPHA_SIZE]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(S string) {
	B := []byte(S)
	cur := t
	for _, b := range B {
		i := b - 'a'
		if cur.children[i] == nil {
			c := Constructor()
			cur.children[i] = &c
		}
		cur = cur.children[i]
	}
	cur.count++
}

func (t *Trie) searchPrefix(S string) *Trie {
	B := []byte(S)
	cur := t
	for _, b := range B {
		i := b - 'a'
		if cur.children[i] == nil {
			return nil
		}
		cur = cur.children[i]
	}
	return cur
}

func (t *Trie) Search(S string) bool {
	c := t.searchPrefix(S)
	return c != nil && c.count > 0
}

func (t *Trie) StartsWith(S string) bool {
	return t.searchPrefix(S) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
