/*
 * Hard
 *
 * 现有一种使用英语字母的火星语言，这门语言的字母顺序与英语顺序不同。
 *
 * 给你一个字符串列表 words ，作为这门语言的词典，words 中的字符串已经 按这门新语言的字母顺序进行了排序 。
 *
 * 请你根据该词典还原出此语言中已知的字母顺序，并 按字母递增顺序 排列。若不存在合法字母顺序，返回 "" 。若存在多种可能的合法字母顺序，返回其中 任意一种 顺序即可。
 *
 * 字符串 s 字典顺序小于 字符串 t 有两种情况：
 *
 * 在第一个不同字母处，如果 s 中的字母在这门外星语言的字母顺序中位于 t 中字母之前，那么 s 的字典顺序小于 t 。
 * 如果前面 min(s.length, t.length) 字母都相同，那么 s.length < t.length 时，s 的字典顺序也小于 t 。
 *
 *
 * 示例 1：
 *
 * 输入：words = ["wrt","wrf","er","ett","rftt"]
 * 输出："wertf"
 * 示例 2：
 *
 * 输入：words = ["z","x"]
 * 输出："zx"
 * 示例 3：
 *
 * 输入：words = ["z","x","z"]
 * 输出：""
 * 解释：不存在合法字母顺序，因此返回 "" 。
 *
 *
 * 提示：
 *
 * 1 <= words.length <= 100
 * 1 <= words[i].length <= 100
 * words[i] 仅由小写英文字母组成
 * 通过次数8,697
 * 提交次数24,773
 */

const alpha int = 26

func alienOrder(words []string) string {
	n := len(words)
	var alphabet []byte
	var index [alpha]int
	for i := range index {
		index[i] = -1
	}
	for _, w := range words {
		for _, b := range []byte(w) {
			if index[int(b-'a')] == -1 {
				index[int(b-'a')] = len(alphabet)
				alphabet = append(alphabet, b)
			}
		}
	}
	G := make([][]int, len(alphabet))
	indegree := make([]int, len(alphabet))
	for i := 0; i+1 < n; i++ {
		if len(words[i]) > len(words[i+1]) && strings.HasPrefix(words[i], words[i+1]) {
			return ""
		}
		less, more := []byte(words[i]), []byte(words[i+1])
		m := min(len(less), len(more))
		for j := 0; j < m; j++ {
			if less[j] != more[j] {
				p, q := index[int(less[j]-'a')], index[int(more[j]-'a')]
				G[p] = append(G[p], q)
				indegree[q]++
				break
			}
		}
	}
	var ans []byte
	var Q []int
	for i, id := range indegree {
		if id == 0 {
			Q = append(Q, i)
			ans = append(ans, alphabet[i])
		}
	}
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		for _, k := range G[q] {
			indegree[k]--
			if indegree[k] == 0 {
				Q = append(Q, k)
				ans = append(ans, alphabet[k])
			}
		}
	}
	if len(ans) < len(alphabet) {
		return ""
	}
	return string(ans)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
