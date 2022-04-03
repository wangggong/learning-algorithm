/*
 * @lc app=leetcode.cn id=stone-game-v lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1563] 石子游戏 V
 *
 * https://leetcode-cn.com/problems/stone-game-v/description/
 *
 * algorithms
 * Hard (38.40%)
 * Total Accepted:    4.3K
 * Total Submissions: 11.1K
 * Testcase Example:  '[6,2,3,4,5,5]'
 *
 * 几块石子 排成一行 ，每块石子都有一个关联值，关联值为整数，由数组 stoneValue 给出。
 *
 * 游戏中的每一轮：Alice 会将这行石子分成两个 非空行（即，左侧行和右侧行）；Bob 负责计算每一行的值，即此行中所有石子的值的总和。Bob
 * 会丢弃值最大的行，Alice 的得分为剩下那行的值（每轮累加）。如果两行的值相等，Bob 让 Alice 决定丢弃哪一行。下一轮从剩下的那一行开始。
 *
 * 只 剩下一块石子 时，游戏结束。Alice 的分数最初为 0 。
 *
 * 返回 Alice 能够获得的最大分数 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：stoneValue = [6,2,3,4,5,5]
 * 输出：18
 * 解释：在第一轮中，Alice 将行划分为 [6，2，3]，[4，5，5] 。左行的值是 11 ，右行的值是 14 。Bob 丢弃了右行，Alice
 * 的分数现在是 11 。
 * 在第二轮中，Alice 将行分成 [6]，[2，3] 。这一次 Bob 扔掉了左行，Alice 的分数变成了 16（11 + 5）。
 * 最后一轮 Alice 只能将行分成 [2]，[3] 。Bob 扔掉右行，Alice 的分数现在是 18（16 +
 * 2）。游戏结束，因为这行只剩下一块石头了。
 *
 *
 * 示例 2：
 *
 * 输入：stoneValue = [7,7,7,7,7,7,7]
 * 输出：28
 *
 *
 * 示例 3：
 *
 * 输入：stoneValue = [4]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= stoneValue.length <= 500
 * 1 <= stoneValue[i] <= 10^6
 *
 *
 */
const MAXN = 500 + 10

// const MAXN = 10

var dp [MAXN][MAXN]int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(arr, preSum []int, p, q int) int {
	if dp[p][q] >= 0 {
		return dp[p][q]
	}
	if p >= q {
		dp[p][q] = 0
		return 0
	}
	if p+1 == q {
		dp[p][q] = arr[p]
		return dp[p][q]
	}
	ans := 0
	for i := p + 1; i < q; i++ {
		//   0 1 2 3
		// X 2 3 4
		// 0 2 5 9
		left := preSum[i] - preSum[p]
		right := preSum[q] - preSum[i]
		cur := 0
		rest := 0
		if left < right {
			if i > p+1 {
				rest = dfs(arr, preSum, p, i)
			}
			cur = left + rest
		} else if left > right {
			if q > i+1 {
				rest = dfs(arr, preSum, i, q)
			}
			cur = right + rest
		} else {
			v1, v2 := 0, 0
			if i > p+1 {
				v1 = dfs(arr, preSum, p, i)
			}
			if q > i+1 {
				v2 = dfs(arr, preSum, i, q)
			}
			cur = left + max(v1, v2)
		}
		ans = max(ans, cur)
	}
	dp[p][q] = ans
	return ans
}

func stoneGameV(stoneValue []int) int {
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			dp[i][j] = -1
		}
	}
	n := len(stoneValue)
	if n == 1 {
		return 0
	}
	prefixSum := make([]int, n+1)
	for i, v := range stoneValue {
		prefixSum[i+1] = prefixSum[i] + v
	}
	v := dfs(stoneValue, prefixSum, 0, n)
	// fmt.Printf("%v\n%v\n", prefixSum, dp)
	return v
}
