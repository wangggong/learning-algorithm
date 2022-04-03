/*
 * @lc app=leetcode.cn id=stone-game lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [877] 石子游戏
 *
 * https://leetcode-cn.com/problems/stone-game/description/
 *
 * algorithms
 * Medium (75.84%)
 * Total Accepted:    73.2K
 * Total Submissions: 96.4K
 * Testcase Example:  '[5,3,4,5]'
 *
 * Alice 和 Bob 用几堆石子在做游戏。一共有偶数堆石子，排成一行；每堆都有 正 整数颗石子，数目为 piles[i] 。
 *
 * 游戏以谁手中的石子最多来决出胜负。石子的 总数 是 奇数 ，所以没有平局。
 *
 * Alice 和 Bob 轮流进行，Alice 先开始 。 每回合，玩家从行的 开始 或 结束 处取走整堆石头。
 * 这种情况一直持续到没有更多的石子堆为止，此时手中 石子最多 的玩家 获胜 。
 *
 * 假设 Alice 和 Bob 都发挥出最佳水平，当 Alice 赢得比赛时返回 true ，当 Bob 赢得比赛时返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：piles = [5,3,4,5]
 * 输出：true
 * 解释：
 * Alice 先开始，只能拿前 5 颗或后 5 颗石子 。
 * 假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
 * 如果 Bob 拿走前 3 颗，那么剩下的是 [4,5]，Alice 拿走后 5 颗赢得 10 分。
 * 如果 Bob 拿走后 5 颗，那么剩下的是 [3,4]，Alice 拿走后 4 颗赢得 9 分。
 * 这表明，取前 5 颗石子对 Alice 来说是一个胜利的举动，所以返回 true 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：piles = [3,7,2,3]
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= piles.length <= 500
 * piles.length 是 偶数
 * 1 <= piles[i] <= 500
 * sum(piles[i]) 是 奇数
 *
 *
 */

const (
	turnAlice = iota
	turnBob
)

const MAXN = 500 + 10

var dp [2][MAXN][MAXN]int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minMax(arr []int, p, q, turn int) int {
	if p >= q {
		return 0
	}
	if dp[turn][p][q] >= 0 {
		return dp[turn][p][q]
	}
	var nextTurn int
	switch turn {
	case turnAlice:
		nextTurn = turnBob
	case turnBob:
		nextTurn = turnAlice
	}
	v1 := minMax(arr, p+1, q, nextTurn) + arr[p]
	v2 := minMax(arr, p, q-1, nextTurn) + arr[q-1]
	switch turn {
	case turnAlice:
		v := max(v1, v2)
		dp[turn][p][q] = v
		return v
	case turnBob:
		v := min(v1, v2)
		dp[turn][p][q] = v
		return v
	}
	return 0
}

func stoneGame(piles []int) bool {
	n := len(piles)
	for k := 0; k < 2; k++ {
		for i := 0; i < MAXN; i++ {
			for j := 0; j < MAXN; j++ {
				dp[k][i][j] = -1
			}
		}
	}
	v := minMax(piles, 0, n, turnAlice)
	total := 0
	for _, p := range piles {
		total += p
	}
	return v*2 >= total
}
