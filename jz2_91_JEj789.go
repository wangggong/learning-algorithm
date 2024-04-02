/*
 * Medium
 *
 * 假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
 *
 * 当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。
 *
 * 例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。
 *
 * 请计算出粉刷完所有房子最少的花费成本。
 *
 *
 *
 * 示例 1：
 *
 * 输入: costs = [[17,2,17],[16,16,5],[14,3,19]]
 * 输出: 10
 * 解释: 将 0 号房子粉刷成蓝色，1 号房子粉刷成绿色，2 号房子粉刷成蓝色。
 *      最少花费: 2 + 5 + 3 = 10。
 * 示例 2：
 *
 * 输入: costs = [[7,6,2]]
 * 输出: 2
 *
 *
 * 提示:
 *
 * costs.length == n
 * costs[i].length == 3
 * 1 <= n <= 100
 * 1 <= costs[i][j] <= 20
 */

import "math"

func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	copy(dp[0], costs[0])
	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = math.MaxInt32
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+costs[i][j])
			}
		}
	}
	ans := math.MaxInt32
	for _, d := range dp[n-1] {
		ans = min(ans, d)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
 * // 滚动数组
 * func minCost(costs [][]int) int {
 * 	var dp [3]int
 * 	copy(dp[:], costs[0])
 * 	for _, c := range costs[1:] {
 * 		var tmp [3]int
 * 		copy(tmp[:], dp[:])
 * 		dp[0] = min(tmp[1], tmp[2]) + c[0]
 * 		dp[1] = min(tmp[0], tmp[2]) + c[1]
 * 		dp[2] = min(tmp[0], tmp[1]) + c[2]
 * 	}
 * 	return min(min(dp[0], dp[1]), dp[2])
 * }
 * 
 * func min(x, y int) int {
 * 	if x < y {
 * 		return x
 * 	}
 * 	return y
 * }
 */
