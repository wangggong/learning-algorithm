/*
 * @lc app=leetcode.cn id=delete-and-earn lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [740] 删除并获得点数
 *
 * https://leetcode-cn.com/problems/delete-and-earn/description/
 *
 * algorithms
 * Medium (62.59%)
 * Total Accepted:    75.1K
 * Total Submissions: 120K
 * Testcase Example:  '[3,4,2]'
 *
 * 给你一个整数数组 nums ，你可以对它进行一些操作。
 *
 * 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和
 * nums[i] + 1 的元素。
 *
 * 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,4,2]
 * 输出：6
 * 解释：
 * 删除 4 获得 4 个点数，因此 3 也被删除。
 * 之后，删除 2 获得 2 个点数。总共获得 6 个点数。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,3,3,3,4]
 * 输出：9
 * 解释：
 * 删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
 * 之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
 * 总共获得 9 个点数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 104
 * 1 <= nums[i] <= 104
 *
 *
 */
const maxn int = 1e4

func deleteAndEarn(nums []int) int {
	counter := make([]int, maxn+5)
	for _, n := range nums {
		counter[n]++
	}
	n := len(nums)
	arr := make([][2]int, 0, n)
	for i, c := range counter {
		if c > 0 {
			arr = append(arr, [2]int{i, c})
		}
	}
	m := len(arr)
	if m == 0 {
		return 0
	}
	dp := make([]int, m)
	dp[0] = arr[0][0] * arr[0][1]
	ans := dp[0]
	for i := 1; i < m; i++ {
		dp[i] = arr[i][0] * arr[i][1]
		v := 0
		for j := 0; j < i; j++ {
			if arr[i][0]-arr[j][0] > 1 {
				v = max(v, dp[j])
			}
		}
		dp[i] += v
		ans = max(ans, dp[i])
	}
	// fmt.Println(arr, dp)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
