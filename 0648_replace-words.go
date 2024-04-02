/*
 * @lc app=leetcode.cn id=replace-words lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [648] 单词替换
 *
 * https://leetcode.cn/problems/replace-words/description/
 *
 * algorithms
 * Medium (60.09%)
 * Total Accepted:    36.5K
 * Total Submissions: 58.5K
 * Testcase Example:  '["cat","bat","rat"]\n"the cattle was rattled by the battery"'
 *
 * 在英语中，我们有一个叫做 词根(root) 的概念，可以词根后面添加其他一些词组成另一个较长的单词——我们称这个词为
 * 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
 *
 * 现在，给定一个由许多词根组成的词典 dictionary 和一个用空格分隔单词形成的句子
 * sentence。你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
 *
 * 你需要输出替换之后的句子。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by
 * the battery"
 * 输出："the cat was rat by the bat"
 *
 *
 * 示例 2：
 *
 *
 * 输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
 * 输出："a a b c"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= dictionary.length <= 1000
 * 1 <= dictionary[i].length <= 100
 * dictionary[i] 仅由小写字母组成。
 * 1 <= sentence.length <= 10^6
 * sentence 仅由小写字母和空格组成。
 * sentence 中单词的总量在范围 [1, 1000] 内。
 * sentence 中每个单词的长度在范围 [1, 1000] 内。
 * sentence 中单词之间由一个空格隔开。
 * sentence 没有前导或尾随空格。
 *
 *
 *
 *
 */

import "strings"

type Trie struct {
	Word     bool
	Children ['z' - 'a' + 1]*Trie
}

func (t *Trie) Add(word string) {
	cur := t
	for i, n := 0, len(word); i < n; i++ {
		idx := int(word[i] - 'a')
		if cur.Children[idx] == nil {
			cur.Children[idx] = NewTrie()
		}
		cur = cur.Children[idx]
	}
	cur.Word = true
	return
}

func (t *Trie) Get(word string) string {
	cur := t
	for i, n := 0, len(word); i < n; i++ {
		if cur.Word {
			return word[:i]
		}
		idx := int(word[i] - 'a')
		if cur.Children[idx] == nil {
			return word
		}
		cur = cur.Children[idx]
	}
	return word
}

func NewTrie() *Trie { return &Trie{} }

func replaceWords(dictionary []string, sentence string) string {
	t := NewTrie()
	for _, d := range dictionary {
		t.Add(d)
	}
	words := strings.Split(sentence, " ")
	var ans []string
	for _, w := range words {
		ans = append(ans, t.Get(w))
	}
	return strings.Join(ans, " ")
}
