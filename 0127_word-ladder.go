/*
 * @lc app=leetcode.cn id=word-ladder lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Hard (47.38%)
 * Total Accepted:    143.7K
 * Total Submissions: 303.2K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 ->
 * s2 -> ... -> sk：
 *
 *
 * 每一对相邻的单词只差一个字母。
 * 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
 * sk == endWord
 *
 *
 * 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列
 * 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * 输出：5
 * 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
 *
 *
 * 示例 2：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * 输出：0
 * 解释：endWord "cog" 不在字典中，所以无法进行转换。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= beginWord.length <= 10
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 5000
 * wordList[i].length == beginWord.length
 * beginWord、endWord 和 wordList[i] 由小写英文字母组成
 * beginWord != endWord
 * wordList 中的所有字符串 互不相同
 *
 *
 */
type wordDist struct {
	word string
	dist int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	index := make(map[string]struct{})
	visit := make(map[string]struct{})
	for _, w := range wordList {
		index[w] = struct{}{}
	}
	var Q []wordDist
	Q = append(Q, wordDist{beginWord, 1})
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		if q.word == endWord {
			return q.dist
		}
		arr := []byte(q.word)
		m := len(arr)
		for i := 0; i < m; i++ {
			for j := 0; j < 26; j++ {
				arr[i] = byte(int(arr[i]-'a')+j)%26 + 'a'
				s := string(arr)
				arr[i] = byte(int(arr[i]-'a')-j+26)%26 + 'a'
				if _, ok := visit[s]; ok {
					continue
				}
				visit[s] = struct{}{}
				if _, ok := index[s]; !ok {
					continue
				}
				Q = append(Q, wordDist{s, q.dist + 1})
			}
		}
	}
	return 0
}
