/*
 * @lc app=leetcode.cn id=maximum-deletions-on-a-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6195] 对字母串可执行的最大删除数
 *
 * https://leetcode.cn/problems/maximum-deletions-on-a-string/description/
 *
 * algorithms
 * Hard (38.70%)
 * Total Accepted:    4.2K
 * Total Submissions: 9.4K
 * Testcase Example:  '"abcabcdabc"'
 *
 * 给你一个仅由小写英文字母组成的字符串 s 。在一步操作中，你可以：
 *
 *
 * 删除 整个字符串 s ，或者
 * 对于满足 1 <= i <= s.length / 2 的任意 i ，如果 s 中的 前 i 个字母和接下来的 i 个字母 相等 ，删除 前 i
 * 个字母。
 *
 *
 * 例如，如果 s = "ababc" ，那么在一步操作中，你可以删除 s 的前两个字母得到 "abc" ，因为 s 的前两个字母和接下来的两个字母都等于
 * "ab" 。
 *
 * 返回删除 s 所需的最大操作数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abcabcdabc"
 * 输出：2
 * 解释：
 * - 删除前 3 个字母（"abc"），因为它们和接下来 3 个字母相等。现在，s = "abcdabc"。
 * - 删除全部字母。
 * 一共用了 2 步操作，所以返回 2 。可以证明 2 是所需的最大操作数。
 * 注意，在第二步操作中无法再次删除 "abc" ，因为 "abc" 的下一次出现并不是位于接下来的 3 个字母。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aaabaab"
 * 输出：4
 * 解释：
 * - 删除第一个字母（"a"），因为它和接下来的字母相等。现在，s = "aabaab"。
 * - 删除前 3 个字母（"aab"），因为它们和接下来 3 个字母相等。现在，s = "aab"。
 * - 删除第一个字母（"a"），因为它和接下来的字母相等。现在，s = "ab"。
 * - 删除全部字母。
 * 一共用了 4 步操作，所以返回 4 。可以证明 4 是所需的最大操作数。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "aaaaa"
 * 输出：5
 * 解释：在每一步操作中，都可以仅删除 s 的第一个字母。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 4000
 * s 仅由小写英文字母组成
 *
 *
 */

const P uint64 = 13331

var hash, mul []uint64

func get(p, q int) uint64 {
	return hash[q] - hash[p]*mul[q-p]
}

func deleteString(s string) int {
	n := len(s)
	hash, mul = make([]uint64, n+1), make([]uint64, n+1)
	mul[0] = 1
	for i := 0; i < n; i++ {
		mul[i+1] = mul[i] * P
		hash[i+1] = hash[i]*P + uint64(s[i]-'a')
	}
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		for j := 1; i+2*j <= n; j++ {
			if get(i, i+j) == get(i+j, i+2*j) && s[i:i+j] == s[i+j:i+2*j] {
				dp[i] = max(dp[i], dp[i+j]+1)
			}
		}
	}
	return dp[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
