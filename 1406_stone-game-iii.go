/*
 * @lc app=leetcode.cn id=stone-game-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1406] 石子游戏 III
 *
 * https://leetcode-cn.com/problems/stone-game-iii/description/
 *
 * algorithms
 * Hard (56.75%)
 * Total Accepted:    5.6K
 * Total Submissions: 9.8K
 * Testcase Example:  '[1,2,3,7]'
 *
 * Alice 和 Bob 用几堆石子在做游戏。几堆石子排成一行，每堆石子都对应一个得分，由数组 stoneValue 给出。
 *
 * Alice 和 Bob 轮流取石子，Alice 总是先开始。在每个玩家的回合中，该玩家可以拿走剩下石子中的的前 1、2 或 3 堆石子
 * 。比赛一直持续到所有石头都被拿走。
 *
 * 每个玩家的最终得分为他所拿到的每堆石子的对应得分之和。每个玩家的初始分数都是 0
 * 。比赛的目标是决出最高分，得分最高的选手将会赢得比赛，比赛也可能会出现平局。
 *
 * 假设 Alice 和 Bob 都采取 最优策略 。如果 Alice 赢了就返回 "Alice" ，Bob 赢了就返回 "Bob"，平局（分数相同）返回
 * "Tie" 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：values = [1,2,3,7]
 * 输出："Bob"
 * 解释：Alice 总是会输，她的最佳选择是拿走前三堆，得分变成 6 。但是 Bob 的得分为 7，Bob 获胜。
 *
 *
 * 示例 2：
 *
 * 输入：values = [1,2,3,-9]
 * 输出："Alice"
 * 解释：Alice 要想获胜就必须在第一个回合拿走前三堆石子，给 Bob 留下负分。
 * 如果 Alice 只拿走第一堆，那么她的得分为 1，接下来 Bob 拿走第二、三堆，得分为 5 。之后 Alice 只能拿到分数 -9
 * 的石子堆，输掉比赛。
 * 如果 Alice 拿走前两堆，那么她的得分为 3，接下来 Bob 拿走第三堆，得分为 3 。之后 Alice 只能拿到分数 -9
 * 的石子堆，同样会输掉比赛。
 * 注意，他们都应该采取 最优策略 ，所以在这里 Alice 将选择能够使她获胜的方案。
 *
 * 示例 3：
 *
 * 输入：values = [1,2,3,6]
 * 输出："Tie"
 * 解释：Alice 无法赢得比赛。如果她决定选择前三堆，她可以以平局结束比赛，否则她就会输。
 *
 *
 * 示例 4：
 *
 * 输入：values = [1,2,3,-1,-2,-3,7]
 * 输出："Alice"
 *
 *
 * 示例 5：
 *
 * 输入：values = [-1,-2,-3]
 * 输出："Tie"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= values.length <= 50000
 * -1000 <= values[i] <= 1000
 *
 *
 */
// import "fmt"
const (
	aliceWin = "Alice"
	bobWin   = "Bob"
	tie      = "Tie"
)

const K = 3
const MAXN = 50000 + 5

/*
 * type Pair struct{ x, y int }
 *
 * var dp [MAXN]*Pair
 * var n int
 *
 * func dfs(arr []int, p int) (a, b int) {
 * 	if p >= n {
 * 		return 0, 0
 * 	}
 * 	if p := dp[p]; p != nil {
 * 		return p.x, p.y
 * 	}
 * 	a = -1e8
 * 	cur := 0
 * 	for i := 0; i < K; i++ {
 * 		v := p + i
 * 		if v >= n {
 * 			continue
 * 		}
 * 		cur += arr[v]
 * 		na, nb := dfs(arr, v+1)
 * 		// fmt.Printf("p = %v, i = %v, cur = %v, na = %v, nb = %v\n", p, i, cur, na, nb)
 * 		if cur+nb > a {
 * 			a, b = cur+nb, na
 * 		}
 * 	}
 * 	dp[p] = &Pair{a, b}
 * 	return
 * }
 *
 * func stoneGameIII(stoneValue []int) string {
 * 	for i := 0; i < MAXN; i++ {
 * 		dp[i] = nil
 * 	}
 * 	n = len(stoneValue)
 * 	alice, bob := dfs(stoneValue, 0)
 * 	// fmt.Printf("%v %v\n", alice, bob)
 * 	switch c := alice - bob; {
 * 	case c > 0:
 * 		return aliceWin
 * 	case c < 0:
 * 		return bobWin
 * 	default:
 * 		return tie
 * 	}
 * }
 */

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	suffixSum := make([]int, n+1)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + stoneValue[i]
	}
	for i := n - 1; i >= 0; i-- {
		cur := dp[i+1]
		for j := 2; j <= K && i+j <= n; j++ {
			if dp[i+j] < cur {
				cur = dp[i+j]
			}
		}
		dp[i] = suffixSum[i] - cur
	}
	total := suffixSum[0]
	// fmt.Printf("%v %v\n", suffixSum, dp)
	switch c := dp[0] * 2; {
	case c > total:
		return aliceWin
	case c < total:
		return bobWin
	default:
		return tie
	}
}
