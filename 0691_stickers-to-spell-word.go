/*
 * @lc app=leetcode.cn id=stickers-to-spell-word lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [691] 贴纸拼词
 *
 * https://leetcode.cn/problems/stickers-to-spell-word/description/
 *
 * algorithms
 * Hard (46.88%)
 * Total Accepted:    8.9K
 * Total Submissions: 16.4K
 * Testcase Example:  '["with","example","science"]\n"thehat"'
 *
 * 我们有 n 种不同的贴纸。每个贴纸上都有一个小写的英文单词。
 *
 * 您想要拼写出给定的字符串 target ，方法是从收集的贴纸中切割单个字母并重新排列它们。如果你愿意，你可以多次使用每个贴纸，每个贴纸的数量是无限的。
 *
 * 返回你需要拼出 target 的最小贴纸数量。如果任务不可能，则返回 -1 。
 *
 * 注意：在所有的测试用例中，所有的单词都是从 1000 个最常见的美国英语单词中随机选择的，并且 target 被选择为两个随机单词的连接。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： stickers = ["with","example","science"], target = "thehat"
 * 输出：3
 * 解释：
 * 我们可以使用 2 个 "with" 贴纸，和 1 个 "example" 贴纸。
 * 把贴纸上的字母剪下来并重新排列后，就可以形成目标 “thehat“ 了。
 * 此外，这是形成目标字符串所需的最小贴纸数量。
 *
 *
 * 示例 2:
 *
 *
 * 输入：stickers = ["notice","possible"], target = "basicbasic"
 * 输出：-1
 * 解释：我们不能通过剪切给定贴纸的字母来形成目标“basicbasic”。
 *
 *
 *
 * 提示:
 *
 *
 * n == stickers.length
 * 1 <= n <= 50
 * 1 <= stickers[i].length <= 10
 * 1 <= target <= 15
 * stickers[i] 和 target 由小写英文单词组成
 *
 *
 */

import "math"

const alphaCnt int = 26

var dp []int

func newState(k int, target, curr string) int {
	for _, c := range curr {
		for i, t := range target {
			if t == c && k&(1<<i) == 0 {
				k = k | 1<<i
				break
			}
		}
	}
	return k
}

func minStickers(stickers []string, target string) int {
	n := len(target)
	lim := 1 << n
	dp = make([]int, lim)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, sticker := range stickers {
		for s := 0; s < lim; s++ {
			ns := newState(s, target, sticker)
			if ns != s {
				dp[ns] = min(dp[ns], dp[s]+1)
			}
		}
	}
	if dp[lim-1] == math.MaxInt32 {
		return -1
	}
	return dp[lim-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
