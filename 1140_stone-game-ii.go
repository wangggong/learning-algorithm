import "math"

/*
 * @lc app=leetcode.cn id=stone-game-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1140] 石子游戏 II
 *
 * https://leetcode-cn.com/problems/stone-game-ii/description/
 *
 * algorithms
 * Medium (65.34%)
 * Total Accepted:    7.5K
 * Total Submissions: 11.4K
 * Testcase Example:  '[2,7,9,4,4]'
 *
 * 爱丽丝和鲍勃继续他们的石子游戏。许多堆石子排成一行，每堆都有正整数颗石子piles[i]。游戏以谁手中的石子最多来决出胜负。
 *
 * 爱丽丝和鲍勃轮流进行，爱丽丝先开始。最初，M = 1。
 *
 * 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
 *
 * 游戏一直持续到所有石子都被拿走。
 *
 * 假设爱丽丝和鲍勃都发挥出最佳水平，返回爱丽丝可以得到的最大数量的石头。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：piles = [2,7,9,4,4]
 * 输出：10
 * 解释：如果一开始Alice取了一堆，Bob取了两堆，然后Alice再取两堆。爱丽丝可以得到2 + 4 + 4 =
 * 10堆。如果Alice一开始拿走了两堆，那么Bob可以拿走剩下的三堆。在这种情况下，Alice得到2 + 7 = 9堆。返回10，因为它更大。
 *
 *
 * 示例 2:
 *
 *
 * 输入：piles = [1,2,3,4,5,100]
 * 输出：104
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= piles.length <= 100
 * 1 <= piles[i] <= 10^4
 *
 *
 */

const (
	turnAlice = iota
	turnBob
)

const MAXN = 100 + 10

type P struct{ x, y int }

var dp [MAXN][MAXN]P
var n int

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

func minMax(arr []int, p, M int) (a, b int) {
	if p == n {
		return 0, 0
	}
	if v := dp[p][M]; v.x >= 0 {
		return v.x, v.y
	}
	cur := 0
	for i := 0; i < 2*M; i++ {
		if p+i >= n {
			continue
		}
		cur += arr[p+i]
		nM := M
		if i+1 > M {
			nM = i + 1
		}
		na, nb := minMax(arr, p+i+1, nM)
		if cur+nb > a {
			a = cur + nb
			b = na
		}
	}
	dp[p][M] = P{a, b}
	return
}

func stoneGameII(piles []int) int {
	n = len(piles)
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			dp[i][j] = P{-1, -1}
		}
	}
	v, _ := minMax(piles, 0, 1)
	return v
}

type Key struct{ K, M int }

func dfs(S []int, k, M int, memo map[Key]int) (ans int) {
	key := Key{k, M}
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	defer func() { memo[key] = ans }()
	if k+2*M >= len(S) {
		ans = S[k]
		return
	}
	other := math.MaxInt32
	for i := 1; i <= 2*M; i++ {
		other = min(other, dfs(S, k+i, max(M, i), memo))
	}
	ans = S[k] - other
	return
}

/*
 * func stoneGameII(S []int) int {
 * 	n := len(S)
 * 	for i := n - 2; i >= 0; i-- {
 * 		S[i] += S[i+1]
 * 	}
 * 	return dfs(S, 0, 1, make(map[Key]int))
 * }
 */
