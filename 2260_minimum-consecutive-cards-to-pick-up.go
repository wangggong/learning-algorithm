/*
 * @lc app=leetcode.cn id=minimum-consecutive-cards-to-pick-up lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2260] 必须拿起的最小连续卡牌数
 *
 * https://leetcode-cn.com/problems/minimum-consecutive-cards-to-pick-up/description/
 *
 * algorithms
 * Medium (48.91%)
 * Total Accepted:    7.4K
 * Total Submissions: 14.9K
 * Testcase Example:  '[3,4,2,3,4,7]'
 *
 * 给你一个整数数组 cards ，其中 cards[i] 表示第 i 张卡牌的 值 。如果两张卡牌的值相同，则认为这一对卡牌 匹配 。
 *
 * 返回你必须拿起的最小连续卡牌数，以使在拿起的卡牌中有一对匹配的卡牌。如果无法得到一对匹配的卡牌，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：cards = [3,4,2,3,4,7]
 * 输出：4
 * 解释：拿起卡牌 [3,4,2,3] 将会包含一对值为 3 的匹配卡牌。注意，拿起 [4,2,3,4] 也是最优方案。
 *
 * 示例 2：
 *
 * 输入：cards = [1,0,5,3]
 * 输出：-1
 * 解释：无法找出含一对匹配卡牌的一组连续卡牌。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= cards.length <= 10^5
 * 0 <= cards[i] <= 10^6
 *
 *
 */
func minimumCardPickup(cards []int) int {
	n := len(cards)
	m := make(map[int]int)
	ans := n + 1
	for i := 0; i < n; i++ {
		if v, ok := m[cards[i]]; ok {
			ans = min(ans, i-v+1)
		}
		m[cards[i]] = i
	}
	if ans > n {
		return -1
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
