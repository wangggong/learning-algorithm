/*
 * @lc app=leetcode.cn id=minimum-operations-to-make-a-subsequence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1713] 得到子序列的最少操作次数
 *
 * https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/description/
 *
 * algorithms
 * Hard (49.53%)
 * Total Accepted:    16.3K
 * Total Submissions: 32.9K
 * Testcase Example:  '[5,1,3]\n[9,4,2,3,4]'
 *
 * 给你一个数组 target ，包含若干 互不相同 的整数，以及另一个整数数组 arr ，arr 可能 包含重复元素。
 *
 * 每一次操作中，你可以在 arr 的任意位置插入任一整数。比方说，如果 arr = [1,4,1,2] ，那么你可以在中间添加 3 得到
 * [1,4,3,1,2] 。你可以在数组最开始或最后面添加整数。
 *
 * 请你返回 最少 操作次数，使得 target 成为 arr 的一个子序列。
 *
 * 一个数组的 子序列 指的是删除原数组的某些元素（可能一个元素都不删除），同时不改变其余元素的相对顺序得到的数组。比方说，[2,7,4] 是
 * [4,2,3,7,2,1,4] 的子序列（加粗元素），但 [2,4,2] 不是子序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：target = [5,1,3], arr = [9,4,2,3,4]
 * 输出：2
 * 解释：你可以添加 5 和 1 ，使得 arr 变为 [5,9,4,1,2,3,4] ，target 为 arr 的子序列。
 *
 *
 * 示例 2：
 *
 * 输入：target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= target.length, arr.length <= 10^5
 * 1 <= target[i], arr[i] <= 10^9
 * target 不包含任何重复元素。
 *
 *
 */
const MAXN = 1e5 + 10

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func invertedIndex(arr []int) map[int]int {
	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}
	return m
}

// LIS 为最长递增子序列 (longest increasing subsequence) 问题二分模版.
func LIS(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	ans := 1
	n := len(arr)
	g := make([]int, 0, n)
	for i := 0; i < n; i++ {
		g = append(g, MAXN)
	}
	g[0] = arr[0]
	l, r, right := 0, 0, 0
	for i := 1; i < n; i++ {
		l, r = 0, right
		for l <= r {
			mid := l + (r-l)/2
			if arr[i] > g[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		right = max(right, l)
		ans = max(ans, l+1)
		g[l] = arr[i]
	}
	return ans
}

func minOperations(target []int, arr []int) int {
	m := invertedIndex(target)
	var helper []int
	for _, v := range arr {
		if k, ok := m[v]; ok {
			helper = append(helper, k)
		}
	}
	v := LIS(helper)
	return len(target) - v
}