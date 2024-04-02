/*
 * @lc app=leetcode.cn id=bulb-switcher-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [672] 灯泡开关 Ⅱ
 *
 * https://leetcode.cn/problems/bulb-switcher-ii/description/
 *
 * algorithms
 * Medium (53.93%)
 * Total Accepted:    25.3K
 * Total Submissions: 41.3K
 * Testcase Example:  '1\n1'
 *
 * 房间中有 n 只已经打开的灯泡，编号从 1 到 n 。墙上挂着 4 个开关 。
 *
 * 这 4 个开关各自都具有不同的功能，其中：
 *
 *
 * 开关 1 ：反转当前所有灯的状态（即开变为关，关变为开）
 * 开关 2 ：反转编号为偶数的灯的状态（即 2, 4, ...）
 * 开关 3 ：反转编号为奇数的灯的状态（即 1, 3, ...）
 * 开关 4 ：反转编号为 j = 3k + 1 的灯的状态，其中 k = 0, 1, 2, ...（即 1, 4, 7, 10, ...）
 *
 *
 * 你必须 恰好 按压开关 presses 次。每次按压，你都需要从 4 个开关中选出一个来执行按压操作。
 *
 * 给你两个整数 n 和 presses ，执行完所有按压之后，返回 不同可能状态 的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 1, presses = 1
 * 输出：2
 * 解释：状态可以是：
 * - 按压开关 1 ，[关]
 * - 按压开关 2 ，[开]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 2, presses = 1
 * 输出：3
 * 解释：状态可以是：
 * - 按压开关 1 ，[关, 关]
 * - 按压开关 2 ，[开, 关]
 * - 按压开关 3 ，[关, 开]
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 3, presses = 1
 * 输出：4
 * 解释：状态可以是：
 * - 按压开关 1 ，[关, 关, 关]
 * - 按压开关 2 ，[关, 开, 关]
 * - 按压开关 3 ，[开, 开, 开]
 * - 按压开关 4 ，[关, 开, 开]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 1000
 * 0 <= presses <= 1000
 *
 *
 */

// 参考: https://leetcode.cn/problems/bulb-switcher-ii/solution/by-ac_oier-3ttx/
// 找——规——律——
func flipLights(n, p int) int {
	// 不按按键: 只有全开状态
	if p == 0 {
		return 1
	}
	// 只有一盏灯: 但凡按一下都能跑满状态
	if n == 1 {
		return 2
	}
	// 两盏灯: 按一下 3 种状态 (00, 10, 01), 再按一下能跑满状态
	if n == 2 {
		if p == 1 {
			return 3
		}
		return 4
	}
	// 至少三盏灯: 和三盏灯相同.
	// 证明分两部分: 1. 六盏灯一周期; 2. 后引入的灯泡与之前的三个灯泡的状态一一对应.
	// 因此直接考虑三盏灯: 按一下 4 种状态 (000, 010, 011, 101), 两下 7 种 (000, 001, 010, 100, 101, 110, 111), 三下拉满
	if p == 1 {
		return 4
	}
	if p == 2 {
		return 7
	}
	return 8
}
