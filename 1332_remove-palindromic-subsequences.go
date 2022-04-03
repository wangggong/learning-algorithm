/*
 * @lc app=leetcode.cn id=remove-palindromic-subsequences lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1332] 删除回文子序列
 *
 * https://leetcode-cn.com/problems/remove-palindromic-subsequences/description/
 *
 * algorithms
 * Easy (69.36%)
 * Total Accepted:    30.4K
 * Total Submissions: 39K
 * Testcase Example:  '"ababa"'
 *
 * 给你一个字符串 s，它仅由字母 'a' 和 'b' 组成。每一次删除操作都可以从 s 中删除一个回文 子序列。
 * 
 * 返回删除给定字符串中所有字符（字符串为空）的最小删除次数。
 * 
 * 「子序列」定义：如果一个字符串可以通过删除原字符串某些字符而不改变原字符顺序得到，那么这个字符串就是原字符串的一个子序列。
 * 
 * 「回文」定义：如果一个字符串向后和向前读是一致的，那么这个字符串就是一个回文。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "ababa"
 * 输出：1
 * 解释：字符串本身就是回文序列，只需要删除一次。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abb"
 * 输出：2
 * 解释："abb" -> "bb" -> "". 
 * 先删除回文子序列 "a"，然后再删除 "bb"。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "baabb"
 * 输出：2
 * 解释："baabb" -> "b" -> "". 
 * 先删除回文子序列 "baab"，然后再删除 "b"。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 1000
 * s 仅包含字母 'a'  和 'b'
 * 
 * 
 */
func removePalindromeSub(s string) int {
	b := []byte(s)
	// containsA, containsB := false, false
	// for _, ch := range b {
	// 	switch ch {
	// 	case 'a':
	// 		containsA = true
	// 	case 'b':
	// 		containsB = true
	// 	}
	// }
	// if (!containsA) || (!containsB) {
	// 	return 1
	// }

	// 如果是回文串, 直接返回 1;
	// 否则返回 2 (先去掉所有 'a', 后去掉所有 'b')
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return 2
		}
	}
	return 1
}
