/*
 * @lc app=leetcode.cn id=freedom-trail lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [514] 自由之路
 *
 * https://leetcode-cn.com/problems/freedom-trail/description/
 *
 * algorithms
 * Hard (50.12%)
 * Total Accepted:    21.8K
 * Total Submissions: 43.5K
 * Testcase Example:  '"godding"\n"gd"'
 *
 * 电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，并使用表盘拼写特定关键词才能开门。
 *
 * 给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
 *
 * 最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00
 * 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。
 *
 * 旋转 ring 拼出 key 字符 key[i] 的阶段中：
 *
 *
 * 您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00
 * 方向对齐，并且这个字符必须等于字符 key[i] 。
 * 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key
 * 的下一个字符（下一阶段）, 直至完成所有拼写。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 *
 *
 * 输入: ring = "godding", key = "gd"
 * 输出: 4
 * 解释:
 * ⁠对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 * ⁠对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 * ⁠当然, 我们还需要1步进行拼写。
 * ⁠因此最终的输出是 4。
 *
 *
 * 示例 2:
 *
 *
 * 输入: ring = "godding", key = "godding"
 * 输出: 13
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= ring.length, key.length <= 100
 * ring 和 key 只包含小写英文字母
 * 保证 字符串 key 一定可以由字符串  ring 旋转拼出
 *
 *
 */

import "math"

func findRotateSteps(ring string, key string) int {
	pattern, text := []byte(ring), []byte(key)
	m, n := len(pattern), len(text)
	var pos [26][]int
	for i, p := range pattern {
		ind := int(p - 'a')
		pos[ind] = append(pos[ind], i)
	}
	// `dp[i][j]` 代表在 `ring[j]` 位置匹配到 `key[0..i]` 子串的最小步数.
	//
	// 状态转移:
	// 枚举 `i-1` 位置的字符在环中的位置, 假设为 `{a_{0}, a_{1}, ..., a_{k}}`;
	// 枚举 `i` 位置的字符在环中的位置, 假设为 `{b_{0}, b_{1}, ..., b_{j}}`;
	// 那么 `dp[i][j] = min(dp[i-1][k] + min(|j-k|, n-|j-k|) + 1)`, `n` 为环的大小.
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for _, i := range pos[int(text[0]-'a')] {
		dp[0][i] = min(i, m-i) + 1
	}
	for i := 1; i < n; i++ {
		for _, j := range pos[int(text[i]-'a')] {
			for _, k := range pos[int(text[i-1]-'a')] {
				dist := abs(j - k)
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(dist, m-dist)+1)
			}
		}
	}
	ans := math.MaxInt32
	for i := 0; i < m; i++ {
		ans = min(ans, dp[n-1][i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// 二刷后看见状态转移勉强能写出来, 很难啊!
