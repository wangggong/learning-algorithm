/*
 * @lc app=leetcode.cn id=stone-game-vii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1690] 石子游戏 VII
 *
 * https://leetcode-cn.com/problems/stone-game-vii/description/
 *
 * algorithms
 * Medium (52.52%)
 * Total Accepted:    4.9K
 * Total Submissions: 9.2K
 * Testcase Example:  '[5,3,1,4,2]'
 *
 * 石子游戏中，爱丽丝和鲍勃轮流进行自己的回合，爱丽丝先开始 。
 *
 * 有 n 块石子排成一排。每个玩家的回合中，可以从行中 移除 最左边的石头或最右边的石头，并获得与该行中剩余石头值之 和
 * 相等的得分。当没有石头可移除时，得分较高者获胜。
 *
 * 鲍勃发现他总是输掉游戏（可怜的鲍勃，他总是输），所以他决定尽力 减小得分的差值 。爱丽丝的目标是最大限度地 扩大得分的差值 。
 *
 * 给你一个整数数组 stones ，其中 stones[i] 表示 从左边开始 的第 i 个石头的值，如果爱丽丝和鲍勃都 发挥出最佳水平 ，请返回他们
 * 得分的差值 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：stones = [5,3,1,4,2]
 * 输出：6
 * 解释：
 * - 爱丽丝移除 2 ，得分 5 + 3 + 1 + 4 = 13 。游戏情况：爱丽丝 = 13 ，鲍勃 = 0 ，石子 = [5,3,1,4] 。
 * - 鲍勃移除 5 ，得分 3 + 1 + 4 = 8 。游戏情况：爱丽丝 = 13 ，鲍勃 = 8 ，石子 = [3,1,4] 。
 * - 爱丽丝移除 3 ，得分 1 + 4 = 5 。游戏情况：爱丽丝 = 18 ，鲍勃 = 8 ，石子 = [1,4] 。
 * - 鲍勃移除 1 ，得分 4 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [4] 。
 * - 爱丽丝移除 4 ，得分 0 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [] 。
 * 得分的差值 18 - 12 = 6 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：stones = [7,90,5,1,100,10,10,2]
 * 输出：122
 *
 *
 *
 * 提示：
 *
 *
 * n == stones.length
 * 2 <= n <= 1000
 * 1 <= stones[i] <= 1000
 *
 *
 */
const (
	aliceTurn = iota
	bobTurn
)

const MAXN = 1e3 + 10

var n int

type Pair struct{ x, y int }

// var dp [2][MAXN][MAXN]int
var dp [MAXN][MAXN]int

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

/*
 * func stoneGameVII(stones []int) int {
 * 	n = len(stones)
 * 	for k := 0; k < 2; k++ {
 * 		for i := 0; i < MAXN; i++ {
 * 			for j := 0; j < MAXN; j++ {
 * 				dp[k][i][j] = -1
 * 			}
 * 		}
 * 	}
 * 	prefixSum := make([]int, n+1)
 * 	for i, v := range stones {
 * 		prefixSum[i+1] = prefixSum[i] + v
 * 	}
 * 	v := dfs(stones, prefixSum, aliceTurn, 0, n)
 * 	//for k := 0; k < 2; k++ {
 * 	//	for i := 0; i < 7; i++ {
 * 	//		for j := 0; j < 7; j++ {
 * 	//			fmt.Printf("%v ", dp[k][i][j])
 * 	//		}
 * 	//		fmt.Println()
 * 	//	}
 * 	//	fmt.Println()
 * 	//}
 * 	return v
 * }
 *
 * func dfs(arr, prefix []int, turn, p, q int) int {
 * 	if p >= q {
 * 		return 0
 * 	}
 * 	if dp[turn][p][q] >= 0 {
 * 		return dp[turn][p][q]
 * 	}
 * 	// X 5 3 1  4  2
 * 	// 0 5 8 9 13 15
 * 	v1 := prefix[q] - prefix[p+1]
 * 	v2 := prefix[q-1] - prefix[p]
 * 	switch turn {
 * 	case aliceTurn:
 * 		dp[turn][p][q] = max(
 * 			v1+dfs(arr, prefix, bobTurn, p+1, q),
 * 			v2+dfs(arr, prefix, bobTurn, p, q-1),
 * 		)
 * 	case bobTurn:
 * 		dp[turn][p][q] = min(
 * 			dfs(arr, prefix, aliceTurn, p+1, q)-v1,
 * 			dfs(arr, prefix, aliceTurn, p, q-1)-v2,
 * 		)
 * 	default:
 * 	}
 * 	return dp[turn][p][q]
 * }
 */

func stoneGameVII(stones []int) int {
	n = len(stones)
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			dp[i][j] = -1
		}
	}
	return maxMin(stones, 0, n)
}

func maxMin(arr []int, p, q int) int {
	if dp[p][q] >= 0 {
		return dp[p][q]
	}
	if p >= q {
		dp[p][q] = 0
		return dp[p][q]
	}
	if p+1 == q {
		dp[p][q] = 0
		return dp[p][q]
	}
	if p+1 == q-1 {
		dp[p][q] = max(arr[p], arr[q-1])
		return dp[p][q]
	}
	// 得分差值仅取决于 Bob 选择的价值.
	dp[p][q] = max(
		min(
			arr[p+1]+maxMin(arr, p+2, q),
			arr[q-1]+maxMin(arr, p+1, q-1),
		),
		min(
			arr[p]+maxMin(arr, p+1, q-1),
			arr[q-2]+maxMin(arr, p, q-2),
		),
	)
	return dp[p][q]
}
