/*
 * @lc app=leetcode.cn id=remove-letter-to-equalize-frequency lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6212] 删除字符使频率相同
 *
 * https://leetcode.cn/problems/remove-letter-to-equalize-frequency/description/
 *
 * algorithms
 * Easy (18.93%)
 * Total Accepted:    3K
 * Total Submissions: 15.7K
 * Testcase Example:  '"abcc"'
 *
 * 给你一个下标从 0 开始的字符串 word ，字符串只包含小写英文字母。你需要选择 一个 下标并 删除 下标处的字符，使得 word 中剩余每个字母出现
 * 频率 相同。
 *
 * 如果删除一个字母后，word 中剩余所有字母的出现频率都相同，那么返回 true ，否则返回 false 。
 *
 * 注意：
 *
 *
 * 字母 x 的 频率 是这个字母在字符串中出现的次数。
 * 你 必须 恰好删除一个字母，不能一个字母都不删除。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：word = "abcc"
 * 输出：true
 * 解释：选择下标 3 并删除该字母，word 变成 "abc" 且每个字母出现频率都为 1 。
 *
 *
 * 示例 2：
 *
 * 输入：word = "aazz"
 * 输出：false
 * 解释：我们必须删除一个字母，所以要么 "a" 的频率变为 1 且 "z" 的频率为 2
 * ，要么两个字母频率反过来。所以不可能让剩余所有字母出现频率相同。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= word.length <= 100
 * word 只包含小写英文字母。
 *
 *
 */

func equalFrequency(word string) bool {
	cnt := make(map[byte]int)
	for _, b := range []byte(word) {
		cnt[b]++
	}
	mxb, cnt1b := byte(0), byte(0)
	for k, c := range cnt {
		if c == 1 {
			cnt1b = k
		}
		if c > cnt[mxb] {
			mxb = k
		}
	}
	if cnt1b != 0 {
		n, ok := -1, true
		for k, c := range cnt {
			if k == cnt1b {
				continue
			}
			if n == -1 {
				n = c
				continue
			}
			if n != c {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	if mxb != 0 {
		n, ok := cnt[mxb]-1, true
		for k, c := range cnt {
			if k == mxb {
				continue
			}
			if c != n {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
