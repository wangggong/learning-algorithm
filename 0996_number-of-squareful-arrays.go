/*
 * @lc app=leetcode.cn id=number-of-squareful-arrays lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [996] 正方形数组的数目
 *
 * https://leetcode-cn.com/problems/number-of-squareful-arrays/description/
 *
 * algorithms
 * Hard (49.36%)
 * Total Accepted:    4.9K
 * Total Submissions: 9.9K
 * Testcase Example:  '[1,17,8]'
 *
 * 给定一个非负整数数组 A，如果该数组每对相邻元素之和是一个完全平方数，则称这一数组为正方形数组。
 *
 * 返回 A 的正方形排列的数目。两个排列 A1 和 A2 不同的充要条件是存在某个索引 i，使得 A1[i] != A2[i]。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,17,8]
 * 输出：2
 * 解释：
 * [1,8,17] 和 [17,8,1] 都是有效的排列。
 *
 *
 * 示例 2：
 *
 * 输入：[2,2,2]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 12
 * 0 <= A[i] <= 1e9
 *
 *
 */

import "sort"

const MAXN = 15

var n int
var ans int
var cur []int
var visit [MAXN]bool
var G [MAXN][MAXN]bool
var dp [MAXN][1 << MAXN]int

/*
 * func numSquarefulPerms(nums []int) int {
 * 	n = len(nums)
 *	// NOTE: 别忘了排序!
 * 	sort.Ints(nums)
 * 	ans, cur = 0, nil
 * 	for i := range visit {
 * 		visit[i] = false
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		for j := i + 1; j < n; j++ {
 * 			v := isSquare(nums[i] + nums[j])
 * 			G[i][j] = v
 * 			G[j][i] = v
 * 		}
 * 	}
 * 	backtrack(nums, 0)
 * 	return ans
 * }
 */

// DP + 状态压缩
func numSquarefulPerms(nums []int) int {
	n = len(nums)
	sort.Ints(nums)
	ans = 0
	for i := 0; i < MAXN; i++ {
		for s := 0; s < 1<<MAXN; s++ {
			dp[i][s] = 0
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v := isSquare(nums[i] + nums[j])
			G[i][j] = v
			G[j][i] = v
		}
	}
	// 初始条件
	for i := 0; i < n; i++ {
		// NOTE: 如果遇到重复数字, 只取第一个. 要去重嘛...
		// 比如这里有 2 个 `2`, 那么第二个 `2` 作为起始位置的排列是 0 个,
		// 因为我们强行限制了第二个 `2` 必须在第一个的后面.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		dp[i][1<<i] = 1
	}
	lim := 1 << n
	// 枚举每个状态
	for s := 1; s < lim; s++ {
		// 我们看看当前位置...
		for i := 0; i < n; i++ {
			// 先剪枝: 如果之前没这个排列法就跳过...
			if dp[i][s] == 0 {
				continue
			}
			// 再看看下个位置...
			for j := 0; j < n; j++ {
				// 已经选中了就跳过...
				if (s>>j)&1 != 0 {
					continue
				}
				// 如果和当前最后位置和不出完全平方也跳过...
				if !G[i][j] {
					continue
				}
				// 这里去个重, 这样下面就不用 Morris 式去重了...
				if j > 0 && nums[j] == nums[j-1] && (s>>(j-1))&1 == 0 {
					continue
				}
				dp[j][s|(1<<j)] += dp[i][s]
			}
		}
	}
	// 统计最终状态, 枚举每个最终位置的值...
	for i := 0; i < n; i++ {
		ans += dp[i][lim-1]
	}
	// 此处去个重 @Morris...
	/*
		rate := 0
		for i := 0; i < n; i++ {
			rate++
			if i+1 == n || nums[i] != nums[i+1] {
				ans /= factor(rate)
				rate = 0
			}
		}
	*/
	return ans
}

func isSquare(x int) bool {
	if x < 0 {
		return false
	}
	if x <= 1 {
		return true
	}
	p, q := 1, x
	for p < q {
		mid := (p + q + 1) >> 1
		if mid*mid > x {
			q = mid - 1
		} else {
			p = mid
		}
	}
	return p*p == x
}

func backtrack(arr []int, k int) {
	if k == n {
		ans++
		return
	}
	for i := 0; i < n; i++ {
		if visit[i] {
			continue
		}
		if i > 0 && arr[i] == arr[i-1] && !visit[i-1] {
			continue
		}
		if k > 0 && !G[cur[len(cur)-1]][i] {
			continue
		}
		visit[i] = true
		cur = append(cur, i)
		backtrack(arr, k+1)
		visit[i] = false
		cur = cur[:len(cur)-1]
	}
}

func factor(x int) int {
	ans := 1
	for x > 0 {
		ans *= x
		x--
	}
	return ans
}
