/*
 * @lc app=leetcode.cn id=stream-of-characters lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1032] 字符流
 *
 * https://leetcode-cn.com/problems/stream-of-characters/description/
 *
 * algorithms
 * Hard (42.36%)
 * Total Accepted:    3.3K
 * Total Submissions: 7.8K
 * Testcase Example:  '["StreamChecker","query","query","query","query","query","query","query","query","query","query","query","query"]\n' +
  '[[["cd","f","kl"]],["a"],["b"],["c"],["d"],["e"],["f"],["g"],["h"],["i"],["j"],["k"],["l"]]'
 *
 * 设计一个算法：接收一个字符流，并检查这些字符的后缀是否是字符串数组 words 中的一个字符串。
 *
 * 例如，words = ["abc", "xyz"] 且字符流中逐个依次加入 4 个字符 'a'、'x'、'y' 和 'z'
 * ，你所设计的算法应当可以检测到 "axyz" 的后缀 "xyz" 与 words 中的字符串 "xyz" 匹配。
 *
 * 按下述要求实现 StreamChecker 类：
 *
 *
 * StreamChecker(String[] words) ：构造函数，用字符串数组 words 初始化数据结构。
 * boolean query(char letter)：从字符流中接收一个新字符，如果字符流中的任一非空后缀能匹配 words 中的某一字符串，返回
 * true ；否则，返回 false。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["StreamChecker", "query", "query", "query", "query", "query", "query",
 * "query", "query", "query", "query", "query", "query"]
 * [[["cd", "f", "kl"]], ["a"], ["b"], ["c"], ["d"], ["e"], ["f"], ["g"],
 * ["h"], ["i"], ["j"], ["k"], ["l"]]
 * 输出：
 * [null, false, false, false, true, false, true, false, false, false, false,
 * false, true]
 *
 * 解释：
 * StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
 * streamChecker.query("a"); // 返回 False
 * streamChecker.query("b"); // 返回 False
 * streamChecker.query("c"); // 返回n False
 * streamChecker.query("d"); // 返回 True ，因为 'cd' 在 words 中
 * streamChecker.query("e"); // 返回 False
 * streamChecker.query("f"); // 返回 True ，因为 'f' 在 words 中
 * streamChecker.query("g"); // 返回 False
 * streamChecker.query("h"); // 返回 False
 * streamChecker.query("i"); // 返回 False
 * streamChecker.query("j"); // 返回 False
 * streamChecker.query("k"); // 返回 False
 * streamChecker.query("l"); // 返回 True ，因为 'kl' 在 words 中
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= words.length <= 2000
 * 1 <= words[i].length <= 2000
 * words[i] 由小写英文字母组成
 * letter 是一个小写英文字母
 * 最多调用查询 4 * 10^4 次
 *
 *
*/

/*
 * // 先来个模拟面试时的超时版本:
 * const mod1 int = 1e9 + 7
 * const base1 int = 47
 * const mod2 int = 15485863
 * const base2 int = 37
 *
 * type StreamChecker struct {
 * 	Hash1     map[int]string
 * 	Hash2     map[int]string
 * 	QuerySize int
 * 	Prefix1   []int
 * 	Prefix2   []int
 * 	Mul       []int
 * 	CheckSize int
 * }
 *
 * func (sc *StreamChecker) hashString(s string) {
 * 	arr := []byte(s)
 * 	n := len(arr)
 * 	h := 0
 * 	for i := 0; i < n; i++ {
 * 		h = (h*base1 + int(arr[i]-'a'+1)) % mod1
 * 	}
 * 	sc.Hash1[h] = s
 * 	h = 0
 * 	for i := 0; i < n; i++ {
 * 		h = (h*base2 + int(arr[i]-'a'+1)) % mod2
 * 	}
 * 	sc.Hash2[h] = s
 * }
 *
 * func Constructor(words []string) StreamChecker {
 * 	sc := StreamChecker{
 * 		Hash1: make(map[int]string),
 * 		Hash2: make(map[int]string),
 * 	}
 * 	for _, w := range words {
 * 		sc.hashString(w)
 * 		if len(w) > sc.CheckSize {
 * 			sc.CheckSize = len(w)
 * 		}
 * 	}
 * 	sc.Prefix1 = make([]int, sc.CheckSize)
 * 	sc.Prefix2 = make([]int, sc.CheckSize)
 * 	return sc
 * }
 *
 * func (sc *StreamChecker) Query(letter byte) bool {
 * 	sc.QuerySize++
 * 	has := false
 * 	if sc.QuerySize <= sc.CheckSize {
 * 		for i := 0; i < sc.QuerySize; i++ {
 * 			sc.Prefix1[i] = (sc.Prefix1[i]*base1 + int(letter-'a'+1)) % mod1
 * 			sc.Prefix2[i] = (sc.Prefix2[i]*base2 + int(letter-'a'+1)) % mod2
 * 			_, ok1 := sc.Hash1[sc.Prefix1[i]]
 * 			_, ok2 := sc.Hash2[sc.Prefix2[i]]
 * 			if ok1 && ok2 {
 * 				has = true
 * 			}
 * 		}
 * 		return has
 * 	}
 * 	sc.Prefix1 = sc.Prefix1[1:]
 * 	sc.Prefix2 = sc.Prefix2[1:]
 * 	sc.Prefix1 = append(sc.Prefix1, 0)
 * 	sc.Prefix2 = append(sc.Prefix2, 0)
 * 	for i := 0; i < sc.CheckSize; i++ {
 * 		sc.Prefix1[i] = (sc.Prefix1[i]*base1 + int(letter-'a'+1)) % mod1
 * 		sc.Prefix2[i] = (sc.Prefix2[i]*base2 + int(letter-'a'+1)) % mod2
 * 		_, ok1 := sc.Hash1[sc.Prefix1[i]]
 * 		_, ok2 := sc.Hash2[sc.Prefix2[i]]
 * 		if ok1 && ok2 {
 * 			has = true
 * 		}
 * 	}
 * 	return has
 * }
 *
 */

// 官解是字典树, 这也就算了, 还是特么倒着建的...
const alpha int = 26

type Trie struct {
	Word     bool
	Children [alpha]*Trie
}

func (t *Trie) insert(arr []byte) {
	if len(arr) == 0 {
		t.Word = true
		return
	}
	ind := int(arr[0] - 'a')
	if t.Children[ind] == nil {
		t.Children[ind] = &Trie{}
	}
	t.Children[ind].insert(arr[1:])
}

func (t *Trie) query(arr []byte) bool {
	if t.Word {
		return true
	}
	if len(arr) == 0 {
		return false
	}
	ind := int(arr[0] - 'a')
	if t.Children[ind] == nil {
		return false
	}
	return t.Children[ind].query(arr[1:])
}

type StreamChecker struct {
	*Trie
	MaxSize int
	Buffer  []byte
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{
		Trie: &Trie{},
	}
	for _, w := range words {
		arr := []byte(w)
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		sc.insert(arr)
		if sc.MaxSize < len(arr) {
			sc.MaxSize = len(arr)
		}
	}
	return sc
}

func (sc *StreamChecker) Query(letter byte) bool {
	sc.Buffer = append([]byte{letter}, sc.Buffer...)
	if len(sc.Buffer) > sc.MaxSize {
		sc.Buffer = sc.Buffer[:sc.MaxSize]
	}
	return sc.query(sc.Buffer)
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
