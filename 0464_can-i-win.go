/*
 * @lc app=leetcode.cn id=can-i-win lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [464] 我能赢吗
 *
 * https://leetcode-cn.com/problems/can-i-win/description/
 *
 * algorithms
 * Medium (35.51%)
 * Total Accepted:    18.2K
 * Total Submissions: 47.9K
 * Testcase Example:  '10\n11'
 *
 * 在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和 达到或超过  100
 * 的玩家，即为胜者。
 *
 * 如果我们将游戏规则改为 “玩家 不能 重复使用整数” 呢？
 *
 * 例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
 *
 * 给定两个整数 maxChoosableInteger （整数池中可选择的最大数）和 desiredTotal（累计和），若先出手的玩家是否能稳赢则返回
 * true ，否则返回 false 。假设两位玩家游戏时都表现 最佳 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：maxChoosableInteger = 10, desiredTotal = 11
 * 输出：false
 * 解释：
 * 无论第一个玩家选择哪个整数，他都会失败。
 * 第一个玩家可以选择从 1 到 10 的整数。
 * 如果第一个玩家选择 1，那么第二个玩家只能选择从 2 到 10 的整数。
 * 第二个玩家可以通过选择整数 10（那么累积和为 11 >= desiredTotal），从而取得胜利.
 * 同样地，第一个玩家选择任意其他整数，第二个玩家都会赢。
 *
 *
 * 示例 2:
 *
 *
 * 输入：maxChoosableInteger = 10, desiredTotal = 0
 * 输出：true
 *
 *
 * 示例 3:
 *
 *
 * 输入：maxChoosableInteger = 10, desiredTotal = 1
 * 输出：true
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= maxChoosableInteger <= 20
 * 0 <= desiredTotal <= 300
 *
 *
 */

const (
	alice = iota
	bob
)

func canIWin(n int, target int) bool {
	if n*(n+1)>>1 < target {
		return false
	}
	memo := make(map[[3]int]int)
	return dfs(n, target, 0, alice, memo) == alice
}

func dfs(n int, target int, visit int, turn int, memo map[[3]int]int) (ans int) {
	status := [3]int{target, visit, turn}
	if ans, ok := memo[status]; ok {
		return ans
	}
	defer func() { memo[status] = ans }()
	// defer func() { fmt.Printf("%v %010b %v %v\n", target, visit, turn, ans) }()
	switch turn {
	case alice:
		ans = bob
		for i := 1; i <= n; i++ {
			if visit&(1<<i) != 0 {
				continue
			}
			if target <= i {
				ans = alice
				return
			}
			if dfs(n, target-i, visit|(1<<i), bob, memo) == alice {
				ans = alice
				return
			}
		}
		return
	case bob:
		ans = alice
		for i := 1; i <= n; i++ {
			if visit&(1<<i) != 0 {
				continue
			}
			if target <= i {
				ans = bob
				return
			}
			if dfs(n, target-i, visit|(1<<i), alice, memo) == bob {
				ans = bob
				return
			}
		}
		return
	default:
	}
	// unreachable
	return alice
}

