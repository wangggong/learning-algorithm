/*
 * @lc app=leetcode.cn id=the-number-of-weak-characters-in-the-game lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1996] 游戏中弱角色的数量
 *
 * https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game/description/
 *
 * algorithms
 * Medium (25.84%)
 * Total Accepted:    25.6K
 * Total Submissions: 62K
 * Testcase Example:  '[[5,5],[6,3],[3,6]]'
 *
 * 你正在参加一个多角色游戏，每个角色都有两个主要属性：攻击 和 防御 。给你一个二维整数数组 properties ，其中 properties[i] =
 * [attacki, defensei] 表示游戏中第 i 个角色的属性。
 *
 * 如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色 。更正式地，如果认为角色 i 弱于 存在的另一个角色
 * j ，那么 attackj > attacki 且 defensej > defensei 。
 *
 * 返回 弱角色 的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：properties = [[5,5],[6,3],[3,6]]
 * 输出：0
 * 解释：不存在攻击和防御都严格高于其他角色的角色。
 *
 *
 * 示例 2：
 *
 *
 * 输入：properties = [[2,2],[3,3]]
 * 输出：1
 * 解释：第一个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。
 *
 *
 * 示例 3：
 *
 *
 * 输入：properties = [[1,5],[10,4],[4,3]]
 * 输出：1
 * 解释：第三个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= properties.length <= 10^5
 * properties[i].length == 2
 * 1 <= attacki, defensei <= 10^5
 *
 *
 */

import "sort"

type PropSlice struct {
	Prop [][]int
}

func (ps PropSlice) Len() int { return len(ps.Prop) }
func (ps PropSlice) Less(p, q int) bool {
	if ps.Prop[p][0] != ps.Prop[q][0] {
		return ps.Prop[p][0] > ps.Prop[q][0]
	}
	return ps.Prop[p][1] < ps.Prop[q][1]
}
func (ps PropSlice) Swap(p, q int) { ps.Prop[p], ps.Prop[q] = ps.Prop[q], ps.Prop[p] }

// [[5 5] [6 3] [3 6]]
// [[3 6] [5 5] [6 3]]
func numberOfWeakCharacters(properties [][]int) int {
	// assert len(properties) > 1 && len(properties) <= 1e5
	// assert attack[i] > 0 && attack[i] <= 1e5
	// assert defense[i] > 0 && defense[i] <= 1e5
	ps := PropSlice{properties}
	sort.Sort(ps)
	ans := 0
	maxDef := 0
	for _, p := range ps.Prop {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	// S := make([][]int, 0, len(ps.Prop))
	// for _, p := range ps.Prop {
	// 	for len(S) > 0 && S[len(S)-1][1] <= p[1] {
	// 		if S[len(S)-1][0] < p[0] && S[len(S)-1][1] < p[1] {
	// 			ans++
	// 		}
	// 		S = S[:len(S)-1]
	// 	}
	// 	S = append(S, p)
	// }
	return ans
}
