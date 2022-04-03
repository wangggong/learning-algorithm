/*
 * @lc app=leetcode.cn id=scramble-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [87] 扰乱字符串
 *
 * https://leetcode-cn.com/problems/scramble-string/description/
 *
 * algorithms
 * Hard (48.43%)
 * Total Accepted:    44.3K
 * Total Submissions: 91.4K
 * Testcase Example:  '"great"\n"rgeat"'
 *
 * 使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
 *
 * 如果字符串的长度为 1 ，算法停止
 * 如果字符串的长度 > 1 ，执行下述步骤：
 *
 * 在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y
 * 。
 * 随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y +
 * x 。
 * 在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
 *
 *
 *
 *
 * 给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "great", s2 = "rgeat"
 * 输出：true
 * 解释：s1 上可能发生的一种情形是：
 * "great" --> "gr/eat" // 在一个随机下标处分割得到两个子字符串
 * "gr/eat" --> "gr/eat" // 随机决定：「保持这两个子字符串的顺序不变」
 * "gr/eat" --> "g/r / e/at" // 在子字符串上递归执行此算法。两个子字符串分别在随机下标处进行一轮分割
 * "g/r / e/at" --> "r/g / e/at" // 随机决定：第一组「交换两个子字符串」，第二组「保持这两个子字符串的顺序不变」
 * "r/g / e/at" --> "r/g / e/ a/t" // 继续递归执行此算法，将 "at" 分割得到 "a/t"
 * "r/g / e/ a/t" --> "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
 * 算法终止，结果字符串和 s2 相同，都是 "rgeat"
 * 这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1 = "abcde", s2 = "caebd"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：s1 = "a", s2 = "a"
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * s1.length == s2.length
 * 1 <= s1.length <= 30
 * s1 和 s2 由小写英文字母组成
 *
 *
 */
const MAXN = 30 + 5

type S struct{ p, q, x, y int }

func same(b1, b2 []byte, s S) bool {
	for i := 0; i < s.q-s.p; i++ {
		if b1[s.p+i] != b2[s.x+i] {
			return false
		}
	}
	return true
}

func freq(b1, b2 []byte, s S) bool {
	var A [26]int
	for i := s.p; i < s.q; i++ {
		A[b1[i]-'a']++
	}
	for i := s.x; i < s.y; i++ {
		A[b2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if A[i] != 0 {
			return false
		}
	}
	return true
}

func dfs(b1, b2 []byte, s S, M map[S]bool) bool {
	if s.q-s.p != s.y-s.x {
		return false
	}
	if v, ok := M[s]; ok {
		return v
	}

	v := false
	defer func() { M[s] = v }()

	if same(b1, b2, s) {
		v = true
		return v
	}

	if !freq(b1, b2, s) {
		v = false
		return v
	}

	for i := 1; i < s.q-s.p; i++ {
		j := (s.q - s.p) - i
		if (dfs(b1, b2, S{s.p, s.p + i, s.x, s.x + i}, M) && dfs(b1, b2, S{s.p + i, s.q, s.x + i, s.y}, M)) ||
			(dfs(b1, b2, S{s.p, s.p + i, s.x + j, s.y}, M) && dfs(b1, b2, S{s.p + i, s.q, s.x, s.x + j}, M)) {
			v = true
			return v
		}
	}

	v = false
	return v
}

func isScramble(s1 string, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)
	n := len(b1)
	// assert len(s1) == len(s2)

	M := make(map[S]bool)
	return dfs(b1, b2, S{0, n, 0, n}, M)
}
