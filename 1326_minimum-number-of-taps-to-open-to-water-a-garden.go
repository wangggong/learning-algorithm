/*
 * @lc app=leetcode.cn id=minimum-number-of-taps-to-open-to-water-a-garden lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1326] 灌溉花园的最少水龙头数目
 *
 * https://leetcode-cn.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/description/
 *
 * algorithms
 * Hard (47.49%)
 * Total Accepted:    6.6K
 * Total Submissions: 13.9K
 * Testcase Example:  '5\n[3,4,1,1,0,0]'
 *
 * 在 x 轴上有一个一维的花园。花园长度为 n，从点 0 开始，到点 n 结束。
 *
 * 花园里总共有 n + 1 个水龙头，分别位于 [0, 1, ..., n] 。
 *
 * 给你一个整数 n 和一个长度为 n + 1 的整数数组 ranges ，其中 ranges[i] （下标从 0 开始）表示：如果打开点 i
 * 处的水龙头，可以灌溉的区域为 [i -  ranges[i], i + ranges[i]] 。
 *
 * 请你返回可以灌溉整个花园的 最少水龙头数目 。如果花园始终存在无法灌溉到的地方，请你返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：n = 5, ranges = [3,4,1,1,0,0]
 * 输出：1
 * 解释：
 * 点 0 处的水龙头可以灌溉区间 [-3,3]
 * 点 1 处的水龙头可以灌溉区间 [-3,5]
 * 点 2 处的水龙头可以灌溉区间 [1,3]
 * 点 3 处的水龙头可以灌溉区间 [2,4]
 * 点 4 处的水龙头可以灌溉区间 [4,4]
 * 点 5 处的水龙头可以灌溉区间 [5,5]
 * 只需要打开点 1 处的水龙头即可灌溉整个花园 [0,5] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 3, ranges = [0,0,0,0]
 * 输出：-1
 * 解释：即使打开所有水龙头，你也无法灌溉整个花园。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 * ranges.length == n + 1
 * 0 <= ranges[i] <= 100
 *
 *
 */

/*
 * // 一顿操作, 大样例过不了...
 * func minOperations(nums1 []int, nums2 []int) int {
 *     n, m := len(nums1), len(nums2)
 *     if (n > m && n > m*6) || (m > n && m > n*6) {
 *         return -1
 *     }
 *     sort.Ints(nums1)
 *     sort.Ints(nums2)
 *     sum1, sum2 := 0, 0
 *     for _, n := range nums1 {
 *         sum1 += n
 *     }
 *     for _, n := range nums2 {
 *         sum2 += n
 *     }
 *     // fmt.Println(sum1, sum2)
 *     if sum1 > sum2 {
 *         ans1 := getUp(nums1, nums2, sum1, sum2)
 *         ans2 := getDown(nums1, nums2, sum1, sum2)
 *         return min(ans1, ans2)
 *     } else {
 *         ans1 := getUp(nums2, nums1, sum2, sum1)
 *         ans2 := getDown(nums2, nums1, sum2, sum1)
 *         return min(ans1, ans2)
 *     }
 *     return -1
 * }
 *
 * func getUp(arr1, arr2 []int, sum1, sum2 int) int {
 *     n, m := len(arr1), len(arr2)
 *     ans := 0
 *     for i := 0; i < n && sum1 > sum2; i++ {
 *         ans++
 *         diff := sum1 - sum2
 *         k := min(diff, arr1[i] - 1)
 *         sum1 -= k
 *     }
 *     for i := 0; i < m && sum1 > sum2; i++ {
 *         ans++
 *         diff := sum1 - sum2
 *         k := min(diff, 6 - arr2[i])
 *         sum2 += k
 *
 *     }
 *     return ans
 * }
 *
 * func getDown(arr1, arr2 []int, sum1, sum2 int) int {
 *     n, m := len(arr1), len(arr2)
 *     ans := 0
 *     for i := 0; i < m && sum1 > sum2; i++ {
 *         ans++
 *         diff := sum1 - sum2
 *         k := min(diff, 6 - arr2[i])
 *         sum2 += k
 *
 *     }
 *     for i := 0; i < n && sum1 > sum2; i++ {
 *         ans++
 *         diff := sum1 - sum2
 *         k := min(diff, arr1[i] - 1)
 *         sum1 -= k
 *     }
 *     return ans
 * }
 */

// 参考: https://leetcode-cn.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/solution/guan-gai-hua-yuan-de-zui-shao-shui-long-tou-shu-3/
//
// 就是跳跃游戏套了个壳子嘛, 傻逼了吧...
// @20230221 更新, 二刷不会系列.
func minTaps(n int, ranges []int) int {
	right := make([]int, n+1)
	right[0] = ranges[0]
	for i := 1; i <= n; i++ {
		l, r := max(i-ranges[i], 0), min(i+ranges[i], n)
		right[i] = max(right[i-1], r)
		right[l] = max(right[l], r)
	}
	ans := 1
	for i, curr, next := 0, right[0], 0; curr < n; curr, ans = next, ans+1 {
		for next = curr; i <= curr; i++ {
			next = max(next, right[i])
		}
		if curr == next {
			return -1
		}
	}
	return ans
}

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
